package main

import (
	"ecommerce-api/db"
    "ecommerce-api/routes"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "fmt"
)


func main() {
	db.Init()
    defer db.Close() // Close the database connection when the app ends

    // Initialize Gin router
    router := gin.Default()
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Set up routes
    routes.SetupRoutes(router)

    // Start the server
    fmt.Println("Server running on http://localhost:8080")
    router.Run(":8080")
}