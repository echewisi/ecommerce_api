package models

import (
	"time"

	"github.com/google/uuid" // Import the UUID package
)

// OrderStatus defines possible statuses for an order
type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Completed OrderStatus = "Completed"
	Canceled  OrderStatus = "Canceled"
)

// Order represents an order placed by a user
type Order struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID      uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"` // User ID also changed to UUID
	Products    []OrderItem `gorm:"foreignKey:OrderID" json:"products"`
	TotalAmount float64     `gorm:"not null" json:"total_amount"`
	Status      OrderStatus `gorm:"default:'Pending'" json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"`
}
