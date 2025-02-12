package controllers

import (
	"ecommerce-api/db"
	"ecommerce-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Checkout - Convert Cart Items into an Order
func Checkout(c *gin.Context) {
	var cartItems []models.Cart
	userIDStr := c.Param("user_id") // 🔹 Get user_id from URL param (as string)

	// ✅ FIX 1: Convert userID from string to uint
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"}) // 🔹 Return error if conversion fails
		return
	}

	// ✅ FIX 2: Fetch all cart items of the user
	if err := db.DB.Where("user_id = ?", uint(userID)).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	// ✅ FIX 3: Check if cart is empty
	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	// ✅ FIX 4: Calculate total amount from cart items
	var totalAmount float64
	for _, item := range cartItems {
		var product models.Product
		if err := db.DB.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product details"})
			return
		}
		totalAmount += float64(item.Quantity) * product.Price
	}

	// ✅ FIX 5: Create an Order
	order := models.Order{
		UserID:      uint(userID), // 🔹 Convert userID correctly
		TotalAmount: totalAmount,
		Status:      "pending",
	}

	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// ✅ FIX 6: Clear the cart after order is placed
	if err := db.DB.Where("user_id = ?", uint(userID)).Delete(&models.Cart{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}

	// ✅ SUCCESS RESPONSE
	c.JSON(http.StatusCreated, gin.H{
		"message": "Order placed successfully",
		"order":   order,
	})
}
