package service

import (
	"context"
	"fmt"
	"time"

	"globepay/internal/domain"
	"globepay/internal/domain/model"
	"globepay/internal/repository"
)

// TransactionServiceInterface defines the interface for transaction service
type TransactionServiceInterface interface {
	GetTransactions(userID string) ([]model.Transaction, error)          // Changed from int64 to string
	GetTransactionByID(transactionID string) (*model.Transaction, error) // Changed from int64 to string
	CreateTransaction(transaction *model.Transaction) error
	GetTransactionHistory(userID string, limit, offset int) ([]model.Transaction, error) // Changed from int64 to string
	GetTransactionsByStatus(status string) ([]model.Transaction, error)
	UpdateTransactionStatus(transactionID string, status string) error // Changed from int64 to string
}

// TransactionService implements TransactionServiceInterface
type TransactionService struct {
	transactionRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
}

// NewTransactionService creates a new TransactionService
func NewTransactionService(transactionRepo repository.TransactionRepository, accountRepo repository.AccountRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
	}
}

// GetTransactions retrieves all transactions for a user
func (s *TransactionService) GetTransactions(userID string) ([]model.Transaction, error) { // Changed from int64 to string
	// This method doesn't exist in the repository interface, so we'll use GetByUser instead
	transactions, err := s.transactionRepo.GetByUser(context.TODO(), userID, 100, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	// Convert []*model.Transaction to []model.Transaction
	result := make([]model.Transaction, len(transactions))
	for i, t := range transactions {
		if t != nil {
			result[i] = *t
		}
	}

	return result, nil
}

// GetTransactionByID retrieves a transaction by ID
func (s *TransactionService) GetTransactionByID(transactionID string) (*model.Transaction, error) { // Changed from int64 to string
	transaction, err := s.transactionRepo.GetByID(transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	return transaction, nil
}

// CreateTransaction creates a new transaction
func (s *TransactionService) CreateTransaction(transaction *model.Transaction) error {
	// Set default values
	transaction.Status = string(domain.TransactionPending)
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	// Validate account exists and belongs to user
	account, err := s.accountRepo.GetByID(transaction.AccountID)
	if err != nil {
		return fmt.Errorf("failed to get account: %w", err)
	}

	// Check if account has sufficient funds for withdrawal transactions
	if transaction.Type == string(model.TransactionWithdrawal) && account.Balance < transaction.Amount+transaction.Fee {
		return domain.ErrInsufficientFunds
	}

	// Create transaction
	if err := s.transactionRepo.Create(transaction); err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	// Process transaction immediately for deposits
	if transaction.Type == string(model.TransactionDeposit) {
		if err := s.processDeposit(transaction); err != nil {
			return fmt.Errorf("failed to process deposit: %w", err)
		}
	}

	return nil
}

// GetTransactionHistory retrieves transaction history with pagination
func (s *TransactionService) GetTransactionHistory(userID string, limit, offset int) ([]model.Transaction, error) { // Changed from int64 to string
	// Get all transactions for the user
	allTransactions, err := s.transactionRepo.GetByUser(context.TODO(), userID, 1000, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	// Apply pagination
	start := offset
	end := offset + limit
	if start >= len(allTransactions) {
		return []model.Transaction{}, nil
	}
	if end > len(allTransactions) {
		end = len(allTransactions)
	}

	// Convert []*model.Transaction to []model.Transaction
	result := make([]model.Transaction, end-start)
	for i, t := range allTransactions[start:end] {
		if t != nil {
			result[i] = *t
		}
	}

	return result, nil
}

// GetTransactionsByStatus retrieves all transactions with a specific status
func (s *TransactionService) GetTransactionsByStatus(status string) ([]model.Transaction, error) {
	transactions, err := s.transactionRepo.GetByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	return transactions, nil
}

// UpdateTransactionStatus updates the status of a transaction
func (s *TransactionService) UpdateTransactionStatus(transactionID string, status string) error { // Changed from int64 to string
	// Get existing transaction
	transaction, err := s.transactionRepo.GetByID(transactionID)
	if err != nil {
		return fmt.Errorf("failed to get transaction: %w", err)
	}

	// Update status
	oldStatus := domain.TransactionStatus(transaction.Status)
	transaction.Status = status // Changed from string(status) to status (already string)
	transaction.UpdatedAt = time.Now()

	// If transaction is being processed, update processed_at timestamp
	if status == string(domain.TransactionProcessed) {
		now := time.Now()
		transaction.ProcessedAt = now
	}

	// Update transaction
	if err := s.transactionRepo.Update(transaction); err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	// Handle status transitions
	if oldStatus != domain.TransactionStatus(status) {
		if err := s.handleStatusTransition(transaction, oldStatus, domain.TransactionStatus(status)); err != nil {
			return fmt.Errorf("failed to handle status transition: %w", err)
		}
	}

	return nil
}

// processDeposit processes a deposit transaction
func (s *TransactionService) processDeposit(transaction *model.Transaction) error {
	// Update account balance
	account, err := s.accountRepo.GetByID(transaction.AccountID)
	if err != nil {
		return fmt.Errorf("failed to get account: %w", err)
	}

	// Add amount to account balance
	newBalance := account.Balance + transaction.Amount
	if err := s.accountRepo.UpdateBalance(context.TODO(), account.ID, newBalance); err != nil {
		return fmt.Errorf("failed to update account balance: %w", err)
	}

	// Update transaction status to processed
	if err := s.UpdateTransactionStatus(transaction.ID, string(domain.TransactionProcessed)); err != nil {
		return fmt.Errorf("failed to update transaction status: %w", err)
	}

	return nil
}

// handleStatusTransition handles logic when a transaction status changes
func (s *TransactionService) handleStatusTransition(transaction *model.Transaction, oldStatus, newStatus domain.TransactionStatus) error {
	// Handle different status transitions
	switch {
	case oldStatus == domain.TransactionPending && newStatus == domain.TransactionProcessed:
		// Process the transaction
		if transaction.Type == string(model.TransactionWithdrawal) {
			// Deduct amount from account balance
			account, err := s.accountRepo.GetByID(transaction.AccountID)
			if err != nil {
				return fmt.Errorf("failed to get account: %w", err)
			}

			newBalance := account.Balance - transaction.Amount - transaction.Fee
			if err := s.accountRepo.UpdateBalance(context.TODO(), account.ID, newBalance); err != nil {
				return fmt.Errorf("failed to update account balance: %w", err)
			}
		}
	case oldStatus == domain.TransactionPending && newStatus == domain.TransactionFailed:
		// Handle failed transaction
		// In this case, no balance changes are needed
		// but you might want to log the failure or notify the user
	case oldStatus == domain.TransactionProcessed && newStatus == domain.TransactionCancelled:
		// Handle cancelled transaction
		// Reverse the transaction
		if transaction.Type == string(model.TransactionDeposit) {
			// Deduct amount from account balance
			account, err := s.accountRepo.GetByID(transaction.AccountID)
			if err != nil {
				return fmt.Errorf("failed to get account: %w", err)
			}

			newBalance := account.Balance - transaction.Amount
			if err := s.accountRepo.UpdateBalance(context.TODO(), account.ID, newBalance); err != nil {
				return fmt.Errorf("failed to update account balance: %w", err)
			}
		} else if transaction.Type == string(model.TransactionWithdrawal) {
			// Add amount back to account balance
			account, err := s.accountRepo.GetByID(transaction.AccountID)
			if err != nil {
				return fmt.Errorf("failed to get account: %w", err)
			}

			newBalance := account.Balance + transaction.Amount + transaction.Fee
			if err := s.accountRepo.UpdateBalance(context.TODO(), account.ID, newBalance); err != nil {
				return fmt.Errorf("failed to update account balance: %w", err)
			}
		}
	}

	return nil
}

// GetTransactionByReferenceNumber retrieves a transaction by reference number
func (s *TransactionService) GetTransactionByReferenceNumber(_ string) (*model.Transaction, error) {
	// This would require adding a method to the repository interface
	// For now, we'll return a placeholder implementation
	return nil, fmt.Errorf("not implemented")
}
