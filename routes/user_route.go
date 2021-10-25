package routes

import (
	"github.com/Mr-Malomz/skintech_be/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.GET("/user/:userId", controllers.GetAUser())
}
