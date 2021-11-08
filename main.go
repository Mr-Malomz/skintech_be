package main

import (
	"net/http"

	"github.com/Mr-Malomz/skintech_be/models"
	"github.com/Mr-Malomz/skintech_be/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello World",
		})
	})

	//cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	//run database
	models.ConnectDB()

	//middlewares
	r.Use(cors.New(config))

	//routes
	routes.AuthRoute(r)
	routes.UserRoute(r)
	routes.ImageRoute(r)
	routes.CollectionRoute(r)

	r.Run()
}
