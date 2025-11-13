package service

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/internal/utils"
)

// TransferService provides transfer-related functionality
type TransferService struct {
	transferRepo    repository.TransferRepository
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
}

// NewTransferService creates a new transfer service
func NewTransferService(
	transferRepo repository.TransferRepository,
	accountRepo repository.AccountRepository,
	transactionRepo repository.TransactionRepository,
	userRepo repository.UserRepository,
) *TransferService {
	return &TransferService{
		transferRepo:    transferRepo,
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
	}
}

// CreateTransfer creates a new money transfer
func (s *TransferService) CreateTransfer(ctx context.Context, transfer *model.Transfer) error {
	// Validate currencies
	if !utils.ValidateCurrencyCode(transfer.SourceCurrency) {
		return &ValidationError{Field: "sourceCurrency", Message: "Invalid source currency code"}
	}

	if !utils.ValidateCurrencyCode(transfer.DestCurrency) {
		return &ValidationError{Field: "destCurrency", Message: "Invalid destination currency code"}
	}

	// Validate country code
	if !utils.ValidateCountryCode(transfer.RecipientCountry) {
		return &ValidationError{Field: "recipientCountry", Message: "Invalid country code"}
	}

	// Validate purpose
	if transfer.Purpose == "" {
		return &ValidationError{Field: "purpose", Message: "Purpose is required"}
	}

	// Validate amounts
	if transfer.SourceAmount <= 0 {
		return &ValidationError{Field: "sourceAmount", Message: "Source amount must be greater than zero"}
	}

	// Get user accounts
	sourceAccount, err := s.accountRepo.GetByUserAndCurrency(ctx, transfer.UserID, transfer.SourceCurrency)
	if err != nil {
		if err == sql.ErrNoRows {
			return &NotFoundError{Message: fmt.Sprintf("Source account for currency %s not found", transfer.SourceCurrency)}
		}
		return err
	}

	// Check if user has sufficient funds
	if sourceAccount.Balance < transfer.SourceAmount {
		return &InsufficientFundsError{Message: "Insufficient funds in source account"}
	}

	// Generate reference number
	transfer.ReferenceNumber = generateReferenceNumber()

	// Set initial status
	transfer.Status = "pending"

	// Set estimated arrival time (24 hours from now)
	transfer.EstimatedArrival = time.Now().Add(24 * time.Hour)

	// Calculate exchange rate and fees (simplified for this example)
	transfer.ExchangeRate = calculateExchangeRate(transfer.SourceCurrency, transfer.DestCurrency)
	transfer.FeeAmount = calculateFee(transfer.SourceAmount)
	transfer.DestAmount = (transfer.SourceAmount - transfer.FeeAmount) * transfer.ExchangeRate

	// Save transfer
	if err := s.transferRepo.Create(transfer); err != nil {
		return err
	}

	return nil
}

// GetTransfersByUser retrieves transfers for a user
func (s *TransferService) GetTransfersByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Transfer, error) {
	return s.transferRepo.GetByUser(ctx, userID, limit, offset)
}

// GetTransferByID retrieves a transfer by ID
func (s *TransferService) GetTransferByID(_ context.Context, transferID string) (*model.Transfer, error) {
	return s.transferRepo.GetByID(transferID)
}

// GetTransferByReferenceNumber retrieves a transfer by reference number
func (s *TransferService) GetTransferByReferenceNumber(_ string) (*model.Transfer, error) {
	return nil, fmt.Errorf("not implemented")
}

// CancelTransfer cancels a pending transfer
func (s *TransferService) CancelTransfer(_ context.Context, transferID string) error {
	transfer, err := s.transferRepo.GetByID(transferID)
	if err != nil {
		return err
	}

	// Check if transfer can be cancelled
	if transfer.Status != "pending" {
		return &ConflictError{Message: "Only pending transfers can be cancelled"}
	}

	// Update status
	transfer.Status = "cancelled"
	transfer.UpdatedAt = time.Now()

	return s.transferRepo.Update(transfer)
}

// ProcessTransfer processes a transfer (simulated)
func (s *TransferService) ProcessTransfer(_ context.Context, transferID string) error {
	transfer, err := s.transferRepo.GetByID(transferID)
	if err != nil {
		return err
	}

	// Check if transfer can be processed
	if transfer.Status != "pending" {
		return &ConflictError{Message: "Only pending transfers can be processed"}
	}

	// Update status
	transfer.Status = "completed"
	transfer.ProcessedAt = time.Now()
	transfer.UpdatedAt = time.Now()

	// Update transfer
	if err := s.transferRepo.Update(transfer); err != nil {
		return err
	}

	return nil
}

// generateReferenceNumber generates a unique reference number
func generateReferenceNumber() string {
	// In a real implementation, this would generate a unique reference number
	// following banking standards
	return fmt.Sprintf("REF%s", utils.GenerateUUID()[:10])
}

// calculateExchangeRate calculates exchange rate (simplified)
func calculateExchangeRate(fromCurrency, toCurrency string) float64 {
	// In a real implementation, this would fetch real exchange rates
	// from a financial service
	rates := map[string]map[string]float64{
		"USD": {"EUR": 0.85, "GBP": 0.75, "JPY": 110.0},
		"EUR": {"USD": 1.18, "GBP": 0.88, "JPY": 129.0},
		"GBP": {"USD": 1.33, "EUR": 1.14, "JPY": 147.0},
		"JPY": {"USD": 0.0091, "EUR": 0.0078, "GBP": 0.0068},
	}

	if fromRates, ok := rates[fromCurrency]; ok {
		if rate, ok := fromRates[toCurrency]; ok {
			return rate
		}
	}

	// Default to 1.0 if no rate found
	return 1.0
}

// calculateFee calculates transfer fee (simplified)
func calculateFee(amount float64) float64 {
	// In a real implementation, this would use a more complex fee structure
	// based on amount, currency, destination, etc.

	// 2.5% fee with minimum of $1
	fee := amount * 0.025
	if fee < 1.0 {
		fee = 1.0
	}

	// Round to 2 decimal places
	return math.Round(fee*100) / 100
}
