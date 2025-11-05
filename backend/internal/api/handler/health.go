package handler

import (
	"net/http"

	"globepay/internal/domain/service"
	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	healthService *service.HealthService
}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler(healthService *service.HealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
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
	// Check if all dependencies are ready
	dbStatus := "unknown"
	redisStatus := "unknown"
	
	if h.healthService != nil {
		dbOk := h.healthService.CheckDatabaseSimple()
		redisOk := h.healthService.CheckRedisSimple()
		
		if dbOk {
			dbStatus = "ok"
		} else {
			dbStatus = "error"
		}
		
		if redisOk {
			redisStatus = "ok"
		} else {
			redisStatus = "error"
		}
	}

	statusCode := http.StatusOK
	if dbStatus != "ok" || redisStatus != "ok" {
		statusCode = http.StatusServiceUnavailable
	}

	c.JSON(statusCode, gin.H{
		"status":  "ready",
		"message": "Globepay API is ready to serve requests",
		"dependencies": map[string]string{
			"database": dbStatus,
			"redis":    redisStatus,
		},
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