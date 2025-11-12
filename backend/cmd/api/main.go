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

	"globepay/internal/api/router"
	"globepay/internal/config"
	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/cache"
	"globepay/internal/infrastructure/database"
	"globepay/internal/infrastructure/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	// Check if a specific config file is specified
	configFile := os.Getenv("CONFIG_FILE")
	if configFile != "" {
		// Set the config file for viper
		config.SetConfigFile(configFile)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logLevel, err := logrus.ParseLevel(cfg.Observability.LogLevel)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger := logger.NewLogger(logLevel, cfg.IsDevelopment())
	logger.Info("Starting Globepay API server...")

	// Initialize database connection
	db, err := database.NewConnection(cfg.GetDatabaseDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Redis client
	redisClient, err := cache.NewRedisClient(cfg.GetRedisAddress())
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize AWS config
	awsConfig := aws.Config{}

	// Initialize service factory
	logger.Info("Initializing service factory...")
	serviceFactory := service.NewFactory(cfg, db, redisClient, awsConfig)
	logger.Info("Service factory initialized")

	// Initialize Gin router
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	router.SetupRoutes(r, serviceFactory, nil) // TODO: Add metrics

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

	// Start server in a goroutine
	go func() {
		logger.Info(fmt.Sprintf("Server starting on port %s...", cfg.GetServerAddress()))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-sigChan
	logger.Info("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	logger.Info("Server stopped successfully")
}
