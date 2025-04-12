package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"inventory-api/config"
	"inventory-api/models"
)

// Get stock for a product
func GetInventoryByProductID(c *gin.Context) {
	productID := c.Param("product_id")
	var inventory models.Inventory

	if err := config.DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// Update stock (add/subtract)
func UpdateInventory(c *gin.Context) {
	productID := c.Param("product_id")
	var inventory models.Inventory

	if err := config.DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	var input struct {
		Quantity int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inventory.Quantity += input.Quantity // bisa tambah atau kurangin
	config.DB.Save(&inventory)

	c.JSON(http.StatusOK, inventory)
}

// Create inventory (first time add stock)
func CreateInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&inventory)
	c.JSON(http.StatusOK, inventory)
}
