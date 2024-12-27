package utils

import (
	_"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{
		"error": err,
	})
}
