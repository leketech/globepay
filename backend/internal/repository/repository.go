package repository

import (
	"context"
	"globepay/internal/domain/model"
)

// UserRepository defines the interface for user repository
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string) error
	GetAll() ([]model.User, error)
	GetByUserAndCurrency(ctx context.Context, userID, currency string) (*model.Account, error)
}

// AccountRepository defines the interface for account repository
type AccountRepository interface {
	Create(account *model.Account) error
	GetByID(id string) (*model.Account, error)
	GetByUser(ctx context.Context, userID string) ([]*model.Account, error)
	GetByNumber(ctx context.Context, accountNumber string) (*model.Account, error)
	GetByUserAndCurrency(ctx context.Context, userID, currency string) (*model.Account, error)
	Update(account *model.Account) error
	Delete(id string) error
	GetAll() ([]model.Account, error)
	UpdateBalance(ctx context.Context, id string, balance float64) error
}

// TransferRepository defines the interface for transfer repository
type TransferRepository interface {
	Create(transfer *model.Transfer) error
	GetByID(id string) (*model.Transfer, error)
	GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Transfer, error)
	Update(transfer *model.Transfer) error
	Delete(id string) error
	GetByNameAndUser(ctx context.Context, name, userID string) (*model.Beneficiary, error)
}

// TransactionRepository defines the interface for transaction repository
type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetByID(id string) (*model.Transaction, error)
	GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Transaction, error)
	GetByAccount(ctx context.Context, accountID string, limit, offset int) ([]*model.Transaction, error)
	GetByTransfer(ctx context.Context, transferID string) ([]*model.Transaction, error)
	Update(transaction *model.Transaction) error
	Delete(id string) error
}

// BeneficiaryRepository defines the interface for beneficiary repository
type BeneficiaryRepository interface {
	Create(beneficiary *model.Beneficiary) error
	GetByID(id string) (*model.Beneficiary, error)
	GetByUser(ctx context.Context, userID string) ([]*model.Beneficiary, error)
	GetByNameAndUser(ctx context.Context, name, userID string) (*model.Beneficiary, error)
	Update(beneficiary *model.Beneficiary) error
	Delete(id string) error
}

// CurrencyRepository defines the interface for currency repository
type CurrencyRepository interface {
	GetAll(ctx context.Context) ([]*model.Currency, error)
	GetByCode(ctx context.Context, code string) (*model.Currency, error)
}

// AuditRepository defines the interface for audit repository
type AuditRepository interface {
	Create(ctx context.Context, auditLog *model.AuditLog) error
	GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.AuditLog, error)
	GetByAction(ctx context.Context, action string, limit, offset int) ([]*model.AuditLog, error)
	GetByTable(ctx context.Context, tableName string, limit, offset int) ([]*model.AuditLog, error)
}

// MoneyRequestRepository defines the interface for money request repository
type MoneyRequestRepository interface {
	Create(ctx context.Context, request *model.MoneyRequest) error
	GetByID(ctx context.Context, id string) (*model.MoneyRequest, error)
	GetByRequester(ctx context.Context, requesterID string) ([]*model.MoneyRequest, error)
	GetByRecipient(ctx context.Context, recipientID string) ([]*model.MoneyRequest, error)
	UpdateStatus(ctx context.Context, id, status string, paidAt *string) error
	UpdatePaymentLink(ctx context.Context, id, paymentLink string) error
}

// UserPreferencesRepository defines the interface for user preferences repository
type UserPreferencesRepository interface {
	GetUserPreferences(ctx context.Context, userID string) (*model.UserPreferences, error)
	CreateUserPreferences(ctx context.Context, preferences *model.UserPreferences) error
	UpdateUserPreferences(ctx context.Context, preferences *model.UserPreferences) error
}