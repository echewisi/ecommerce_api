package repositories

import (
	"github.com/echewisi/ecommerce_api/models"
	"errors"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// CreateOrder creates a new order in the database
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

// GetOrdersByUserID fetches orders for a specific user
func (r *OrderRepository) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Where("user_id = ?", userID).Preload("Products").Find(&orders).Error
	return orders, err
}

// FindOrderByID fetches an order by ID
func (r *OrderRepository) FindOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.DB.Preload("Products").First(&order, id).Error
	return &order, err
}

// UpdateOrder updates the status or details of an order
func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	return r.DB.Save(order).Error
}

// CancelOrder cancels an order if its status is "Pending"
func (r *OrderRepository) CancelOrder(orderID uint) error {
	var order models.Order
	err := r.DB.First(&order, orderID).Error
	if err != nil {
		return err
	}

	// Check if the order is in "Pending" status
	if order.Status != models.Pending {
		return errors.New("only orders with 'Pending' status can be canceled")
	}

	// Update the status to "Canceled"
	order.Status = models.Canceled
	return r.DB.Save(&order).Error
}