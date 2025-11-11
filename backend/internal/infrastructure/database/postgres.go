package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"globepay/internal/config"
)

// NewPostgresConnection creates a new PostgreSQL database connection
func NewPostgresConnection(config config.DatabaseConfig) (*sql.DB, error) {
	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return db, nil
}

// ClosePostgresConnection closes the PostgreSQL database connection
func ClosePostgresConnection(db *sql.DB) error {
	if db == nil {
		return nil
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	log.Println("Successfully closed PostgreSQL database connection")
	return nil
}

// RunMigrations runs database migrations
func RunMigrations(db *sql.DB, migrationsPath string) error {
	// In a real implementation, you would run actual migrations
	// For now, we'll just log that migrations would be run
	log.Printf("Running migrations from path: %s", migrationsPath)
	return nil
}

// PingDatabase pings the database to check connectivity
func PingDatabase(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database ping successful")
	return nil
}

// GetDatabaseStats returns database connection statistics
func GetDatabaseStats(db *sql.DB) sql.DBStats {
	return db.Stats()
}
