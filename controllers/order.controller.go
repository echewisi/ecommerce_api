package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/echewisi/ecommerce_api/models"
	"github.com/echewisi/ecommerce_api/services"
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
	userID, _ := c.Get("userID")

	var request struct {
		Items []models.OrderItem `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := oc.OrderService.PlaceOrder(userID.(uint), request.Items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order placed successfully", "order": order})
}

// CancelOrder handles order cancellation
func (oc *OrderController) CancelOrder(c *gin.Context) {
	orderID := c.Param("id")

	if err := oc.OrderService.CancelOrder(parseUint(orderID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}

// parseUint is a helper function to parse uint from string
func parseUint(s string) uint {
	value, _ := strconv.ParseUint(s, 10, 32)
	return uint(value)
}

// UpdateOrder handles updating the status or details of an order
func (oc *OrderController) UpdateOrder(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if !isAdmin.(bool) {
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

