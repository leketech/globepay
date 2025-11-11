package handler

import (
	"context"
	"net/http"
	"time"

	"globepay/internal/domain/service"

	"github.com/gin-gonic/gin"
)

// HealthResponse represents the response for health checks
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// ReadinessResponse represents the response for readiness checks
type ReadinessResponse struct {
	Status    string            `json:"status"`
	Services  map[string]string `json:"services"`
	Timestamp time.Time         `json:"timestamp"`
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
	// Get the service factory from the context
	serviceFactory := c.MustGet("serviceFactory").(*service.Factory)

	// Get the health service
	healthService := serviceFactory.GetHealthService()

	// Check all services
	ctx := context.Background()
	services := healthService.CheckAll(ctx)

	// Determine overall status
	status := "ready"
	for _, serviceStatus := range services {
		if serviceStatus != "connected" {
			status = "not ready"
			break
		}
	}

	response := ReadinessResponse{
		Status:    status,
		Services:  services,
		Timestamp: time.Now(),
	}

	c.JSON(http.StatusOK, response)
}
