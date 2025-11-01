package mocks

import (
	"globepay/internal/domain"
	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) Create(transaction *domain.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) GetByID(id int64) (*domain.Transaction, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) GetByUserID(userID int64) ([]domain.Transaction, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) GetByAccountID(accountID int64) ([]domain.Transaction, error) {
	args := m.Called(accountID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) Update(transaction *domain.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}