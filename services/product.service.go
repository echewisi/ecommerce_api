package services

import (
	"github.com/echewisi/ecommerce_api/models"
	"github.com/echewisi/ecommerce_api/repositories"
)

type ProductService struct {
	ProductRepo *repositories.ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}

// CreateProduct adds a new product
func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.ProductRepo.CreateProduct(product)
}

// GetAllProducts retrieves all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.ProductRepo.GetAllProducts()
}

// GetProductByID retrieves a single product by ID
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.ProductRepo.FindProductByID(id)
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.ProductRepo.UpdateProduct(product)
}

// DeleteProduct removes a product by ID
func (s *ProductService) DeleteProduct(id uint) error {
	return s.ProductRepo.DeleteProduct(id)
}
