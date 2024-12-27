package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/echewisi/ecommerce_api/models"
	"github.com/echewisi/ecommerce_api/services"
)

type ProductController struct {
	ProductService *services.ProductService
}

// NewProductController creates a new instance of ProductController
func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

// CreateProduct handles adding a new product (Admin only)
func (pc *ProductController) CreateProduct(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.ProductService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}

// GetAllProducts retrieves all products
func (pc *ProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}