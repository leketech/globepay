package service

import (
	"database/sql"
	"strconv"

	infraconfig "globepay/internal/infrastructure/config"
	"globepay/internal/config"
	"globepay/internal/infrastructure/cache"
	"globepay/internal/infrastructure/email"
	"globepay/internal/infrastructure/sms"
	"globepay/internal/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sirupsen/logrus"
)

// Factory creates and manages application services
type Factory struct {
	config              *config.Config
	db                  *sql.DB
	redisClient         *cache.RedisClient
	awsConfig           aws.Config
	repos               *repository.RepositoryFactory
	userService         *UserService
	accountService      *AccountService
	transferService     *TransferService
	authService         *AuthService
	moneyRequestService *MoneyRequestService
}

// NewFactory creates a new service factory
func NewFactory(
	config *config.Config,
	db *sql.DB,
	redisClient *cache.RedisClient,
	awsConfig aws.Config,
) *Factory {
	factory := &Factory{
		config:      config,
		db:          db,
		redisClient: redisClient,
		awsConfig:   awsConfig,
	}

	// Initialize repository factory
	factory.repos = repository.NewRepositoryFactory(db)

	return factory
}

// GetUserService returns the user service
func (f *Factory) GetUserService() *UserService {
	if f.userService == nil {
		f.userService = NewUserService(f.repos.GetUserRepository())
		// Set the user preferences repository
		f.userService.SetUserPreferencesRepo(f.repos.GetUserPreferencesRepository())
	}
	return f.userService
}

// GetAccountService returns the account service
func (f *Factory) GetAccountService() *AccountService {
	if f.accountService == nil {
		f.accountService = NewAccountService(
			f.repos.GetAccountRepository(),
			f.repos.GetUserRepository(),
		)
	}
	return f.accountService
}

// GetTransferService returns the transfer service
func (f *Factory) GetTransferService() *TransferService {
	if f.transferService == nil {
		f.transferService = NewTransferService(
			f.repos.GetTransferRepository(),
			f.repos.GetAccountRepository(),
			f.repos.GetTransactionRepository(),
			f.repos.GetUserRepository(),
		)
	}
	return f.transferService
}

// GetTransactionService returns the transaction service
func (f *Factory) GetTransactionService() *TransactionService {
	return NewTransactionService(
		f.repos.GetTransactionRepository(),
		f.repos.GetAccountRepository(),
		f.repos.GetTransferRepository(),
	)
}

// GetBeneficiaryService returns the beneficiary service
func (f *Factory) GetBeneficiaryService() *BeneficiaryService {
	return NewBeneficiaryService(f.repos.GetBeneficiaryRepository())
}

// GetCurrencyService returns the currency service
func (f *Factory) GetCurrencyService() *CurrencyService {
	return NewCurrencyService(f.repos.GetCurrencyRepository())
}

// GetHealthService returns the health service
func (f *Factory) GetHealthService() *HealthService {
	return NewHealthService(f.db, f.redisClient)
}

// GetAuthService returns the authentication service
func (f *Factory) GetAuthService() *AuthService {
	if f.authService == nil {
		f.authService = NewAuthService(
			f.GetUserService(),
			f.config.JWT.Secret,
			f.config.JWT.Expiration,
		)
	}
	return f.authService
}

// GetNotificationService returns the notification service
func (f *Factory) GetNotificationService() *NotificationService {
	emailClient := email.NewSESClient(f.awsConfig)
	smsClient := sms.NewSNSClient(f.awsConfig)

	return NewNotificationService(emailClient, smsClient, "noreply@globepay.com")
}

// GetAuditService returns the audit service
func (f *Factory) GetAuditService() *AuditService {
	return NewAuditService(f.repos.GetAuditRepository())
}

// GetCacheService returns the cache service
func (f *Factory) GetCacheService() *CacheService {
	return NewCacheService(f.redisClient)
}

// GetConfigService returns the configuration service
func (f *Factory) GetConfigService() *ConfigService {
	return NewConfigService(&infraconfig.Config{
		Environment:   f.config.Environment,
		ServerPort:    strconv.Itoa(f.config.Server.Port),
		JWTSecret:     f.config.JWT.Secret,
		JWTExpiration: f.config.JWT.Expiration,
		DatabaseURL:   f.config.GetDatabaseDSN(),
		RedisURL:      f.config.GetRedisAddress(),
		AWSRegion:     f.config.AWS.Region,
		LogLevel:      logrus.InfoLevel, // Default value
		Debug:         f.config.IsDevelopment(),
	})
}

// GetJWTSecret returns the JWT secret from the configuration
func (f *Factory) GetJWTSecret() string {
	return f.config.JWT.Secret
}

// GetConfig returns the application configuration
func (f *Factory) GetConfig() *config.Config {
	return f.config
}

// GetMoneyRequestService returns the money request service
func (f *Factory) GetMoneyRequestService() *MoneyRequestService {
	if f.moneyRequestService == nil {
		f.moneyRequestService = NewMoneyRequestService(
			f.repos.GetMoneyRequestRepository(),
			f.repos.GetAccountRepository(),
		)
	}
	return f.moneyRequestService
}