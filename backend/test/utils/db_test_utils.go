package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// TestDB represents a test database connection
type TestDB struct {
	DB *sql.DB
}

// NewTestDB creates a new test database connection
func NewTestDB() *TestDB {
	// Get database configuration from environment variables or use defaults
	host := getEnv("TEST_DB_HOST", "localhost")
	port := getEnv("TEST_DB_PORT", "5432")
	user := getEnv("TEST_DB_USER", "postgres")
	password := getEnv("TEST_DB_PASSWORD", "postgres")
	dbname := getEnv("TEST_DB_NAME", "globepay_test")

	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to open test database connection: %v", err)
		return nil
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Printf("Failed to ping test database: %v", err)
		if closeErr := db.Close(); closeErr != nil {
			log.Printf("Failed to close database connection: %v", closeErr)
		}
		return nil
	}

	// Configure connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	log.Println("Successfully connected to test database")
	return &TestDB{DB: db}
}

// Close closes the test database connection
func (tdb *TestDB) Close() {
	if tdb != nil && tdb.DB != nil {
		if err := tdb.DB.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
		log.Println("Test database connection closed")
	}
}

// ClearTables clears all data from test tables
func (tdb *TestDB) ClearTables() {
	if tdb == nil || tdb.DB == nil {
		return
	}

	tables := []string{
		"transactions",
		"transfers",
		"accounts",
		"users",
	}

	for _, table := range tables {
		_, err := tdb.DB.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			log.Printf("Failed to clear table %s: %v", table, err)
		}
	}

	log.Println("Test tables cleared")
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
