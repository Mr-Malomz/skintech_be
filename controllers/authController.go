package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Mr-Malomz/skintech_be/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = models.GetCollection(models.DB, "user")
var validate = validator.New()

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(hashed)
}

func verifyPassword(hashedPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))

	return err == nil
}

//signup
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		//check if email already exist on the database
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel() //if it exceeds stipulated time, then cancel request
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for email"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This email already exist"})
			return
		}

		//hashing the password
		password := hashPassword(user.Password)

		//creating a user
		user.Password = password
		user.Id = primitive.NewObjectID()
		user.Firstname = c.PostForm("firstname")
		user.Lastname = c.PostForm("lastname")
		user.Email = c.PostForm("email")
		user.Created_At = time.Now()
		user.IsActive = false
		user.IsVerified = false

		_, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, gin.H{"message": `User created successfully!`})
	}
}
