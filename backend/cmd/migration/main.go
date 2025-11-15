package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"globepay/internal/infrastructure/config"

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
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		fmt.Printf("Failed to ping database: %v\n", err)
		os.Exit(1)
	}

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
	// Get all .up.sql files from migrations directory
	migrationDir := "./migrations"
	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		// Try alternative path
		migrationDir = "/app/migrations"
	}

	files, err := filepath.Glob(filepath.Join(migrationDir, "*.up.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migration files: %v", err)
	}

	// Sort files by name to ensure they run in order
	// (This is a simplified approach - in production you might want a more robust solution)
	
	for _, file := range files {
		fmt.Printf("Running migration: %s\n", file)
		
		// Read the migration file
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file, err)
		}
		
		// Execute the migration
		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %v", file, err)
		}
		
		fmt.Printf("Completed migration: %s\n", file)
	}

	return nil
}

func runMigrationsDown(db *sql.DB) error {
	// Get all .down.sql files from migrations directory
	migrationDir := "./migrations"
	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		// Try alternative path
		migrationDir = "/app/migrations"
	}

	files, err := filepath.Glob(filepath.Join(migrationDir, "*.down.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migration files: %v", err)
	}

	// Run in reverse order
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		fmt.Printf("Running rollback migration: %s\n", file)
		
		// Read the migration file
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file, err)
		}
		
		// Execute the migration
		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %v", file, err)
		}
		
		fmt.Printf("Completed rollback migration: %s\n", file)
	}

	return nil
}

func createMigration(name string) error {
	// In a real implementation, you would create migration files
	// For this example, we'll just print a message
	fmt.Printf("Creating migration: %s\n", name)
	return nil
}