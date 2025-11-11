package repository

import (
	"database/sql"
)

// RepositoryFactory creates and manages repository instances
type RepositoryFactory struct {
	db *sql.DB
}

// NewRepositoryFactory creates a new repository factory
func NewRepositoryFactory(db *sql.DB) *RepositoryFactory {
	return &RepositoryFactory{
		db: db,
	}
}

// GetUserRepository returns the user repository
func (f *RepositoryFactory) GetUserRepository() UserRepository {
	return NewUserRepository(f.db)
}

// GetAccountRepository returns the account repository
func (f *RepositoryFactory) GetAccountRepository() AccountRepository {
	return NewAccountRepository(f.db)
}

// GetTransferRepository returns the transfer repository
func (f *RepositoryFactory) GetTransferRepository() TransferRepository {
	return NewTransferRepository(f.db)
}

// GetTransactionRepository returns the transaction repository
func (f *RepositoryFactory) GetTransactionRepository() TransactionRepository {
	return NewTransactionRepository(f.db)
}

// GetBeneficiaryRepository returns the beneficiary repository
func (f *RepositoryFactory) GetBeneficiaryRepository() BeneficiaryRepository {
	return NewBeneficiaryRepository(f.db)
}

// GetCurrencyRepository returns the currency repository
func (f *RepositoryFactory) GetCurrencyRepository() CurrencyRepository {
	return NewCurrencyRepository(f.db)
}

// GetAuditRepository returns the audit repository
func (f *RepositoryFactory) GetAuditRepository() AuditRepository {
	return NewAuditRepository(f.db)
}

// GetMoneyRequestRepository returns the money request repository
func (f *RepositoryFactory) GetMoneyRequestRepository() MoneyRequestRepository {
	return NewMoneyRequestRepository(f.db)
}

// GetUserPreferencesRepository returns the user preferences repository
func (f *RepositoryFactory) GetUserPreferencesRepository() UserPreferencesRepository {
	return NewUserPreferencesRepository(f.db)
}