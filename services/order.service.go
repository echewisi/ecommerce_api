package services

import (
	"errors"
	"github.com/echewisi/ecommerce_api/models"
	"github.com/echewisi/ecommerce_api/repositories"
)

type OrderService struct {
	OrderRepo   *repositories.OrderRepository
	ProductRepo *repositories.ProductRepository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(orderRepo *repositories.OrderRepository, productRepo *repositories.ProductRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo, ProductRepo: productRepo}
}

// PlaceOrder places a new order
func (s *OrderService) PlaceOrder(userID uint, items []models.OrderItem) (*models.Order, error) {
	var totalAmount float64

	// Validate product availability and calculate total
	for i, item := range items {
		product, err := s.ProductRepo.FindProductByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}
		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}
		// Update item price and total
		items[i].Price = product.Price
		totalAmount += float64(item.Quantity) * product.Price
	}

	// Create the order
	order := &models.Order{
		UserID:      userID,
		Products:    items,
		TotalAmount: totalAmount,
		Status:      models.Pending,
	}
	err := s.OrderRepo.CreateOrder(order)
	return order, err
}

// CancelOrder cancels an order if it's in "Pending" status
func (s *OrderService) CancelOrder(orderID uint) error {
	return s.OrderRepo.CancelOrder(orderID)
}

// GetUserOrders retrieves all orders for a user
func (s *OrderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.OrderRepo.GetOrdersByUserID(userID)
}

// UpdateOrder updates the status or details of an existing order
func (s *OrderService) UpdateOrder(order *models.Order) error {
	// Retrieve the existing order
	existingOrder, err := s.OrderRepo.FindOrderByID(order.ID)
	if err != nil {
		return errors.New("order not found")
	}

	// Update fields
	existingOrder.Status = order.Status
	existingOrder.TotalAmount = order.TotalAmount // Optional if the total amount needs updates
	return s.OrderRepo.UpdateOrder(existingOrder)
}

