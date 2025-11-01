package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthResponse represents the response for health checks
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// ReadinessResponse represents the response for readiness checks
type ReadinessResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
	Timestamp time.Time        `json:"timestamp"`
}

// HealthCheck handles health check requests
func HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
	}
	
	c.JSON(http.StatusOK, response)
}

// ReadinessCheck handles readiness check requests
func ReadinessCheck(c *gin.Context) {
	// In a real implementation, you would check:
	// - Database connectivity
	// - Cache connectivity
	// - External service connectivity
	
	services := map[string]string{
		"database": "connected",
		"cache":    "connected",
	}
	
	response := ReadinessResponse{
		Status:    "ready",
		Services:  services,
		Timestamp: time.Now(),
	}
	
	c.JSON(http.StatusOK, response)
}