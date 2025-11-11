package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"globepay/internal/domain/model"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) Create(account *model.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountRepositoryMock) GetByID(id string) (*model.Account, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Account), args.Error(1)
}

func (m *AccountRepositoryMock) GetByUser(ctx context.Context, userID string) ([]*model.Account, error) {
	args := m.Called(ctx, userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.Account), args.Error(1)
}

func (m *AccountRepositoryMock) GetByNumber(ctx context.Context, accountNumber string) (*model.Account, error) {
	args := m.Called(ctx, accountNumber)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Account), args.Error(1)
}

func (m *AccountRepositoryMock) GetByUserAndCurrency(ctx context.Context, userID, currency string) (*model.Account, error) {
	args := m.Called(ctx, userID, currency)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Account), args.Error(1)
}

func (m *AccountRepositoryMock) Update(account *model.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountRepositoryMock) UpdateBalance(ctx context.Context, id string, balance float64) error {
	args := m.Called(ctx, id, balance)
	return args.Error(0)
}

func (m *AccountRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *AccountRepositoryMock) GetAll() ([]model.Account, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]model.Account), args.Error(1)
}
