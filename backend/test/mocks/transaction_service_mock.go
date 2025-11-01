package mocks

import (
	"globepay/internal/domain"
	"github.com/stretchr/testify/mock"
)

type TransactionServiceMock struct {
	mock.Mock
}

func (m *TransactionServiceMock) GetTransactions(userID int64) ([]domain.Transaction, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.Transaction), args.Error(1)
}

func (m *TransactionServiceMock) GetTransactionByID(transactionID int64) (*domain.Transaction, error) {
	args := m.Called(transactionID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.Transaction), args.Error(1)
}

func (m *TransactionServiceMock) CreateTransaction(transaction *domain.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionServiceMock) GetTransactionHistory(userID int64, limit, offset int) ([]domain.Transaction, error) {
	args := m.Called(userID, limit, offset)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.Transaction), args.Error(1)
}

func (m *TransactionServiceMock) GetTransactionsByStatus(status domain.TransactionStatus) ([]domain.Transaction, error) {
	args := m.Called(status)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.Transaction), args.Error(1)
}

func (m *TransactionServiceMock) UpdateTransactionStatus(transactionID int64, status domain.TransactionStatus) error {
	args := m.Called(transactionID, status)
	return args.Error(0)
}