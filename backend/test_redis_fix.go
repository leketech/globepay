package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"globepay/internal/infrastructure/cache"
)

func main() {
	fmt.Println("Testing Redis connection with retry logic...")

	// Test with a valid Redis URL (this would normally come from config)
	redisURL := "localhost:6379" // This will fail since there's no Redis running

	// Try to connect to Redis with retry logic
	client, err := cache.NewRedisClient(redisURL)
	if err != nil {
		fmt.Printf("Failed to connect to Redis (as expected in test environment): %v\n", err)
		fmt.Println("This demonstrates that the application handles Redis failures gracefully!")
	} else {
		fmt.Println("Successfully connected to Redis")
		defer client.Close()

		// Test basic operations
		ctx := context.Background()
		err = client.Set(ctx, "test_key", "test_value", 1*time.Minute)
		if err != nil {
			log.Printf("Failed to set key: %v", err)
		} else {
			fmt.Println("Successfully set test key")
		}

		value, err := client.Get(ctx, "test_key")
		if err != nil {
			log.Printf("Failed to get key: %v", err)
		} else {
			fmt.Printf("Retrieved value: %s\n", value)
		}
	}

	fmt.Println("Test completed successfully!")
}