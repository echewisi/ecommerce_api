package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/echewisi/ecommerce_api/config"
	"github.com/echewisi/ecommerce_api/controllers"
	"github.com/echewisi/ecommerce_api/database"
	"github.com/echewisi/ecommerce_api/routes"
	"github.com/echewisi/ecommerce_api/services"
	"github.com/echewisi/ecommerce_api/repositories"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

// @title E-Commerce API
// @version 1.0
// @description This is a RESTful API for managing an e-commerce application.

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	config := config.LoadConfig()

	// Initialize database connection using GORM
	db, err := database.ConnectDB(config.DB)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, config.JWTKey)
	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo, productRepo)

	// Initialize controllers
	controllers := struct {
		Auth    *controllers.AuthController
		Product *controllers.ProductController
		Order   *controllers.OrderController
	}{
		Auth:    controllers.NewAuthController(authService),
		Product: controllers.NewProductController(productService),
		Order:   controllers.NewOrderController(orderService),
	}

	// Initialize router and routes
	router := gin.Default()
	routes.InitRoutes(router, config.JWTKey, controllers)

	// Serve your manually created Swagger JSON file
	// router.GET("/swagger/swagger.json", func(c *gin.Context) {
	// 	c.File("../docs/swagger.json") 
	// })

	// Configure Swagger UI to use your custom Swagger JSON file
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/docs/swagger.json")))
	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
