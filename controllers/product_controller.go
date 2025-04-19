package controllers

import (
	"net/http"
	"inventory-api/config"
	"inventory-api/models"
	"github.com/gin-gonic/gin"
)

// Create product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Get product by ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

//Update product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
		return
	}

	var updated models.Product
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = updated.Name
	product.Description = updated.Description
	product.Price = updated.Price
	product.Category = updated.Category
	product.ImageURL = updated.ImageURL

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}


// Delete product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
		return
	}
	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product dihapus"})
}