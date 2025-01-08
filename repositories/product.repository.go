package repositories

import (
	"github.com/echewisi/ecommerce_api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// CreateProduct creates a new product in the database
func (r *ProductRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

// GetAllProducts fetches all products
func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}

// FindProductByID fetches a product by ID
func (r *ProductRepository) FindProductByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, "id = ?", id).Error
	return &product, err
}

// UpdateProduct updates a product in the database
func (pr *ProductRepository) UpdateProduct(product *models.Product) error {
	// Build a map to update only the non-zero fields
	updates := make(map[string]interface{})

	// Check and add only non-empty fields to the updates map
	if product.Name != "" {
		updates["name"] = product.Name
	}
	if product.Description != "" {
		updates["description"] = product.Description
	}
	if product.Price != 0 {
		updates["price"] = product.Price
	}
	if product.Stock != 0 {
		updates["stock"] = product.Stock
	}

	// Only update fields that are not empty
	if len(updates) > 0 {
		if err := pr.DB.Model(&models.Product{}).Where("id = ?", product.ID).Updates(updates).Error; err != nil {
			return err
		}
	}
	return nil
}


// // You can also fetch the updated product (if needed) after the update like this:
// func (pr *ProductRepository) GetProductByID(id uuid.UUID) (*models.Product, error) {
// 	var product models.Product
// 	if err := pr.DB.First(&product, "id = ?", id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &product, nil
// }


// DeleteProduct deletes a product by ID
func (r *ProductRepository) DeleteProduct(id uuid.UUID) error {
	return r.DB.Delete(&models.Product{}, "id = ?", id).Error
}
