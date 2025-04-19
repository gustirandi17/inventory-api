package routes

import (
	"github.com/gin-gonic/gin"
	"inventory-api/controllers"
)

func FileRoutes(r *gin.Engine) {
	r.POST("/upload/:product_id", controllers.UploadImage)
	r.GET("/download/:product_id", controllers.DownloadImage)
}
