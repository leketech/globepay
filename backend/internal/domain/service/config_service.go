package service

import (
	"time"

	"globepay/internal/infrastructure/config"
)

// ConfigService provides configuration management functionality
type ConfigService struct {
	appConfig *config.Config
}

// NewConfigService creates a new configuration service
func NewConfigService(appConfig *config.Config) *ConfigService {
	return &ConfigService{
		appConfig: appConfig,
	}
}

// GetEnvironment returns the current environment
func (s *ConfigService) GetEnvironment() string {
	return s.appConfig.Environment
}

// IsDevelopment returns true if the environment is development
func (s *ConfigService) IsDevelopment() bool {
	return s.appConfig.Environment == "development"
}

// IsProduction returns true if the environment is production
func (s *ConfigService) IsProduction() bool {
	return s.appConfig.Environment == "production"
}

// GetServerPort returns the server port
func (s *ConfigService) GetServerPort() string {
	return s.appConfig.ServerPort
}

// GetJWTSecret returns the JWT secret
func (s *ConfigService) GetJWTSecret() string {
	return s.appConfig.JWTSecret
}

// GetJWTExpiration returns the JWT expiration duration
func (s *ConfigService) GetJWTExpiration() time.Duration {
	return s.appConfig.JWTExpiration
}

// GetDatabaseURL returns the database URL
func (s *ConfigService) GetDatabaseURL() string {
	return s.appConfig.DatabaseURL
}

// GetRedisURL returns the Redis URL
func (s *ConfigService) GetRedisURL() string {
	return s.appConfig.RedisURL
}

// GetAWSRegion returns the AWS region
func (s *ConfigService) GetAWSRegion() string {
	return s.appConfig.AWSRegion
}

// IsDebug returns true if debug mode is enabled
func (s *ConfigService) IsDebug() bool {
	return s.appConfig.Debug
}

// GetLogLevel returns the log level
func (s *ConfigService) GetLogLevel() string {
	return s.appConfig.LogLevel.String()
}

// GetConfig returns the entire configuration
func (s *ConfigService) GetConfig() *config.Config {
	return s.appConfig
}
