package unit

import (
	"testing"
	"time"

	"globepay/internal/domain"
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
	transfer := &domain.Transfer{
		SenderID:          1,
		ReceiverID:        2,
		SenderAccountID:   1,
		ReceiverAccountID: 2,
		Amount:            100.0,
		Currency:          "USD",
		Fee:               1.0,
	}

	senderAccount := &domain.Account{
		ID:      1,
		UserID:  1,
		Balance: 200.0,
		Currency: "USD",
	}

	receiverAccount := &domain.Account{
		ID:      2,
		UserID:  2,
		Balance: 50.0,
		Currency: "USD",
	}

	// Set up mock expectations
	mockAccountRepo.On("GetByID", int64(1)).Return(senderAccount, nil)
	mockAccountRepo.On("GetByID", int64(2)).Return(receiverAccount, nil)
	mockTransferRepo.On("Create", mock.AnythingOfType("*domain.Transfer")).Return(nil)
	mockAccountRepo.On("UpdateBalance", int64(1), 99.0).Return(nil)
	mockAccountRepo.On("UpdateBalance", int64(2), 150.0).Return(nil)
	mockTransferRepo.On("Update", mock.AnythingOfType("*domain.Transfer")).Return(nil)
	mockTransactionRepo.On("Create", mock.AnythingOfType("*domain.Transaction")).Return(nil)

	// Call the method under test
	err := transferService.CreateTransfer(transfer)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer.ReferenceNumber)
	assert.Equal(t, "processed", transfer.Status)
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
	userID := int64(1)
	expectedTransfers := []domain.Transfer{
		{
			ID:                1,
			SenderID:          1,
			ReceiverID:        2,
			SenderAccountID:   1,
			ReceiverAccountID: 2,
			Amount:            100.0,
			Currency:          "USD",
			Fee:               1.0,
			Status:            "processed",
			ReferenceNumber:   "TRF001",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
	}

	// Set up mock expectations
	mockTransferRepo.On("GetByUserID", userID).Return(expectedTransfers, nil)

	// Call the method under test
	transfers, err := transferService.GetTransfers(userID)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, transfers, 1)
	assert.Equal(t, expectedTransfers, transfers)

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