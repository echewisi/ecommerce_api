package repositories

import (
	"github.com/echewisi/ecommerce_api/models"
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
func (r *ProductRepository) FindProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}

// UpdateProduct updates a product in the database
func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	return r.DB.Save(product).Error
}

// DeleteProduct deletes a product by ID
func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}
