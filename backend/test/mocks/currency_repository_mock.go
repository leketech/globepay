package mocks

import (
	"context"
	"globepay/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type CurrencyRepositoryMock struct {
	mock.Mock
}

func (m *CurrencyRepositoryMock) GetAll(ctx context.Context) ([]*model.Currency, error) {
	args := m.Called(ctx)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.Currency), args.Error(1)
}

func (m *CurrencyRepositoryMock) GetByCode(ctx context.Context, code string) (*model.Currency, error) {
	args := m.Called(ctx, code)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Currency), args.Error(1)
}