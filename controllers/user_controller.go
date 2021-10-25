package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Mr-Malomz/skintech_be/dtos"
	"github.com/Mr-Malomz/skintech_be/models"
	"github.com/gin-gonic/gin"

	// "github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var authCollection *mongo.Collection = models.GetCollection(models.DB, "user")

// var validate = validator.New()

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		userID := c.Param("userId")
		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "User not found"})
			return
		}

		foundUser := models.User{
			Id:             user.Id,
			Firstname:      user.Firstname,
			Lastname:       user.Lastname,
			Phone_Number:   user.Phone_Number,
			Email:          user.Email,
			MDCN:           user.MDCN,
			Anual_lics_num: user.Anual_lics_num,
			Folio_num:      user.Folio_num,
			School_grad:    user.School_grad,
			Year_grad:      user.Year_grad,
			Year_fellow:    user.Year_fellow,
			Cert_url:       user.Cert_url,
			IsActive:       user.IsActive,
			IsVerified:     user.IsVerified,
			Created_At:     user.Created_At,
		}

		c.JSON(http.StatusOK,
			dtos.Response{Status: http.StatusOK, Message: "Login successful!", Data: map[string]interface{}{"user": foundUser}},
		)
	}
}

func GetAllUsers() {

}

func DeleteUserAccount() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}