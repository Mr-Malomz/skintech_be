package main

import (
	"net/http"

	"github.com/Mr-Malomz/skintech_be/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello World",
		})
	})

	//run database
	models.ConnectDB()

	r.Run()
}
