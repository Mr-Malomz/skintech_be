package routes

import (
	"github.com/Mr-Malomz/skintech_be/controllers"
	"github.com/Mr-Malomz/skintech_be/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	authorized := r.Group("/")
	authorized.Use(middlewares.AuthJWTMiddleware())

	{
		authorized.GET("/users", controllers.GetAllUsers())
		authorized.GET("/user/:userId", controllers.GetAUser())
		authorized.POST("/user/:userId", controllers.UpdateUser())
		authorized.DELETE("/user/:userId", controllers.DeleteUserAccount())
	}
}
