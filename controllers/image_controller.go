package controllers

import (
	"net/http"

	"github.com/Mr-Malomz/skintech_be/dtos"
	"github.com/Mr-Malomz/skintech_be/helper"
	"github.com/gin-gonic/gin"
)

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "Select a file to upload"})
			return
		}

		imageUrl, err := helper.ImageUploadHelper(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "Error uploading file"})
			return
		}

		c.JSON(http.StatusOK,
			dtos.Response{Status: http.StatusOK, Message: "success!", Data: map[string]interface{}{"imageURL": imageUrl}},
		)
	}
}
