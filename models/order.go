package models

import "time"

// OrderStatus defines possible statuses for an order
type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Completed OrderStatus = "Completed"
	Canceled  OrderStatus = "Canceled"
)

// Order represents an order placed by a user
type Order struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `gorm:"not null" json:"user_id"`
	Products    []OrderItem `gorm:"foreignKey:OrderID" json:"products"`
	TotalAmount float64     `gorm:"not null" json:"total_amount"`
	Status      OrderStatus `gorm:"default:'Pending'" json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
}
