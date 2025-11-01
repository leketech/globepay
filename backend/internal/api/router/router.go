package router

import (
	"fmt"
	"globepay/internal/api/handler"
	"globepay/internal/api/middleware"
	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/metrics"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, serviceFactory *service.ServiceFactory, metrics *metrics.Metrics) {
	fmt.Println("Setting up routes...")
	
	// Apply CORS middleware to all routes
	r.Use(middleware.CORS())
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
	
	// Transfer routes (protected by authentication)
	transferGroup := r.Group("/api/v1/transfers")
	transferGroup.Use(middleware.AuthMiddleware(serviceFactory.config.JWTSecret))
	{
		transferGroup.GET("", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers request")
			handler.GetTransfers(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/transfers endpoint")
		
		transferGroup.GET("/:id", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers/:id request")
			handler.GetTransfer(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/transfers/:id endpoint")
		
		transferGroup.POST("", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers request")
			handler.CreateTransfer(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/transfers endpoint")
		
		transferGroup.POST("/:id/cancel", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/transfers/:id/cancel request")
			handler.CancelTransfer(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/transfers/:id/cancel endpoint")
	}
	
	// ✅ This is the route the frontend is expecting at /api/exchange-rate
	r.GET("/api/exchange-rate", func(c *gin.Context) {
		fmt.Println("Handling /api/exchange-rate request")
		from := c.Query("from")
		to := c.Query("to")
		amountStr := c.Query("amount")
		
		// Set default values if not provided
		if from == "" {
			from = "USD"
		}
		if to == "" {
			to = "NGN"
		}
		if amountStr == "" {
			amountStr = "1"
		}
		
		// Convert amount to float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount parameter"})
			return
		}
		
		// Call the currency service directly
		currencyService := serviceFactory.GetCurrencyService()
		rateResp, err := currencyService.GetExchangeRate(c.Request.Context(), from, to, amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exchange rate"})
			return
		}
		
		// Return the response in the format expected by the frontend
		c.JSON(http.StatusOK, gin.H{
			"fromCurrency":    from,
			"toCurrency":      to,
			"rate":            rateResp.Rate,
			"fee":             rateResp.Fee,
			"amount":          amount,
			"convertedAmount": rateResp.ConvertedAmount,
			"timestamp":       rateResp.Timestamp.Format(time.RFC3339),
		})
	})
	fmt.Println("Registered /api/exchange-rate endpoint")
	
	// ✅ This is the route the frontend is expecting at /api/v1/exchange-rates
	r.GET("/api/v1/exchange-rates", func(c *gin.Context) {
		fmt.Println("Handling /api/v1/exchange-rates request")
		from := c.Query("from")
		to := c.Query("to")
		amountStr := c.Query("amount")
		
		// Set default values if not provided
		if from == "" {
			from = "USD"
		}
		if to == "" {
			to = "NGN"
		}
		if amountStr == "" {
			amountStr = "1"
		}
		
		// Convert amount to float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount parameter"})
			return
		}
		
		// Call the currency service directly
		currencyService := serviceFactory.GetCurrencyService()
		rateResp, err := currencyService.GetExchangeRate(c.Request.Context(), from, to, amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exchange rate"})
			return
		}
		
		// Return the response in the format expected by the frontend
		c.JSON(http.StatusOK, gin.H{
			"fromCurrency":    from,
			"toCurrency":      to,
			"rate":            rateResp.Rate,
			"fee":             rateResp.Fee,
			"amount":          amount,
			"convertedAmount": rateResp.ConvertedAmount,
			"timestamp":       rateResp.Timestamp.Format(time.RFC3339),
		})
	})
	fmt.Println("Registered /api/v1/exchange-rates endpoint")
	
	// Wallet routes (protected by authentication)
	wallet := r.Group("/api/v1/wallet")
	wallet.Use(middleware.AuthMiddleware(serviceFactory.config.JWTSecret))
	{
		walletHandler := handler.NewWalletHandler(serviceFactory, metrics)
		
		// Add money to wallet
		wallet.POST("/add", walletHandler.AddMoney)
		fmt.Println("Registered /api/v1/wallet/add endpoint")
		
		// Request money from another user
		wallet.POST("/request", walletHandler.RequestMoney)
		fmt.Println("Registered /api/v1/wallet/request endpoint")
		
		// Get money requests
		wallet.GET("/requests", walletHandler.GetMoneyRequests)
		fmt.Println("Registered /api/v1/wallet/requests endpoint")
	}
	
	// User routes (protected by authentication)
	user := r.Group("/api/v1/user")
	user.Use(middleware.AuthMiddleware(serviceFactory.config.JWTSecret))
	{
		user.GET("/profile", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/profile request")
			handler.GetUserProfile(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/user/profile endpoint")
		
		user.PUT("/profile", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/profile request")
			handler.UpdateUserProfile(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/user/profile endpoint")
		
		user.GET("/accounts", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/accounts request")
			handler.GetUserAccounts(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/user/accounts endpoint")
		
		user.POST("/accounts", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/accounts request")
			handler.CreateUserAccount(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/user/accounts endpoint")
	
		// User preferences endpoints
		user.GET("/preferences", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/preferences request")
			handler.GetUserPreferences(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/user/preferences endpoint")
		
		user.PUT("/preferences", func(c *gin.Context) {
			fmt.Println("Handling /api/v1/user/preferences request")
			handler.UpdateUserPreferences(c, serviceFactory)
		})
		fmt.Println("Registered /api/v1/user/preferences endpoint")
	}
	
	fmt.Println("All routes registered successfully")
}