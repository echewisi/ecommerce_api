package routes

import (
	"github.com/echewisi/ecommerce_api/controllers"
	"github.com/echewisi/ecommerce_api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, jwtSecret string, controllers struct {
	Auth    *controllers.AuthController
	Product *controllers.ProductController
	Order   *controllers.OrderController
}) {
	// Public routes
	router.POST("/register", controllers.Auth.Register)
	router.POST("/login", controllers.Auth.Login)

	// Protected routes (Authentication required)
	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// Order routes
		authGroup.POST("/orders", controllers.Order.PlaceOrder)
		authGroup.PUT("/orders/:id/cancel", controllers.Order.CancelOrder)
		authGroup.GET("/orders", controllers.Order.GetUserOrders)
	}

	// Admin routes (Authentication and Admin Access required)
	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.AuthMiddleware(jwtSecret), middleware.AdminMiddleware())
	{
		// Product routes
		adminGroup.POST("/products", controllers.Product.CreateProduct)
		adminGroup.GET("/products", controllers.Product.GetAllProducts)
		adminGroup.PUT("/products/:id", controllers.Product.UpdateProduct)
		adminGroup.DELETE("/products/:id", controllers.Product.DeleteProduct)

		// Order routes
		adminGroup.PUT("/orders/:id", controllers.Order.UpdateOrder)
	}
}
