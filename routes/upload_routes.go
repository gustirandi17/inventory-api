package routes

import (
	"github.com/gin-gonic/gin"
	"inventory-api/controllers"
)

func UploadRoutes(r *gin.Engine) {
	r.POST("/upload", controllers.UploadImage)
}
