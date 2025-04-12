package routes

import (
	"github.com/gin-gonic/gin"
	"inventory-api/controllers"
)

func ProductRoutes(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.POST("/", controllers.CreateProduct)
		products.GET("/", controllers.GetProducts)
		products.GET("/:id", controllers.GetProductByID)
		products.PUT("/:id", controllers.UpdateProduct)
		products.DELETE("/:id", controllers.DeleteProduct)
	}
}
