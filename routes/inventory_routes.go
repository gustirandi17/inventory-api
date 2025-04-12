package routes

import (
	"github.com/gin-gonic/gin"
	"inventory-api/controllers"
)

func InventoryRoutes(r *gin.Engine) {
	inv := r.Group("/inventory")
	{
		inv.POST("/", controllers.CreateInventory)
		inv.GET("/:product_id", controllers.GetInventoryByProductID)
		inv.PUT("/:product_id", controllers.UpdateInventory)
	}
}
