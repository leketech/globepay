package service

import (
	"fmt"
	"time"

	"globepay/internal/domain"
	"globepay/internal/domain/model"
	"globepay/internal/repository"
)

// TransferServiceInterface defines the interface for transfer service
type TransferServiceInterface interface {
	GetTransfers(userID string) ([]model.Transfer, error)       // Changed from int64 to string
	GetTransferByID(transferID string) (*model.Transfer, error) // Changed from int64 to string
	CreateTransfer(transfer *model.Transfer) error
	GetExchangeRates() ([]domain.ExchangeRate, error)
	CalculateTransferFee(amount float64, fromCurrency, toCurrency string) (float64, error)
	GetTransferByReferenceNumber(referenceNumber string) (*model.Transfer, error)
	UpdateTransferStatus(transferID string, status string) error // Changed from int64 to string
}

// TransferService implements TransferServiceInterface
type TransferService struct {
	transferRepo    repository.TransferRepository
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

// NewTransferService creates a new TransferService
func NewTransferService(transferRepo repository.TransferRepository, accountRepo repository.AccountRepository, transactionRepo repository.TransactionRepository) *TransferService {
	return &TransferService{
		transferRepo:    transferRepo,
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

// GetTransfers retrieves all transfers for a user (as sender or receiver)
func (s *TransferService) GetTransfers(userID string) ([]model.Transfer, error) { // Changed from int64 to string
	// This method doesn't exist in the repository interface, so we'll use GetByUser instead
	transfers, err := s.transferRepo.GetByUser(nil, userID, 100, 0) // Pass nil context for now
	if err != nil {
		return nil, fmt.Errorf("failed to get transfers: %w", err)
	}

	// Convert []*model.Transfer to []model.Transfer
	result := make([]model.Transfer, len(transfers))
	for i, t := range transfers {
		if t != nil {
			result[i] = *t
		}
	}

	return result, nil
}

// GetTransferByID retrieves a transfer by ID
func (s *TransferService) GetTransferByID(transferID string) (*model.Transfer, error) { // Changed from int64 to string
	transfer, err := s.transferRepo.GetByID(transferID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer: %w", err)
	}

	return transfer, nil
}

// CreateTransfer creates a new transfer
func (s *TransferService) CreateTransfer(transfer *model.Transfer) error {
	// Set default values
	transfer.Status = string(domain.TransferPending)
	transfer.CreatedAt = time.Now()
	transfer.UpdatedAt = time.Now()
	transfer.ReferenceNumber = s.generateReferenceNumber()

	// For now, we'll skip the account validation since the model structure is different
	// In a real implementation, you would need to adapt this to the actual model structure

	// Create transfer
	if err := s.transferRepo.Create(transfer); err != nil {
		return fmt.Errorf("failed to create transfer: %w", err)
	}

	// Process transfer immediately
	if err := s.processTransfer(transfer); err != nil {
		return fmt.Errorf("failed to process transfer: %w", err)
	}

	return nil
}

// GetExchangeRates retrieves current exchange rates
func (s *TransferService) GetExchangeRates() ([]domain.ExchangeRate, error) {
	// In a real implementation, you would fetch this from an external service
	// or maintain it in your database
	// For now, we'll return some placeholder rates
	rates := []domain.ExchangeRate{
		{
			ID:           1,
			FromCurrency: "USD",
			ToCurrency:   "EUR",
			Rate:         0.85,
			LastUpdated:  time.Now(),
		},
		{
			ID:           2,
			FromCurrency: "USD",
			ToCurrency:   "GBP",
			Rate:         0.75,
			LastUpdated:  time.Now(),
		},
		{
			ID:           3,
			FromCurrency: "EUR",
			ToCurrency:   "USD",
			Rate:         1.18,
			LastUpdated:  time.Now(),
		},
	}

	return rates, nil
}

// CalculateTransferFee calculates the fee for a transfer
func (s *TransferService) CalculateTransferFee(amount float64, fromCurrency, toCurrency string) (float64, error) {
	// In a real implementation, you would have a more complex fee calculation
	// based on various factors like amount, currency, destination, etc.
	// For now, we'll use a simple percentage-based fee

	// Base fee of 1% of the transfer amount
	fee := amount * 0.01

	// Additional fee for international transfers
	if fromCurrency != toCurrency {
		fee += amount * 0.005 // 0.5% additional fee for currency conversion
	}

	return fee, nil
}

// GetTransferByReferenceNumber retrieves a transfer by reference number
func (s *TransferService) GetTransferByReferenceNumber(referenceNumber string) (*model.Transfer, error) {
	// This method doesn't exist in the repository interface
	// We'll need to implement it or find an alternative
	return nil, fmt.Errorf("not implemented")
}

// UpdateTransferStatus updates the status of a transfer
func (s *TransferService) UpdateTransferStatus(transferID string, status string) error { // Changed from int64 to string
	// Get existing transfer
	transfer, err := s.transferRepo.GetByID(transferID)
	if err != nil {
		return fmt.Errorf("failed to get transfer: %w", err)
	}

	// Update status
	oldStatus := domain.TransferStatus(transfer.Status)
	transfer.Status = status
	transfer.UpdatedAt = time.Now()

	// If transfer is being processed, update processed_at timestamp
	if status == string(domain.TransferProcessed) {
		now := time.Now()
		transfer.ProcessedAt = now
	}

	// Update transfer
	if err := s.transferRepo.Update(transfer); err != nil {
		return fmt.Errorf("failed to update transfer: %w", err)
	}

	// Handle status transitions
	if oldStatus != domain.TransferStatus(status) {
		if err := s.handleTransferStatusTransition(transfer, oldStatus, domain.TransferStatus(status)); err != nil {
			return fmt.Errorf("failed to handle status transition: %w", err)
		}
	}

	return nil
}

// processTransfer processes a transfer
func (s *TransferService) processTransfer(transfer *model.Transfer) error {
	// For now, we'll skip the account processing since the model structure is different
	// In a real implementation, you would need to adapt this to the actual model structure

	// Update transfer status to processed
	if err := s.UpdateTransferStatus(transfer.ID, string(domain.TransferProcessed)); err != nil {
		return fmt.Errorf("failed to update transfer status: %w", err)
	}

	return nil
}

// handleTransferStatusTransition handles logic when a transfer status changes
func (s *TransferService) handleTransferStatusTransition(transfer *model.Transfer, oldStatus, newStatus domain.TransferStatus) error {
	// Handle different status transitions
	switch {
	case oldStatus == domain.TransferPending && newStatus == domain.TransferProcessed:
		// Transfer was processed successfully
		// In a real implementation, you might want to send notifications, etc.
	case oldStatus == domain.TransferPending && newStatus == domain.TransferFailed:
		// Transfer failed
		// In a real implementation, you would reverse any partial changes
		// and notify the user
	case oldStatus == domain.TransferProcessed && newStatus == domain.TransferCancelled:
		// Transfer was cancelled after processing
		// In a real implementation, you would need to reverse the transfer
		// and handle any fees or penalties
	}

	return nil
}

// generateReferenceNumber generates a unique reference number for a transfer
func (s *TransferService) generateReferenceNumber() string {
	// In a real implementation, you would generate a unique reference number
	// based on your business requirements
	// For now, we'll generate a simple placeholder
	return fmt.Sprintf("TRF%012d", time.Now().UnixNano()%1000000000000)
}
