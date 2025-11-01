package config

import (
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// Config holds application configuration
type Config struct {
	Environment     string
	ServerPort      string
	JWTSecret       string
	JWTExpiration   time.Duration
	DatabaseURL     string
	RedisURL        string
	AWSRegion       string
	LogLevel        logrus.Level
	Debug           bool
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	config := &Config{
		Environment:     getEnv("ENVIRONMENT", "development"),
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		JWTSecret:       getEnv("JWT_SECRET", "secret"),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		RedisURL:        getEnv("REDIS_URL", ""),
		AWSRegion:       getEnv("AWS_REGION", "us-east-1"),
		Debug:           getEnvAsBool("DEBUG", false),
	}

	// Parse JWT expiration
	jwtExpiration := getEnvAsInt("JWT_EXPIRATION_HOURS", 24)
	config.JWTExpiration = time.Duration(jwtExpiration) * time.Hour

	// Parse log level
	logLevel, err := logrus.ParseLevel(getEnv("LOG_LEVEL", "info"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	config.LogLevel = logLevel

	return config, nil
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		}
		return boolValue
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return intValue
	}
	return defaultValue
}