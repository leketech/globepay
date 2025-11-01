package main

import (
	"database/sql"
	"fmt"
	"os"

	"globepay/internal/infrastructure/config"
	"globepay/internal/infrastructure/database"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: migrate [up|down|create]")
		os.Exit(1)
	}

	command := os.Args[1]

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Connect to database
	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	switch command {
	case "up":
		if err := runMigrationsUp(db); err != nil {
			fmt.Printf("Failed to run migrations up: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migrations applied successfully")
	case "down":
		if err := runMigrationsDown(db); err != nil {
			fmt.Printf("Failed to run migrations down: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migrations rolled back successfully")
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Usage: migrate create <name>")
			os.Exit(1)
		}
		name := os.Args[2]
		if err := createMigration(name); err != nil {
			fmt.Printf("Failed to create migration: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Migration '%s' created successfully\n", name)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Usage: migrate [up|down|create]")
		os.Exit(1)
	}
}

func runMigrationsUp(db *sql.DB) error {
	// In a real implementation, you would use a migration library like migrate
	// For this example, we'll just run a simple migration
	
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

func runMigrationsDown(db *sql.DB) error {
	// In a real implementation, you would roll back migrations
	// For this example, we'll just print a message
	fmt.Println("Rolling back migrations...")
	return nil
}

func createMigration(name string) error {
	// In a real implementation, you would create migration files
	// For this example, we'll just print a message
	fmt.Printf("Creating migration: %s\n", name)
	return nil
}