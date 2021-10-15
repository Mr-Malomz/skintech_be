package main

import (
	"net/http"

	"github.com/Mr-Malomz/skintech_be/models"
	"github.com/Mr-Malomz/skintech_be/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello World",
		})
	})

	//run database
	models.ConnectDB()

	//routes
	routes.AuthRoute(r)
	routes.UserRoute(r)
	routes.ImageRoute(r)
	routes.CollectionRoute(r)

	r.Run()
}
