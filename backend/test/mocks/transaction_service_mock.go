package mocks

import (
	"github.com/stretchr/testify/mock"

	"globepay/internal/domain/model"
)

type TransactionServiceMock struct {
	mock.Mock
}

func (m *TransactionServiceMock) GetTransactions(userID int64) ([]model.Transaction, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]model.Transaction), args.Error(1)
}

func (m *TransactionServiceMock) GetTransactionByID(transactionID int64) (*model.Transaction, error) {
	args := m.Called(transactionID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Transaction), args.Error(1)
}

func (m *TransactionServiceMock) CreateTransaction(transaction *model.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionServiceMock) GetTransactionHistory(userID int64, limit, offset int) ([]model.Transaction, error) {
	args := m.Called(userID, limit, offset)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]model.Transaction), args.Error(1)
}

func (m *TransactionServiceMock) GetTransactionsByStatus(status string) ([]model.Transaction, error) {
	args := m.Called(status)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]model.Transaction), args.Error(1)
}

func (m *TransactionServiceMock) UpdateTransactionStatus(transactionID int64, status string) error {
	args := m.Called(transactionID, status)
	return args.Error(0)
}
