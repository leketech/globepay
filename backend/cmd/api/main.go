package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"globepay/internal/api/router"
	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/cache"
	"globepay/internal/infrastructure/config"
	"globepay/internal/infrastructure/database"
	"globepay/internal/infrastructure/logger"
	"globepay/internal/infrastructure/metrics"

	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	logger := logger.NewLogger(cfg.LogLevel, cfg.Debug)
	logger.Info("Starting Globepay API server")

	// Initialize database
	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Redis cache
	redisClient, err := cache.NewRedisClient(cfg.RedisURL)
	if err != nil {
		logger.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Load AWS configuration
	awsConfig, err := awscfg.LoadDefaultConfig(context.TODO(), awscfg.WithRegion(cfg.AWSRegion))
	if err != nil {
		logger.Fatalf("Failed to load AWS configuration: %v", err)
	}

	// Initialize service factory
	serviceFactory := service.NewServiceFactory(cfg, db, redisClient, awsConfig)

	// Initialize metrics
	metrics := metrics.NewMetrics()

	// Create Gin engine
	gin.SetMode(gin.ReleaseMode)
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// Setup routes
	router.SetupRoutes(r, serviceFactory, metrics)

	// Start server
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	// Run server in a goroutine
	go func() {
		logger.Infof("Server starting on port %s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Info("Server exiting")
}