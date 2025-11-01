package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	// Add dependencies here
}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck handles health check requests
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Globepay API is running",
		"version": "1.0.0",
	})
}

// ReadyCheck handles readiness check requests
func (h *HealthHandler) ReadyCheck(c *gin.Context) {
	// Add logic to check if all dependencies are ready
	c.JSON(http.StatusOK, gin.H{
		"status":  "ready",
		"message": "Globepay API is ready to serve requests",
	})
}

// LiveCheck handles liveness check requests
func (h *HealthHandler) LiveCheck(c *gin.Context) {
	// Add logic to check if the service is alive
	c.JSON(http.StatusOK, gin.H{
		"status":  "alive",
		"message": "Globepay API is alive",
	})
}