package service

import (
	"database/sql"

	"globepay/internal/infrastructure/cache"
	"globepay/internal/infrastructure/config"
	"globepay/internal/infrastructure/email"
	"globepay/internal/infrastructure/sms"
	"globepay/internal/repository"
	
	"github.com/aws/aws-sdk-go-v2/aws"
)

// ServiceFactory creates and manages application services
type ServiceFactory struct {
	config             *config.Config
	db                 *sql.DB
	redisClient        *cache.RedisClient
	awsConfig          aws.Config
	repos              *RepositoryFactory
	userService        *UserService
	accountService     *AccountService
	transferService    *TransferService
	authService        *AuthService
	moneyRequestService *MoneyRequestService
}

// NewServiceFactory creates a new service factory
func NewServiceFactory(
	config *config.Config,
	db *sql.DB,
	redisClient *cache.RedisClient,
	awsConfig aws.Config,
) *ServiceFactory {
	factory := &ServiceFactory{
		config:      config,
		db:          db,
		redisClient: redisClient,
		awsConfig:   awsConfig,
	}
	
	// Initialize repository factory
	factory.repos = NewRepositoryFactory(db)
	
	return factory
}

// GetUserService returns the user service
func (f *ServiceFactory) GetUserService() *UserService {
	if f.userService == nil {
		f.userService = NewUserService(f.repos.GetUserRepository())
		// Set the user preferences repository
		f.userService.SetUserPreferencesRepo(f.repos.GetUserPreferencesRepository())
	}
	return f.userService
}

// GetAccountService returns the account service
func (f *ServiceFactory) GetAccountService() *AccountService {
	if f.accountService == nil {
		f.accountService = NewAccountService(
			f.repos.GetAccountRepository(),
			f.repos.GetUserRepository(),
		)
	}
	return f.accountService
}

