package controllers

import (
	"ecommerce-api/db"
	"ecommerce-api/models"
	"ecommerce-api/utils" // Ensure utils package exists
	"net/http"
    "fmt"
	"github.com/gin-gonic/gin"
)

// ShortenURL - Creates a new short URL
func ShortenURL(c *gin.Context) {
	var request struct {
		URL string `json:"url"`
	}

	// Bind request JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return // 🔴 MISSING RETURN FIXED
	}

	// Check if URL already exists in DB
	var existingURL models.URL
	if err := db.DB.Where("original_url = ?", request.URL).First(&existingURL).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"short_url": existingURL.ShortURL})
		return
	}

	// Generate short URL
	shortURL := utils.GenerateShortURL(request.URL)
	fmt.Printf("dataaaa",shortURL)

	// Save to MySQL
	url := models.URL{ShortURL: shortURL, OriginalURL: request.URL}
	if err := db.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// RedirectURL - Redirects to the original URL
func RedirectURL(c *gin.Context) {
	shortURL := c.Param("short_url")

	// Fetch from DB
	var url models.URL
	if err := db.DB.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Redirect to original URL
	c.Redirect(http.StatusFound, url.OriginalURL)
}
