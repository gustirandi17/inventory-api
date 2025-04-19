// file_controller.go
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"inventory-api/models"
	"inventory-api/config"
)

func UploadImage(c *gin.Context) {
	productID := c.Param("product_id")
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membaca file"})
		return
	}
	defer file.Close()

	// Validasi ukuran (maks 2MB)
	if header.Size > 2<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file maksimal 2MB"})
		return
	}

	// Validasi ekstensi
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
		return
	}

	// Simpan file
	filename := fmt.Sprintf("%s%s", productID, ext)
	filepath := "uploads/" + filename
	out, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}
	defer out.Close()
	io.Copy(out, file)

	// Update URL gambar ke database
	var product models.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Update ImageURL dengan path file yang telah diupload
	product.ImageURL = "/uploads/" + filename

	// Simpan perubahan ke database
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui produk"})
		return
	}

	// Respons berhasil
	c.JSON(http.StatusOK, gin.H{
		"message":   "File berhasil diupload",
		"filename":  filename,
		"image_url": "/uploads/" + filename,
	})
}


func DownloadImage(c *gin.Context) {
	productID := c.Param("product_id")

	var product models.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Ambil nama file dari field ImageURL
	imagePath := "." + product.ImageURL // contoh: ./uploads/1.jpg

	c.File(imagePath)
}
