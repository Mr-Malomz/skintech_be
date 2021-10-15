package routes

import (
	"github.com/Mr-Malomz/skintech_be/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp())
	r.POST("/login", controllers.Login())
}
