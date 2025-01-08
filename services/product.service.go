package services

import (
	"github.com/echewisi/ecommerce_api/models"
	"github.com/echewisi/ecommerce_api/repositories"
	"github.com/google/uuid"
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
func (s *ProductService) GetProductByID(id uuid.UUID) (*models.Product, error) {
	return s.ProductRepo.FindProductByID(id)
}

// UpdateProduct updates an existing product
func (ps *ProductService) UpdateProduct(product *models.Product) error {
	// Assuming you have a repository that handles the database update
	err := ps.ProductRepo.UpdateProduct(product)
	if err != nil {
		return err
	}
	// Retrieve the updated product from the database after the update
	updatedProduct, err := ps.ProductRepo.FindProductByID(product.ID)
	if err != nil {
		return err
	}

	// Update the product object with the latest data
	*product = *updatedProduct
	return nil
}

// DeleteProduct removes a product by ID
func (s *ProductService) DeleteProduct(id uuid.UUID) error {
	return s.ProductRepo.DeleteProduct(id)
}
