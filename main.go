package main

import (
	"inventory-api/config"
	"inventory-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()

	routes.ProductRoutes(r)
	routes.InventoryRoutes(r)
	routes.OrderRoutes(r)
	routes.UploadRoutes(r)

	r.Static("/uploads", "./uploads")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running!"})
	})

	r.Run(":8080")
}
