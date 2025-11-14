package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"globepay/internal/config"
	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/cache"
	"globepay/internal/infrastructure/database"
	"globepay/internal/infrastructure/logger"
	"globepay/internal/infrastructure/metrics"
	"globepay/internal/api/router"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Parse log level
	logLevel, err := logrus.ParseLevel(cfg.Observability.LogLevel)
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	// Initialize logger
	logger := logger.NewLogger(logLevel, cfg.IsDevelopment())
	logger.Info("Starting Globepay API server...")

	// Initialize database connection
	db, err := database.NewConnection(cfg.GetDatabaseDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Redis client
	var redisClient *cache.RedisClient
	redisClient, err = cache.NewRedisClient(cfg.GetRedisAddress())
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to Redis: %v. Starting without Redis support.", err))
		// Continue without Redis - the application can still function with reduced capabilities
		redisClient = nil
	}

	// Only defer close if redisClient is not nil
	if redisClient != nil {
		defer redisClient.Close()
	}

	// Initialize metrics
	var appMetrics *metrics.Metrics
	if redisClient != nil {
		appMetrics = metrics.NewMetrics()
		logger.Info("Metrics initialized")
	} else {
		logger.Info("Metrics skipped (Redis not available)")
	}

	// Initialize AWS config
	awsConfig := aws.Config{}

	// Initialize service factory
	logger.Info("Initializing service factory...")
	serviceFactory := service.NewFactory(cfg, db, redisClient, awsConfig)
	logger.Info("Service factory initialized")

	// Initialize Gin router
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	router.SetupRoutes(r, serviceFactory, appMetrics)

	// Start server
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.GetServerAddress()),
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,  // Prevent Slowloris attacks
		ReadTimeout:       30 * time.Second,  // Maximum time to read entire request
		WriteTimeout:      30 * time.Second,  // Maximum time to write response
		IdleTimeout:       120 * time.Second, // Maximum time to keep idle connections
	}

	// Create a channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		logger.Info(fmt.Sprintf("Starting server on port %s", cfg.GetServerAddress()))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	// Wait for interrupt signal
	<-sigChan
	logger.Info("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	logger.Info("Server exited")
}