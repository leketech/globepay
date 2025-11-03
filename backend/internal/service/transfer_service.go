package service

import (
	"fmt"
	"time"

	"globepay/internal/domain"
	"globepay/internal/repository"
	"globepay/internal/domain/model"
)

// TransferServiceInterface defines the interface for transfer service
type TransferServiceInterface interface {
	GetTransfers(userID int64) ([]model.Transfer, error)
	GetTransferByID(transferID int64) (*model.Transfer, error)
	CreateTransfer(transfer *model.Transfer) error
	GetExchangeRates() ([]domain.ExchangeRate, error)
	CalculateTransferFee(amount float64, fromCurrency, toCurrency string) (float64, error)
	GetTransferByReferenceNumber(referenceNumber string) (*model.Transfer, error)
	UpdateTransferStatus(transferID int64, status string) error
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
func (s *TransferService) GetTransfers(userID int64) ([]model.Transfer, error) {
	transfers, err := s.transferRepo.GetByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfers: %w", err)
	}

	return transfers, nil
}

// GetTransferByID retrieves a transfer by ID
func (s *TransferService) GetTransferByID(transferID int64) (*model.Transfer, error) {
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

	// Validate sender account exists and belongs to sender
	senderAccount, err := s.accountRepo.GetByID(transfer.SenderAccountID)
	if err != nil {
		return fmt.Errorf("failed to get sender account: %w", err)
	}

	if senderAccount.UserID != transfer.SenderID {
		return domain.ErrAccountNotFound
	}

	// Validate receiver account exists
	receiverAccount, err := s.accountRepo.GetByID(transfer.ReceiverAccountID)
	if err != nil {
		return fmt.Errorf("failed to get receiver account: %w", err)
	}

	// Check if sender account has sufficient funds
	totalAmount := transfer.Amount + transfer.Fee
	if senderAccount.Balance < totalAmount {
		return domain.ErrInsufficientFunds
	}

	// Check if accounts are in the same currency
	if senderAccount.Currency != receiverAccount.Currency {
		// For cross-currency transfers, we would need to apply exchange rates
		// This is a simplified implementation
		return domain.ErrInvalidCurrency
	}

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
	transfer, err := s.transferRepo.GetByReferenceNumber(referenceNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer: %w", err)
	}

	return transfer, nil
}

// UpdateTransferStatus updates the status of a transfer
func (s *TransferService) UpdateTransferStatus(transferID int64, status string) error {
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
	if status == domain.TransferProcessed {
		transfer.ProcessedAt = time.Now()
	}

	// Update transfer
	if err := s.transferRepo.Update(transfer); err != nil {
		return fmt.Errorf("failed to update transfer: %w", err)
	}

	// Handle status transitions
	if oldStatus != status {
		if err := s.handleTransferStatusTransition(transfer, oldStatus, status); err != nil {
			return fmt.Errorf("failed to handle status transition: %w", err)
		}
	}

	return nil
}

// processTransfer processes a transfer
func (s *TransferService) processTransfer(transfer *model.Transfer) error {
	// Get sender and receiver accounts
	senderAccount, err := s.accountRepo.GetByID(transfer.SenderAccountID)
	if err != nil {
		return fmt.Errorf("failed to get sender account: %w", err)
	}

	receiverAccount, err := s.accountRepo.GetByID(transfer.ReceiverAccountID)
	if err != nil {
		return fmt.Errorf("failed to get receiver account: %w", err)
	}

	// Deduct amount and fee from sender account
	senderNewBalance := senderAccount.Balance - transfer.Amount - transfer.Fee
	if err := s.accountRepo.UpdateBalance(senderAccount.ID, senderNewBalance); err != nil {
		return fmt.Errorf("failed to update sender account balance: %w", err)
	}

	// Add amount to receiver account
	receiverNewBalance := receiverAccount.Balance + transfer.Amount
	if err := s.accountRepo.UpdateBalance(receiverAccount.ID, receiverNewBalance); err != nil {
		return fmt.Errorf("failed to update receiver account balance: %w", err)
	}

	// Update transfer status to processed
	if err := s.UpdateTransferStatus(transfer.ID, domain.TransferProcessed); err != nil {
		return fmt.Errorf("failed to update transfer status: %w", err)
	}

	// Create transaction records for both accounts
	senderTransaction := &domain.Transaction{
		UserID:          transfer.SenderID,
		AccountID:       transfer.SenderAccountID,
		Type:            string(domain.TransactionTransfer),
		Status:          string(domain.TransactionProcessed),
		Amount:          transfer.Amount,
		Currency:        transfer.Currency,
		Fee:             transfer.Fee,
		Description:     fmt.Sprintf("Transfer to account %s", receiverAccount.AccountNumber),
		ReferenceNumber: transfer.ReferenceNumber,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		ProcessedAt:     time.Now(),
	}

	if err := s.transactionRepo.Create(senderTransaction); err != nil {
		return fmt.Errorf("failed to create sender transaction: %w", err)
	}

	receiverTransaction := &domain.Transaction{
		UserID:          transfer.ReceiverID,
		AccountID:       transfer.ReceiverAccountID,
		Type:            string(domain.TransactionTransfer),
		Status:          string(domain.TransactionProcessed),
		Amount:          transfer.Amount,
		Currency:        transfer.Currency,
		Fee:             0, // Receiver doesn't pay fee
		Description:     fmt.Sprintf("Transfer from account %s", senderAccount.AccountNumber),
		ReferenceNumber: transfer.ReferenceNumber,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		ProcessedAt:     time.Now(),
	}

	if err := s.transactionRepo.Create(receiverTransaction); err != nil {
		return fmt.Errorf("failed to create receiver transaction: %w", err)
	}

	return nil
}

// handleTransferStatusTransition handles logic when a transfer status changes
func (s *TransferService) handleTransferStatusTransition(transfer *model.Transfer, oldStatus, newStatus domain.TransferStatus) error {
	// Handle different status transitions
	switch {
	case oldStatus == domain.TransferPending && newStatus == domain.TransferFailed:
		// Handle failed transfer - reverse the transaction
		// Get sender and receiver accounts
		senderAccount, err := s.accountRepo.GetByID(transfer.SenderAccountID)
		if err != nil {
			return fmt.Errorf("failed to get sender account: %w", err)
		}

		receiverAccount, err := s.accountRepo.GetByID(transfer.ReceiverAccountID)
		if err != nil {
			return fmt.Errorf("failed to get receiver account: %w", err)
		}

		// Add amount and fee back to sender account
		senderNewBalance := senderAccount.Balance + transfer.Amount + transfer.Fee
		if err := s.accountRepo.UpdateBalance(senderAccount.ID, senderNewBalance); err != nil {
			return fmt.Errorf("failed to update sender account balance: %w", err)
		}

		// Deduct amount from receiver account
		receiverNewBalance := receiverAccount.Balance - transfer.Amount
		if err := s.accountRepo.UpdateBalance(receiverAccount.ID, receiverNewBalance); err != nil {
			return fmt.Errorf("failed to update receiver account balance: %w", err)
		}
	}

	return nil
}

// generateReferenceNumber generates a unique reference number
func (s *TransferService) generateReferenceNumber() string {
	// In a real implementation, you would generate a unique reference number
	// based on your business requirements
	// For now, we'll generate a simple placeholder
	return fmt.Sprintf("TRF%012d", time.Now().UnixNano()%1000000000000)
}