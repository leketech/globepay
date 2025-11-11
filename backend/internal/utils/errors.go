package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError represents a standardized API error
type APIError struct {
	Error     string `json:"error"`
	Code      string `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

// ErrorResponse sends a standardized error response
func ErrorResponse(c *gin.Context, statusCode int, code, message string) {
	c.JSON(statusCode, APIError{
		Error:     message,
		Code:      code,
		Timestamp: c.GetTime("request_time").Unix(),
	})
}

// BadRequest sends a 400 Bad Request error
func BadRequest(c *gin.Context, code, message string) {
	ErrorResponse(c, http.StatusBadRequest, code, message)
}

// Unauthorized sends a 401 Unauthorized error
func Unauthorized(c *gin.Context, code, message string) {
	ErrorResponse(c, http.StatusUnauthorized, code, message)
}

// Forbidden sends a 403 Forbidden error
func Forbidden(c *gin.Context, code, message string) {
	ErrorResponse(c, http.StatusForbidden, code, message)
}

// NotFound sends a 404 Not Found error
func NotFound(c *gin.Context, code, message string) {
	ErrorResponse(c, http.StatusNotFound, code, message)
}

// InternalServerError sends a 500 Internal Server Error
func InternalServerError(c *gin.Context, code, message string) {
	ErrorResponse(c, http.StatusInternalServerError, code, message)
}
