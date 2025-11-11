package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Create a channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start worker
	log.Println("Starting Globepay worker...")

	// Worker loop
	go workerLoop()

	// Wait for interrupt signal
	<-sigChan
	log.Println("Shutting down worker...")

	fmt.Println("Worker stopped successfully")
}

func workerLoop() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Process background jobs
		processJobs()
	}
}

func processJobs() {
	log.Println("Processing background jobs...")
	// Add your background job processing logic here
}
