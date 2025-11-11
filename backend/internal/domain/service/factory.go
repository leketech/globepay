// Factory creates and manages application services
type Factory struct {
	config              *config.Config
	db                  *sql.DB
	redisClient         *cache.RedisClient
	awsConfig           aws.Config
	repos               *RepositoryFactory
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
	factory.repos = NewRepositoryFactory(db)

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
			f.config.JWTSecret,
			f.config.JWTExpiration,
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
	return NewConfigService(f.config)
}

// GetJWTSecret returns the JWT secret from the configuration
func (f *Factory) GetJWTSecret() string {
	return f.config.JWTSecret
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