// GetTransferService returns the transfer service
func (f *ServiceFactory) GetTransferService() *TransferService {
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
func (f *ServiceFactory) GetTransactionService() *TransactionService {
	return NewTransactionService(
		f.repos.GetTransactionRepository(),
		f.repos.GetAccountRepository(),
		f.repos.GetTransferRepository(),
	)
}

// GetBeneficiaryService returns the beneficiary service
func (f *ServiceFactory) GetBeneficiaryService() *BeneficiaryService {
	return NewBeneficiaryService(f.repos.GetBeneficiaryRepository())
}

// GetCurrencyService returns the currency service
func (f *ServiceFactory) GetCurrencyService() *CurrencyService {
	return NewCurrencyService(f.repos.GetCurrencyRepository())
}

// GetHealthService returns the health service
func (f *ServiceFactory) GetHealthService() *HealthService {
	return NewHealthService(f.db, f.redisClient)
}

// GetAuthService returns the authentication service
func (f *ServiceFactory) GetAuthService() *AuthService {
	if f.authService == nil {
		f.authService = NewAuthService(
			f.GetUserService(),
			f.config.JWTSecret,
			f.config.JWTExpiration,
		)
	}
	return f.authService
}

// GetNotificationService returns the notification service
func (f *ServiceFactory) GetNotificationService() *NotificationService {
	emailClient := email.NewSESClient(f.awsConfig)
	smsClient := sms.NewSNSClient(f.awsConfig)
	
	return NewNotificationService(emailClient, smsClient, "noreply@globepay.com")
}

// GetAuditService returns the audit service
func (f *ServiceFactory) GetAuditService() *AuditService {
	return NewAuditService(f.repos.GetAuditRepository())
}

// GetCacheService returns the cache service
func (f *ServiceFactory) GetCacheService() *CacheService {
	return NewCacheService(f.redisClient)
}

// GetConfigService returns the configuration service
func (f *ServiceFactory) GetConfigService() *ConfigService {
	return NewConfigService(f.config)
}

// GetJWTSecret returns the JWT secret from the configuration
func (f *ServiceFactory) GetJWTSecret() string {
	return f.config.JWTSecret
}

// GetConfig returns the application configuration
func (f *ServiceFactory) GetConfig() *config.Config {
	return f.config
}

// GetMoneyRequestService returns the money request service
func (f *ServiceFactory) GetMoneyRequestService() *MoneyRequestService {
	if f.moneyRequestService == nil {
		f.moneyRequestService = NewMoneyRequestService(
			f.repos.GetMoneyRequestRepository(),
			f.repos.GetAccountRepository(),
		)
	}
	return f.moneyRequestService
}

// RepositoryFactory creates and manages repositories
type RepositoryFactory struct {
	db                   *sql.DB
	userRepo             repository.UserRepository
	accountRepo          repository.AccountRepository
	transferRepo         repository.TransferRepository
	transactionRepo      repository.TransactionRepository
	beneficiaryRepo      repository.BeneficiaryRepository
	currencyRepo         repository.CurrencyRepository
	auditRepo            repository.AuditRepository
	moneyRequestRepo     repository.MoneyRequestRepository
	userPreferencesRepo  repository.UserPreferencesRepository
}

// NewRepositoryFactory creates a new repository factory
func NewRepositoryFactory(db *sql.DB) *RepositoryFactory {
	return &RepositoryFactory{
		db: db,
	}
}

// GetUserRepository returns the user repository
func (f *RepositoryFactory) GetUserRepository() repository.UserRepository {
	if f.userRepo == nil {
		f.userRepo = repository.NewUserRepository(f.db)
	}
	return f.userRepo
}

// GetAccountRepository returns the account repository
func (f *RepositoryFactory) GetAccountRepository() repository.AccountRepository {
	if f.accountRepo == nil {
		f.accountRepo = repository.NewAccountRepository(f.db)
	}
	return f.accountRepo
}

// GetTransferRepository returns the transfer repository
func (f *RepositoryFactory) GetTransferRepository() repository.TransferRepository {
	if f.transferRepo == nil {
		f.transferRepo = repository.NewTransferRepository(f.db)
	}
	return f.transferRepo
}

// GetTransactionRepository returns the transaction repository
func (f *RepositoryFactory) GetTransactionRepository() repository.TransactionRepository {
	if f.transactionRepo == nil {
		f.transactionRepo = repository.NewTransactionRepository(f.db)
	}
	return f.transactionRepo
}

// GetBeneficiaryRepository returns the beneficiary repository
func (f *RepositoryFactory) GetBeneficiaryRepository() repository.BeneficiaryRepository {
	if f.beneficiaryRepo == nil {
		f.beneficiaryRepo = repository.NewBeneficiaryRepository(f.db)
	}
	return f.beneficiaryRepo
}

// GetCurrencyRepository returns the currency repository
func (f *RepositoryFactory) GetCurrencyRepository() repository.CurrencyRepository {
	if f.currencyRepo == nil {
		f.currencyRepo = repository.NewCurrencyRepository(f.db)
	}
	return f.currencyRepo
}

// GetAuditRepository returns the audit repository
func (f *RepositoryFactory) GetAuditRepository() repository.AuditRepository {
	if f.auditRepo == nil {
		f.auditRepo = repository.NewAuditRepository(f.db)
	}
	return f.auditRepo
}

// GetMoneyRequestRepository returns the money request repository
func (f *RepositoryFactory) GetMoneyRequestRepository() repository.MoneyRequestRepository {
	if f.moneyRequestRepo == nil {
		f.moneyRequestRepo = repository.NewMoneyRequestRepository(f.db)
	}
	return f.moneyRequestRepo
}

// GetUserPreferencesRepository returns the user preferences repository
func (f *RepositoryFactory) GetUserPreferencesRepository() repository.UserPreferencesRepository {
	if f.userPreferencesRepo == nil {
		f.userPreferencesRepo = repository.NewUserPreferencesRepository(f.db)
	}
	return f.userPreferencesRepo
}