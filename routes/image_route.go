package routes

import (
	"github.com/Mr-Malomz/skintech_be/controllers"
	"github.com/gin-gonic/gin"
)

func ImageRoute(r *gin.Engine) {
	r.POST("/upload", controllers.UploadImage())
}
