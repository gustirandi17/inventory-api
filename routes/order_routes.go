package routes

import (
	"inventory-api/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	orders := r.Group("/orders")
	{
		orders.POST("/", controllers.CreateOrder)
		orders.GET("/:id", controllers.GetOrderByID)
	}
}
