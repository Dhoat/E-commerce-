package routes

import (
	"ecommerce-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Define the routes for the API
	api := router.Group("/api")
	{
		// Register route
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/products", controllers.ProductIndex)
		api.POST("/products", controllers.ProductAdd)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)


		// cart 
		api.POST("/cart", controllers.AddToCart)
		api.GET("/cart", controllers.ViewAllCart)
		api.GET("/cart/:user_id", controllers.ViewCart)
		api.DELETE("/cart/:id", controllers.RemoveFromCart)
		
		
       
        // order 
		api.POST("/checkout/:user_id", controllers.Checkout)



		api.POST("/shorten", controllers.ShortenURL)
		api.GET("/:short_url", controllers.RedirectURL)
		
	}
}
