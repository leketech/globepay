package mocks

import (
	"context"
	"globepay/internal/domain/model"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) Create(transaction *model.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) GetByID(id string) (*model.Transaction, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Transaction, error) {
	args := m.Called(ctx, userID, limit, offset)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) GetByAccount(ctx context.Context, accountID string, limit, offset int) ([]*model.Transaction, error) {
	args := m.Called(ctx, accountID, limit, offset)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) GetByTransfer(ctx context.Context, transferID string) ([]*model.Transaction, error) {
	args := m.Called(ctx, transferID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) Update(transaction *model.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) GetByStatus(status string) ([]model.Transaction, error) {
	args := m.Called(status)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]model.Transaction), args.Error(1)
}
