package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Mr-Malomz/skintech_be/dtos"
	"github.com/Mr-Malomz/skintech_be/helper"
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
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "Error occured while checking for email"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "This email already exist"})
			return
		}

		//creating a user
		otp := helper.GenerateOTP().String()
		created_Date, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newUser := models.User{
			Password:   hashPassword(user.Password),
			Id:         primitive.NewObjectID(),
			Firstname:  user.Firstname,
			Lastname:   user.Lastname,
			Email:      user.Email,
			Created_At: created_Date,
			OTP:        otp,
			IsActive:   false,
			IsVerified: false,
		}

		_, insertErr := userCollection.InsertOne(ctx, newUser)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, dtos.Response{Status: http.StatusOK, Message: "User created successfully!", Data: nil})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		//check if email exist on the db
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.Response{Status: http.StatusInternalServerError, Message: "Incorrect email or password"},
			)
			return
		}

		//check if password is correct
		isValidPass := verifyPassword(foundUser.Password, user.Password)
		defer cancel()
		if !isValidPass {
			c.JSON(http.StatusInternalServerError,
				dtos.Response{Status: http.StatusInternalServerError, Message: "Incorrect email or password"},
			)
			return
		}

		//jwt

		//send response
		resp := models.User{
			Id:         foundUser.Id,
			Firstname:  foundUser.Firstname,
			Lastname:   foundUser.Lastname,
			Email:      foundUser.Email,
			Created_At: foundUser.Created_At,
			OTP:        foundUser.OTP,
			IsActive:   foundUser.IsActive,
			IsVerified: foundUser.IsVerified,
		}
		c.JSON(http.StatusOK, 
			dtos.Response{Status: http.StatusOK, Message: "Login successful!", Data: map[string]interface{}{"user": resp}},
		)
	}
}
