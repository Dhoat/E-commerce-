package controllers

import (
    "ecommerce-api/models"
    "ecommerce-api/db"
    "net/http"
    "github.com/gin-gonic/gin"
   "golang.org/x/crypto/bcrypt"
   "ecommerce-api/utils"
    
)

// Register function to handle user registration
// func Register(c *gin.Context) {
//     var user models.User

//     if err := c.ShouldBindJSON(&user); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
//         return
//     }

//     // Check if user exists
//     var existingUser models.User
//     if err := db.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
//         c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
//         return
//     }

//     if err := db.DB.Create(&user).Error; err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
//         return
//     }

//     c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
// }
// func Login(c *gin.Context) {
//     var requestBody struct {
//         Username string `json:"username"`
//         Password string `json:"password"`
//     }

//     if err := c.ShouldBindJSON(&requestBody); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
//         return
//     }

//     var user models.User
//     if err := db.DB.Where("username = ?", requestBody.Username).First(&user).Error; err != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
//         return
//     }

//     if user.Password != requestBody.Password {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
// }


func Register(c *gin.Context){
    
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil{
        c.JSON(http.StatusBadRequest,gin.H{"error": "Invalid input"})
    } 


    // check if user already regsiter
    
    var existingUser   models.User

    if err := db.DB.Where("username = ? ", user.Username).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusBadRequest,gin.H{"error": "user already exit "})
        return
    }

    // hash for password 
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user.Password = string(hashedPassword)

   if err := db.DB.Create(&user).Error; err != nil {
     
    c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to create user"})
    return

   }

   c.JSON(http.StatusCreated,gin.H{"message": "user Regsirter succesfully"})
}


func Login(c *gin.Context) {
    var requestBody struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var user models.User
    if err := db.DB.Where("username = ?", requestBody.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // Compare the hashed password with the input password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token, err := utils.GenerateJWT(user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token,"username":user.Username})
}