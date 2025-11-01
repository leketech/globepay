package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// APIInfo represents the API information response
type APIInfo struct {
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Description string    `json:"description"`
	DocsURL     string    `json:"docs_url"`
	Timestamp   time.Time `json:"timestamp"`
	Endpoints   []string  `json:"endpoints"`
}

// RootHandler handles root path requests
func RootHandler(c *gin.Context) {
	apiInfo := APIInfo{
		Name:        "Globepay API",
		Version:     "1.0.0",
		Description: "Globepay is a modern payment processing platform API",
		DocsURL:     "/swagger", // Placeholder for Swagger docs
		Timestamp:   time.Now(),
		Endpoints: []string{
			"GET /health - Health check endpoint",
			"GET /health/ready - Readiness check endpoint",
			"GET /metrics - Prometheus metrics endpoint",
			"POST /api/v1/auth/login - User login",
			"POST /api/v1/auth/register - User registration",
			"POST /api/v1/auth/refresh - Refresh authentication token",
			"POST /api/v1/auth/forgot-password - Forgot password",
			"POST /api/v1/auth/reset-password - Reset password",
			"GET /api/v1/user/profile - Get user profile (protected)",
			"PUT /api/v1/user/profile - Update user profile (protected)",
			"GET /api/v1/user/accounts - Get user accounts (protected)",
			"POST /api/v1/user/accounts - Create user account (protected)",
			"GET /api/v1/transfers - Get transfers (protected)",
			"GET /api/v1/transfers/:id - Get transfer by ID (protected)",
			"POST /api/v1/transfers - Create transfer (protected)",
			"POST /api/v1/transfers/:id/cancel - Cancel transfer (protected)",
			"GET /api/v1/transfers/rates - Get exchange rates (protected)",
			"GET /api/v1/transactions - Get transactions (protected)",
			"GET /api/v1/transactions/:id - Get transaction by ID (protected)",
			"GET /api/v1/beneficiaries - Get beneficiaries (protected)",
			"POST /api/v1/beneficiaries - Create beneficiary (protected)",
			"PUT /api/v1/beneficiaries/:id - Update beneficiary (protected)",
			"DELETE /api/v1/beneficiaries/:id - Delete beneficiary (protected)",
		},
	}

	c.JSON(http.StatusOK, apiInfo)
}