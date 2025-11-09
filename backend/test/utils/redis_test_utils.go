package utils

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// TestRedis represents a test Redis connection
type TestRedis struct {
	Client *redis.Client
}

// NewTestRedis creates a new test Redis connection
func NewTestRedis() *TestRedis {
	// Get Redis configuration from environment variables or use defaults
	host := getEnv("TEST_REDIS_HOST", "localhost")
	port := getEnv("TEST_REDIS_PORT", "6379")
	password := getEnv("TEST_REDIS_PASSWORD", "")
	// db := getEnv("TEST_REDIS_DB", "0")  // Comment out unused variable

	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0, // Use default DB for now
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to ping test Redis: %v", err)
		if closeErr := client.Close(); closeErr != nil {
			log.Printf("Failed to close Redis client: %v", closeErr)
		}
		return nil
	}

	log.Println("Successfully connected to test Redis")
	return &TestRedis{Client: client}
}

// Close closes the test Redis connection
func (tr *TestRedis) Close() {
	if tr != nil && tr.Client != nil {
		if err := tr.Client.Close(); err != nil {
			log.Printf("Failed to close Redis connection: %v", err)
		}
		log.Println("Test Redis connection closed")
	}
}

// Clear clears all data from Redis
func (tr *TestRedis) Clear() {
	if tr == nil || tr.Client == nil {
		return
	}

	ctx := context.Background()
	err := tr.Client.FlushAll(ctx).Err()
	if err != nil {
		log.Printf("Failed to clear Redis: %v", err)
	}
	log.Println("Test Redis cleared")
}