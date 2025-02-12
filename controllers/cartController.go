package controllers

import (
	"ecommerce-api/db"
	"ecommerce-api/models"
	"net/http"
	

	"github.com/gin-gonic/gin"
)


// Add to Cart
func AddToCart(c *gin.Context) {
    
	var cartItem models.Cart

	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "invalid put"})
	}

	// Check if user exists
    var user models.User
    if err := db.DB.First(&user, cartItem.UserID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
        return
    }

    // Check if product exists
    var product models.Product
    if err := db.DB.First(&product, cartItem.ProductID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Product does not exist"})
        return
    }

	if err := db.DB.Create(&cartItem).Error; err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": "failed to add to cart"})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success fully add  to cart "})

}



func ViewCart(c *gin.Context){

	var cartItems []models.Cart

	UserID := c.Param("user_id")


	if err := db.DB.Where("user_id = ? ",UserID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "invalid input"})
		return
	}

	c.JSON(http.StatusOK,cartItems)
}


func ViewAllCart(c *gin.Context){

	var cartItems []models.Cart


	if err := db.DB.Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "invalid input"})
		return
	}

	c.JSON(http.StatusOK,cartItems)
}


func RemoveFromCart(c *gin.Context){
	var cartItem models.Cart

	cartID := c.Param("id")

	if err := db.DB.Where("id = ?", cartID).Delete(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}