package service

import (
	"fmt"
	"time"

	"globepay/internal/domain"
	"globepay/internal/repository"
	"globepay/internal/domain/model"
)

// TransactionServiceInterface defines the interface for transaction service
type TransactionServiceInterface interface {
	GetTransactions(userID int64) ([]model.Transaction, error)
	GetTransactionByID(transactionID int64) (*model.Transaction, error)
	CreateTransaction(transaction *model.Transaction) error)
	GetTransactionHistory(userID int64, limit, offset int) ([]model.Transaction, error)
	GetTransactionsByStatus(status string) ([]model.Transaction, error)
	UpdateTransactionStatus(transactionID int64, status string) error
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
func (s *TransactionService) GetTransactions(userID int64) ([]model.Transaction, error) {
	transactions, err := s.transactionRepo.GetByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	return transactions, nil
}

// GetTransactionByID retrieves a transaction by ID
func (s *TransactionService) GetTransactionByID(transactionID int64) (*model.Transaction, error) {
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

	if account.UserID != transaction.UserID {
		return domain.ErrAccountNotFound
	}

	// Check if account has sufficient funds for withdrawal transactions
	if transaction.Type == string(domain.TransactionWithdrawal) && account.Balance < transaction.Amount+transaction.Fee {
		return domain.ErrInsufficientFunds
	}

	// Create transaction
	if err := s.transactionRepo.Create(transaction); err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	// Process transaction immediately for deposits
	if transaction.Type == string(domain.TransactionDeposit) {
		if err := s.processDeposit(transaction); err != nil {
			return fmt.Errorf("failed to process deposit: %w", err)
		}
	}

	return nil
}

// GetTransactionHistory retrieves transaction history with pagination
func (s *TransactionService) GetTransactionHistory(userID int64, limit, offset int) ([]model.Transaction, error) {
	// Get all transactions for the user
	allTransactions, err := s.transactionRepo.GetByUserID(userID)
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

	return allTransactions[start:end], nil
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
func (s *TransactionService) UpdateTransactionStatus(transactionID int64, status string) error {
	// Get existing transaction
	transaction, err := s.transactionRepo.GetByID(transactionID)
	if err != nil {
		return fmt.Errorf("failed to get transaction: %w", err)
	}

	// Update status
	oldStatus := domain.TransactionStatus(transaction.Status)
	transaction.Status = string(status)
	transaction.UpdatedAt = time.Now()

	// If transaction is being processed, update processed_at timestamp
	if status == domain.TransactionProcessed {
		transaction.ProcessedAt = time.Now()
	}

	// Update transaction
	if err := s.transactionRepo.Update(transaction); err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	// Handle status transitions
	if oldStatus != status {
		if err := s.handleStatusTransition(transaction, oldStatus, status); err != nil {
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
	if err := s.accountRepo.UpdateBalance(account.ID, newBalance); err != nil {
		return fmt.Errorf("failed to update account balance: %w", err)
	}

	// Update transaction status to processed
	if err := s.UpdateTransactionStatus(transaction.ID, domain.TransactionProcessed); err != nil {
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
		if transaction.Type == string(domain.TransactionWithdrawal) {
			// Deduct amount from account balance
			account, err := s.accountRepo.GetByID(transaction.AccountID)
			if err != nil {
				return fmt.Errorf("failed to get account: %w", err)
			}

			newBalance := account.Balance - transaction.Amount - transaction.Fee
			if err := s.accountRepo.UpdateBalance(account.ID, newBalance); err != nil {
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
		if transaction.Type == string(domain.TransactionDeposit) {
			// Deduct amount from account balance
			account, err := s.accountRepo.GetByID(transaction.AccountID)
			if err != nil {
				return fmt.Errorf("failed to get account: %w", err)
			}

			newBalance := account.Balance - transaction.Amount
			if err := s.accountRepo.UpdateBalance(account.ID, newBalance); err != nil {
				return fmt.Errorf("failed to update account balance: %w", err)
			}
		} else if transaction.Type == string(domain.TransactionWithdrawal) {
			// Add amount back to account balance
			account, err := s.accountRepo.GetByID(transaction.AccountID)
			if err != nil {
				return fmt.Errorf("failed to get account: %w", err)
			}

			newBalance := account.Balance + transaction.Amount + transaction.Fee
			if err := s.accountRepo.UpdateBalance(account.ID, newBalance); err != nil {
				return fmt.Errorf("failed to update account balance: %w", err)
			}
		}
	}

	return nil
}

// GetTransactionByReferenceNumber retrieves a transaction by reference number
func (s *TransactionService) GetTransactionByReferenceNumber(referenceNumber string) (*model.Transaction, error) {
	// This would require adding a method to the repository interface
	// For now, we'll return a placeholder implementation
	return nil, fmt.Errorf("not implemented")
}