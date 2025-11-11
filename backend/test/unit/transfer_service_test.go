package unit

import (
	"testing"
	"time"

	"globepay/internal/domain"
	"globepay/internal/domain/model"
	"globepay/internal/service"
	"globepay/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransferService_CreateTransfer(t *testing.T) {
	// Create mock repositories
	mockTransferRepo := new(mocks.TransferRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)
	mockTransactionRepo := new(mocks.TransactionRepositoryMock)

	// Create transfer service with mock repositories
	transferService := service.NewTransferService(mockTransferRepo, mockAccountRepo, mockTransactionRepo)

	// Test data
	transfer := &model.Transfer{
		ID:                     "1",
		UserID:                 "1",
		RecipientName:          "John Doe",
		RecipientEmail:         "john@example.com",
		RecipientCountry:       "US",
		RecipientBankName:      "Test Bank",
		RecipientAccountNumber: "123456789",
		RecipientSwiftCode:     "TESTUS01",
		SourceCurrency:         "USD",
		DestCurrency:           "USD",
		SourceAmount:           100.0,
		DestAmount:             100.0,
		ExchangeRate:           1.0,
		FeeAmount:              1.0,
		Purpose:                "Test transfer",
		Status:                 string(domain.TransferPending),
	}

	// Set up mock expectations
	mockTransferRepo.On("Create", mock.AnythingOfType("*model.Transfer")).Return(nil)
	mockTransferRepo.On("GetByID", "1").Return(transfer, nil)
	mockTransferRepo.On("Update", mock.AnythingOfType("*model.Transfer")).Return(nil)

	// Call the method under test
	err := transferService.CreateTransfer(transfer)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer.ReferenceNumber)
	assert.Equal(t, string(domain.TransferProcessed), transfer.Status)
	assert.WithinDuration(t, time.Now(), transfer.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), transfer.UpdatedAt, time.Second)

	// Verify mock expectations
	mockTransferRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestTransferService_GetTransfers(t *testing.T) {
	// Create mock repositories
	mockTransferRepo := new(mocks.TransferRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)
	mockTransactionRepo := new(mocks.TransactionRepositoryMock)

	// Create transfer service with mock repositories
	transferService := service.NewTransferService(mockTransferRepo, mockAccountRepo, mockTransactionRepo)

	// Test data
	userID := "1"
	expectedTransfers := []*model.Transfer{
		{
			ID:                     "1",
			UserID:                 "1",
			RecipientName:          "John Doe",
			RecipientEmail:         "john@example.com",
			RecipientCountry:       "US",
			RecipientBankName:      "Test Bank",
			RecipientAccountNumber: "123456789",
			RecipientSwiftCode:     "TESTUS01",
			SourceCurrency:         "USD",
			DestCurrency:           "USD",
			SourceAmount:           100.0,
			DestAmount:             100.0,
			ExchangeRate:           1.0,
			FeeAmount:              1.0,
			Purpose:                "Test transfer",
			Status:                 "processed",
			ReferenceNumber:        "TRF001",
			CreatedAt:              time.Now(),
			UpdatedAt:              time.Now(),
		},
	}

	// Set up mock expectations
	mockTransferRepo.On("GetByUser", mock.Anything, userID, 100, 0).Return(expectedTransfers, nil)

	// Call the method under test
	transfers, err := transferService.GetTransfers(userID)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, transfers, 1)
	// Note: We can't directly compare since GetTransfers returns []model.Transfer but we have []*model.Transfer
	assert.Equal(t, expectedTransfers[0].ID, transfers[0].ID)

	// Verify mock expectations
	mockTransferRepo.AssertExpectations(t)
}

func TestTransferService_CalculateTransferFee(t *testing.T) {
	// Create mock repositories
	mockTransferRepo := new(mocks.TransferRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)
	mockTransactionRepo := new(mocks.TransactionRepositoryMock)

	// Create transfer service with mock repositories
	transferService := service.NewTransferService(mockTransferRepo, mockAccountRepo, mockTransactionRepo)

	// Test cases
	tests := []struct {
		name         string
		amount       float64
		fromCurrency string
		toCurrency   string
		expectedFee  float64
	}{
		{
			name:         "Same currency transfer",
			amount:       100.0,
			fromCurrency: "USD",
			toCurrency:   "USD",
			expectedFee:  1.0, // 1% of 100
		},
		{
			name:         "Cross currency transfer",
			amount:       100.0,
			fromCurrency: "USD",
			toCurrency:   "EUR",
			expectedFee:  1.5, // 1% + 0.5% of 100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fee, err := transferService.CalculateTransferFee(tt.amount, tt.fromCurrency, tt.toCurrency)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedFee, fee)
		})
	}
}
