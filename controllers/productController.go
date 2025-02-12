package controllers

import (
    "net/http"
	"ecommerce-api/models"
    "ecommerce-api/db"
    "fmt"
    "github.com/gin-gonic/gin"
)

// ProductIndex function to handle product listing
func ProductIndex(c *gin.Context) {
	 var products []models.Product
	 result := db.DB.Find(&products)
	 if result.Error != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error": result.Error.Error()})
		return
	 }
	 fmt.Printf("Products: %+v\n", products)
    c.JSON(http.StatusOK, products)
}


func ProductAdd(c *gin.Context){
    
    var product models.Product

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "product is invalid"})
        return // Add return to stop further execution
    }
	fmt.Printf("Product received: %+v\n", product) 

    if err := db.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product added"})
}



func UpdateProduct(c *gin.Context){
    id := c.Param("id")

    // here we query on database  match the id from request and  match with databse if noty then show error
    var product models.Product
    if err := db.DB.Where("id = ?", id).First(&product).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    if err := db.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})

}


func DeleteProduct(c *gin.Context){
    id := c.Param("id")

    var product models.Product
    if err := db.DB.Where("id = ?", id).First(&product).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := db.DB.Delete(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
        return
    }

    // Respond with a success message
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})



}