package controllers

import (
	"net/http"

	"github.com/echewisi/ecommerce_api/models"
	"github.com/echewisi/ecommerce_api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	OrderService *services.OrderService
}

// NewOrderController creates a new instance of OrderController
func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

// PlaceOrder handles placing a new order
func (oc *OrderController) PlaceOrder(c *gin.Context) {
    // Get user ID from context
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Type assert userID as uuid.UUID
    uid, ok := userID.(uuid.UUID)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
        return
    }

    // Create a request structure for order items
    var request struct {
        Items []models.OrderItem `json:"items" binding:"required"`
    }

    // Bind the JSON request body to the request struct
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Logic to place the order using the 'uid' and 'request.Items'
    order, err := oc.OrderService.PlaceOrder(uid, request.Items)
    if err != nil {
        // If error is related to insufficient stock or product not found, handle it
        if err.Error() == "insufficient stock" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for one or more products"})
        } else if err.Error() == "product not found" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "One or more products not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order: " + err.Error()})
        }
        return
    }

    // Return success response with order details
    c.JSON(http.StatusCreated, gin.H{
        "message": "Order placed successfully",
        "order":   order, // Optionally include the order details in the response
    })
}


// CancelOrder handles order cancellation
func (oc *OrderController) CancelOrder(c *gin.Context) {
	orderID := c.Param("id")
	oid, err := uuid.Parse(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := oc.OrderService.CancelOrder(oid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}

// GetUserOrders retrieves all orders for the authenticated user
func (oc *OrderController) GetUserOrders(c *gin.Context) {
    // Get user ID from context
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Type assert userID as uuid.UUID
    uid, ok := userID.(uuid.UUID)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
        return
    }

    // Logic to retrieve orders for the user using the 'uid'
    orders, err := oc.OrderService.GetUserOrders(uid)
    if err != nil {
        // Handle specific error scenarios
        if err.Error() == "no orders found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "No orders found for the user"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders: " + err.Error()})
        }
        return
    }

    // Return success response with the orders
    c.JSON(http.StatusOK, gin.H{
        "orders": orders, // Return the list of orders
    })
}


// UpdateOrder handles updating the status or details of an order
func (oc *OrderController) UpdateOrder(c *gin.Context) {
	isAdmin, exists := c.Get("isAdmin")
	if !exists || !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var request models.Order
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := oc.OrderService.UpdateOrder(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}
