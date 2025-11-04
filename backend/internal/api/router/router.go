package router

import (
	"fmt"

	"globepay/internal/api/handler"
	"globepay/internal/api/middleware"
	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/metrics"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, serviceFactory *service.ServiceFactory, metrics *metrics.Metrics) {
	fmt.Println("Setting up routes...")

	// Apply CORS middleware to all routes
	r.Use(func(c *gin.Context) {
		c.Set("serviceFactory", serviceFactory)
		c.Next()
	})
	r.Use(middleware.CORS())
	r.Use(middleware.MetricsMiddleware(metrics))
	fmt.Println("CORS middleware applied")

	// Root endpoint
	r.GET("/", handler.RootHandler)
	fmt.Println("Registered root endpoint")

	// Health check endpoints
	r.GET("/health", handler.HealthCheck)
	r.GET("/health/ready", handler.ReadinessCheck)
	fmt.Println("Registered health endpoints")

	// Metrics endpoint
	r.GET("/metrics", handler.Metrics)
	fmt.Println("Registered metrics endpoint")

	// Simple test route
	r.GET("/test", func(c *gin.Context) {
		fmt.Println("Handling /test request")
		c.JSON(200, gin.H{"message": "Test route working"})
	})
	fmt.Println("Registered test endpoint")

	// Debug route to check if routes are working
	r.GET("/debug-routes", func(c *gin.Context) {
		fmt.Println("Handling /debug-routes request")
		c.JSON(200, gin.H{"message": "Debug route working"})
	})
	fmt.Println("Registered /debug-routes endpoint")

	// Authentication routes
	r.POST("/api/v1/auth/login", func(c *gin.Context) {
		fmt.Println("Handling /api/v1/auth/login request")
		handler.Login(c, serviceFactory)
	})
	fmt.Println("Registered /api/v1/auth/login endpoint")

	r.POST("/api/v1/auth/register", func(c *gin.Context) {
		fmt.Println("Handling /api/v1/auth/register request")
		handler.Register(c, serviceFactory)
	})
	fmt.Println("Registered /api/v1/auth/register endpoint")

	r.POST("/api/v1/auth/refresh", func(c *gin.Context) {
		fmt.Println("Handling /api/v1/auth/refresh request")
		handler.RefreshToken(c, serviceFactory)
	})
	fmt.Println("Registered /api/v1/auth/refresh endpoint")

	r.POST("/api/v1/auth/forgot-password", func(c *gin.Context) {
		fmt.Println("Handling /api/v1/auth/forgot-password request")
		handler.ForgotPassword(c, serviceFactory)
	})
	fmt.Println("Registered /api/v1/auth/forgot-password endpoint")

	r.POST("/api/v1/auth/reset-password", func(c *gin.Context) {
		fmt.Println("Handling /api/v1/auth/reset-password request")
		handler.ResetPassword(c, serviceFactory)
	})
	fmt.Println("Registered /api/v1/auth/reset-password endpoint")

	// Protected routes group
	protected := r.Group("/api/v1")
	// We need to get the JWT secret from the service factory
	jwtSecret := serviceFactory.GetJWTSecret()
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// User routes
		protected.GET("/user/profile", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/profile request")
			handler.GetUserProfile(c, serviceFactory)
		})
		protected.PUT("/user/profile", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/profile update request")
			handler.UpdateUserProfile(c, serviceFactory)
		})

		// Account routes
		protected.GET("/user/accounts", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/accounts request")
			handler.GetUserAccounts(c, serviceFactory)
		})
		protected.POST("/user/accounts", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/accounts create request")
			handler.CreateUserAccount(c, serviceFactory)
		})

		// Transfer routes
		protected.GET("/transfers", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers request")
			handler.GetTransfers(c, serviceFactory)
		})
		protected.GET("/transfers/:id", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers/:id request")
			handler.GetTransfer(c, serviceFactory)
		})
		protected.POST("/transfers", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers create request")
			handler.CreateTransfer(c, serviceFactory)
		})
		protected.POST("/transfers/:id/cancel", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers/:id/cancel request")
			handler.CancelTransfer(c, serviceFactory)
		})

		// Exchange rate routes
		protected.GET("/transfers/rates", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers/rates request")
			handler.GetExchangeRates(c, serviceFactory)
		})

		// Transaction routes
		protected.GET("/transactions", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transactions request")
			handler.GetTransactions(c, serviceFactory)
		})
		protected.GET("/transactions/:id", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transactions/:id request")
			handler.GetTransaction(c, serviceFactory)
		})

		// Beneficiary routes
		protected.GET("/beneficiaries", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/beneficiaries request")
			handler.GetBeneficiaries(c, serviceFactory)
		})
		protected.POST("/beneficiaries", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/beneficiaries create request")
			handler.CreateBeneficiary(c, serviceFactory)
		})
		protected.PUT("/beneficiaries/:id", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/beneficiaries/:id update request")
			handler.UpdateBeneficiary(c, serviceFactory)
		})
		protected.DELETE("/beneficiaries/:id", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/beneficiaries/:id delete request")
			handler.DeleteBeneficiary(c, serviceFactory)
		})
	}

	fmt.Println("All routes registered successfully")
}