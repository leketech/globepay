package service

import (
	"context"
	"fmt"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/internal/utils"
)

// TransactionService provides transaction-related functionality
type TransactionService struct {
	transactionRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
	transferRepo    repository.TransferRepository
}

// NewTransactionService creates a new transaction service
func NewTransactionService(
	transactionRepo repository.TransactionRepository,
	accountRepo repository.AccountRepository,
	transferRepo repository.TransferRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
		transferRepo:    transferRepo,
	}
}

// CreateTransaction creates a new transaction
func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *model.Transaction) error {
	// Validate currency
	if !utils.ValidateCurrencyCode(transaction.Currency) {
		return &ValidationError{Field: "currency", Message: "Invalid currency code"}
	}
	
	// Validate amount
	if transaction.Amount <= 0 {
		return &ValidationError{Field: "amount", Message: "Amount must be greater than zero"}
	}
	
	// Validate type
	validTypes := map[string]bool{
		"DEPOSIT":    true,
		"WITHDRAWAL": true,
		"TRANSFER":   true,
		"FEE":        true,
	}
	
	if !validTypes[transaction.Type] {
		return &ValidationError{Field: "type", Message: "Invalid transaction type"}
	}
	
	// Generate reference number if not provided
	if transaction.ReferenceNumber == "" {
		transaction.ReferenceNumber = generateTransactionReference()
	}
	
	// Set processed time
	transaction.ProcessedAt = time.Now()
	
	// Save transaction
	if err := s.transactionRepo.Create(transaction); err != nil {
		return err
	}
	
	return nil
}

// GetTransactionsByUser retrieves transactions for a user
func (s *TransactionService) GetTransactionsByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Transaction, error) {
	return s.transactionRepo.GetByUser(ctx, userID, limit, offset)
}

// GetTransactionByID retrieves a transaction by ID
func (s *TransactionService) GetTransactionByID(ctx context.Context, transactionID string) (*model.Transaction, error) {
	return s.transactionRepo.GetByID(transactionID)
}

// GetTransactionsByAccount retrieves transactions for an account
func (s *TransactionService) GetTransactionsByAccount(ctx context.Context, accountID string, limit, offset int) ([]*model.Transaction, error) {
	return s.transactionRepo.GetByAccount(ctx, accountID, limit, offset)
}

// GetTransactionsByTransfer retrieves transactions for a transfer
func (s *TransactionService) GetTransactionsByTransfer(ctx context.Context, transferID string) ([]*model.Transaction, error) {
	return s.transactionRepo.GetByTransfer(ctx, transferID)
}

// generateTransactionReference generates a unique transaction reference
func generateTransactionReference() string {
	// In a real implementation, this would generate a unique reference number
	// following banking standards
	return fmt.Sprintf("TXN%s", utils.GenerateUUID()[:10])
}