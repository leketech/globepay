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

func TestTransactionService_CreateTransaction(t *testing.T) {
	// Create mock repositories
	mockTransactionRepo := new(mocks.TransactionRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	// Create transaction service with mock repositories
	transactionService := service.NewTransactionService(mockTransactionRepo, mockAccountRepo)

	// Test data for deposit transaction
	transaction := &domain.Transaction{
		UserID:          1,
		AccountID:       1,
		Type:            string(domain.TransactionDeposit),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             0.0,
		Description:     "Test deposit",
		ReferenceNumber: "DEP001",
	}

	account := &domain.Account{
		ID:      1,
		UserID:  1,
		Balance: 200.0,
		Currency: "USD",
	}

	// Set up mock expectations
	mockAccountRepo.On("GetByID", int64(1)).Return(account, nil)
	mockTransactionRepo.On("Create", mock.AnythingOfType("*domain.Transaction")).Return(nil)
	mockAccountRepo.On("UpdateBalance", int64(1), 300.0).Return(nil)
	mockTransactionRepo.On("Update", mock.AnythingOfType("*domain.Transaction")).Return(nil)

	// Call the method under test
	err := transactionService.CreateTransaction(transaction)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "processed", transaction.Status)
	assert.WithinDuration(t, time.Now(), transaction.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), transaction.UpdatedAt, time.Second)

	// Verify mock expectations
	mockTransactionRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactions(t *testing.T) {
	// Create mock repositories
	mockTransactionRepo := new(mocks.TransactionRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	// Create transaction service with mock repositories
	transactionService := service.NewTransactionService(mockTransactionRepo, mockAccountRepo)

	// Test data
	userID := int64(1)
	expectedTransactions := []domain.Transaction{
		{
			ID:              1,
			UserID:          1,
			AccountID:       1,
			Type:            string(domain.TransactionDeposit),
			Status:          "processed",
			Amount:          100.0,
			Currency:        "USD",
			Fee:             0.0,
			Description:     "Test deposit",
			ReferenceNumber: "DEP001",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ProcessedAt:     time.Now(),
		},
	}

	// Set up mock expectations
	mockTransactionRepo.On("GetByUserID", userID).Return(expectedTransactions, nil)

	// Call the method under test
	transactions, err := transactionService.GetTransactions(userID)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, transactions, 1)
	assert.Equal(t, expectedTransactions, transactions)

	// Verify mock expectations
	mockTransactionRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionHistory(t *testing.T) {
	// Create mock repositories
	mockTransactionRepo := new(mocks.TransactionRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	// Create transaction service with mock repositories
	transactionService := service.NewTransactionService(mockTransactionRepo, mockAccountRepo)

	// Test data
	userID := int64(1)
	allTransactions := []domain.Transaction{
		{ID: 1, UserID: 1, Amount: 100.0},
		{ID: 2, UserID: 1, Amount: 200.0},
		{ID: 3, UserID: 1, Amount: 300.0},
		{ID: 4, UserID: 1, Amount: 400.0},
		{ID: 5, UserID: 1, Amount: 500.0},
	}

	// Set up mock expectations
	mockTransactionRepo.On("GetByUserID", userID).Return(allTransactions, nil)

	// Test cases
	tests := []struct {
		name     string
		limit    int
		offset   int
		expected int
	}{
		{name: "First page", limit: 2, offset: 0, expected: 2},
		{name: "Second page", limit: 2, offset: 2, expected: 2},
		{name: "Last page partial", limit: 2, offset: 4, expected: 1},
		{name: "Offset beyond data", limit: 2, offset: 10, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactions, err := transactionService.GetTransactionHistory(userID, tt.limit, tt.offset)

			// Assertions
			assert.NoError(t, err)
			assert.Len(t, transactions, tt.expected)
		})
	}

	// Verify mock expectations
	mockTransactionRepo.AssertExpectations(t)
}