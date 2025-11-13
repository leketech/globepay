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

	// Run migrations to ensure schema is up to date
	if err := runTestMigrations(db); err != nil {
		log.Printf("Failed to run test migrations: %v", err)
		if closeErr := db.Close(); closeErr != nil {
			log.Printf("Failed to close database connection: %v", closeErr)
		}
		return nil
	}

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

// runTestMigrations runs the necessary migrations for the test database
func runTestMigrations(db *sql.DB) error {
	// Create users table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			first_name VARCHAR(100) NOT NULL,
			last_name VARCHAR(100) NOT NULL,
			phone_number VARCHAR(20),
			date_of_birth DATE,
			country_code CHAR(2),
			kyc_status VARCHAR(20) DEFAULT 'pending',
			account_status VARCHAR(20) DEFAULT 'active',
			email_verified BOOLEAN DEFAULT FALSE,
			phone_verified BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// Create accounts table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			currency_code CHAR(3) NOT NULL,
			balance DECIMAL(15,2) DEFAULT 0.00,
			account_number VARCHAR(50) UNIQUE NOT NULL,
			account_type VARCHAR(20) DEFAULT 'checking',
			status VARCHAR(20) DEFAULT 'active',
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// Create transfers table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transfers (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			recipient_name VARCHAR(255) NOT NULL,
			recipient_email VARCHAR(255),
			recipient_country CHAR(2) NOT NULL,
			recipient_bank_name VARCHAR(255) NOT NULL,
			recipient_account_number VARCHAR(100) NOT NULL,
			recipient_swift_code VARCHAR(20),
			source_currency CHAR(3) NOT NULL,
			destination_currency CHAR(3) NOT NULL,
			source_amount DECIMAL(15,2) NOT NULL,
			destination_amount DECIMAL(15,2) NOT NULL,
			exchange_rate DECIMAL(10,6) NOT NULL,
			fee_amount DECIMAL(15,2) NOT NULL,
			purpose VARCHAR(100) NOT NULL,
			status VARCHAR(20) DEFAULT 'pending',
			reference_number VARCHAR(50) UNIQUE,
			estimated_arrival TIMESTAMP WITH TIME ZONE,
			processed_at TIMESTAMP WITH TIME ZONE,
			failure_reason TEXT,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// Create transactions table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
			transfer_id UUID REFERENCES transfers(id) ON DELETE SET NULL,
			type VARCHAR(20) NOT NULL,
			status VARCHAR(20) DEFAULT 'completed',
			amount DECIMAL(15,2) NOT NULL,
			currency_code CHAR(3) NOT NULL,
			fee_amount DECIMAL(15,2) DEFAULT 0.00,
			description TEXT,
			reference_number VARCHAR(50) UNIQUE,
			processed_at TIMESTAMP WITH TIME ZONE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// Create beneficiaries table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS beneficiaries (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255),
			country CHAR(2) NOT NULL,
			bank_name VARCHAR(255) NOT NULL,
			account_number VARCHAR(100) NOT NULL,
			swift_code VARCHAR(20),
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// Create currencies table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS currencies (
			code CHAR(3) PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			symbol VARCHAR(10) NOT NULL,
			decimal_places INTEGER DEFAULT 2,
			active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// Insert default currencies
	_, err = db.Exec(`
		INSERT INTO currencies (code, name, symbol) VALUES 
		('USD', 'US Dollar', '$'),
		('EUR', 'Euro', '€'),
		('GBP', 'British Pound', '£'),
		('JPY', 'Japanese Yen', '¥')
		ON CONFLICT (code) DO NOTHING
	`)
	if err != nil {
		return err
	}

	return nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}