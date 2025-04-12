package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"inventory-api/config"
	"inventory-api/models"
)

// Create order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.OrderDate = time.Now()

	// Update inventory stock (reduce)
	var inventory models.Inventory
	if err := config.DB.Where("product_id = ?", order.ProductID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product inventory not found"})
		return
	}

	if inventory.Quantity < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock"})
		return
	}

	inventory.Quantity -= order.Quantity
	config.DB.Save(&inventory)

	// Save order
	config.DB.Create(&order)
	c.JSON(http.StatusOK, order)
}

// Get order by ID
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}