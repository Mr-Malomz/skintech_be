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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var imageCollection *mongo.Collection = models.GetCollection(models.DB, "image")

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var image models.Images

		if err := c.BindJSON(&image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(image); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		//check if user exist
		_, err := imageCollection.CountDocuments(ctx, bson.M{"userId": image.UserId})
		defer cancel() //if it exceeds stipulated time, then cancel request
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "Invalid userId"})
			return
		}

		file, handler, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "Select a file to upload"})
			return
		}
		defer file.Close()
 
		imageUrl, err := helper.ImageUploadHelper(file, handler)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "Error uploading file"})
			return
		}

		//creating/inserting image
		newImage := models.Images{
			UserId:      image.UserId,
			Skin_diag:   image.Skin_diag,
			Loc_disease: image.Loc_disease,
			Dif_diag:    image.Dif_diag,
			Gender:      image.Gender,
			Age:         image.Age,
			Country:     image.Country,
			State:       image.State,
			Comp_desc:   image.Comp_desc,
			Fav_count:   image.Fav_count,
			Img_url:     imageUrl,
		}

		_, insertErr := imageCollection.InsertOne(ctx, newImage)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK,
			dtos.Response{Status: http.StatusOK, Message: "success!", Data: map[string]interface{}{"data": newImage}},
		)
	}
}
