package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	// Form file dengan key "image"
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	// Buat nama unik untuk file-nya
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))

	// Pastikan folder uploads ada
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}

	// Simpan ke folder uploads/
	filepath := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "File uploaded successfully",
		"image_url":  "/uploads/" + filename,
	})
}
