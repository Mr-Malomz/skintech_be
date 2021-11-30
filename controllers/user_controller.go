package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Mr-Malomz/skintech_be/dtos"
	"github.com/Mr-Malomz/skintech_be/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		defer cancel()

		if err != nil {
			log.Fatal(err)
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
			dtos.Response{Status: http.StatusOK, Message: "success!", Data: map[string]interface{}{"user": foundUser}},
		)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.UpdateUser //from json body

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": user})
		defer cancel()

		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "Error updating user's detail"})
			return
		}

		//get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
			defer cancel()

			if err != nil {
				log.Fatal(err)
				c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "User not found"})
				return
			}
		}

		newUpdatedUser := models.User{
			Id:             updatedUser.Id,
			Firstname:      updatedUser.Firstname,
			Lastname:       updatedUser.Lastname,
			Phone_Number:   updatedUser.Phone_Number,
			Email:          updatedUser.Email,
			MDCN:           updatedUser.MDCN,
			Anual_lics_num: updatedUser.Anual_lics_num,
			Folio_num:      updatedUser.Folio_num,
			School_grad:    updatedUser.School_grad,
			Year_grad:      updatedUser.Year_grad,
			Year_fellow:    updatedUser.Year_fellow,
			Cert_url:       updatedUser.Cert_url,
			IsActive:       updatedUser.IsActive,
			IsVerified:     updatedUser.IsVerified,
			Created_At:     updatedUser.Created_At,
		}

		c.JSON(http.StatusOK,
			dtos.Response{Status: http.StatusOK, Message: "success!", Data: map[string]interface{}{"user": newUpdatedUser}},
		)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User

		results, err := userCollection.Find(ctx, bson.M{})
		defer cancel()

		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "Error getting list of users"})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				log.Fatal(err)
			}
			//filter returned data
			filterUser := models.User{
				Id:             singleUser.Id,
				Firstname:      singleUser.Firstname,
				Lastname:       singleUser.Lastname,
				Phone_Number:   singleUser.Phone_Number,
				Email:          singleUser.Email,
				MDCN:           singleUser.MDCN,
				Anual_lics_num: singleUser.Anual_lics_num,
				Folio_num:      singleUser.Folio_num,
				School_grad:    singleUser.School_grad,
				Year_grad:      singleUser.Year_grad,
				Year_fellow:    singleUser.Year_fellow,
				Cert_url:       singleUser.Cert_url,
				IsActive:       singleUser.IsActive,
				IsVerified:     singleUser.IsVerified,
				Created_At:     singleUser.Created_At,
			}

			users = append(users, filterUser)
		}

		c.JSON(http.StatusOK,
			dtos.Response{Status: http.StatusOK, Message: "success!", Data: map[string]interface{}{"users": users}},
		)
	}
}

func DeleteUserAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		defer cancel()

		if err != nil {
			log.Fatal(err)
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				dtos.Response{Status: http.StatusNotFound, Message: "User with specified ID not found!"},
			)
			return
		}

		c.JSON(http.StatusOK,
			dtos.Response{Status: http.StatusOK, Message: "success!", Data: map[string]interface{}{"user": "User successfully deleted!"}},
		)
	}
}
