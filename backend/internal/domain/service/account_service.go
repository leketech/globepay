package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/internal/utils"
)

// AccountService provides account-related functionality
type AccountService struct {
	accountRepo repository.AccountRepository
	userRepo    repository.UserRepository
}

// NewAccountService creates a new account service
func NewAccountService(accountRepo repository.AccountRepository, userRepo repository.UserRepository) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
		userRepo:    userRepo,
	}
}

// CreateAccount creates a new account for a user
func (s *AccountService) CreateAccount(ctx context.Context, userID, currency string) (*model.Account, error) {
	// Validate currency
	if !utils.ValidateCurrencyCode(currency) {
		return nil, &ValidationError{Field: "currency", Message: "Invalid currency code"}
	}

	// Check if user exists
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &NotFoundError{Message: "User not found"}
		}
		return nil, err
	}

	// Check if account already exists for this currency
	existingAccount, err := s.accountRepo.GetByUserAndCurrency(ctx, userID, currency)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if existingAccount != nil {
		return nil, &ConflictError{Message: fmt.Sprintf("Account for currency %s already exists", currency)}
	}

	// Generate account number
	accountNumber := generateAccountNumber()

	// Create new account
	account := model.NewAccount(userID, currency, accountNumber)

	// Save account
	if err := s.accountRepo.Create(account); err != nil {
		return nil, err
	}

	return account, nil
}

// GetAccountsByUser retrieves all accounts for a user
func (s *AccountService) GetAccountsByUser(ctx context.Context, userID string) ([]*model.Account, error) {
	return s.accountRepo.GetByUser(ctx, userID)
}

// GetAccountByID retrieves an account by ID
func (s *AccountService) GetAccountByID(_ context.Context, accountID string) (*model.Account, error) {
	return s.accountRepo.GetByID(accountID)
}

// GetAccountByNumber retrieves an account by account number
func (s *AccountService) GetAccountByNumber(_ context.Context, accountNumber string) (*model.Account, error) {
	return s.accountRepo.GetByNumber(context.Background(), accountNumber)
}

// GetAccountByUserIDAndCurrency retrieves an account by user ID and currency
func (s *AccountService) GetAccountByUserIDAndCurrency(_ context.Context, userID, currency string) (*model.Account, error) {
	return s.accountRepo.GetByUserAndCurrency(context.Background(), userID, currency)
}

// UpdateAccount updates account information
func (s *AccountService) UpdateAccount(_ context.Context, account *model.Account) error {
	account.UpdatedAt = time.Now()
	return s.accountRepo.Update(account)
}

// UpdateAccountBalance updates the balance of an account
func (s *AccountService) UpdateAccountBalance(_ context.Context, accountID string, newBalance float64) error {
	return s.accountRepo.UpdateBalance(context.Background(), accountID, newBalance)
}

// DeleteAccount deletes an account
func (s *AccountService) DeleteAccount(_ context.Context, accountID string) error {
	return s.accountRepo.Delete(accountID)
}

// generateAccountNumber generates a unique account number
func generateAccountNumber() string {
	// In a real implementation, this would generate a unique account number
	// following banking standards
	return fmt.Sprintf("ACC%s", utils.GenerateUUID()[:12])
}
