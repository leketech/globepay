package mocks

import (
	"context"
	"globepay/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type TransferRepositoryMock struct {
	mock.Mock
}

func (m *TransferRepositoryMock) Create(transfer *model.Transfer) error {
	args := m.Called(transfer)
	return args.Error(0)
}

func (m *TransferRepositoryMock) GetByID(id string) (*model.Transfer, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Transfer), args.Error(1)
}

func (m *TransferRepositoryMock) GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Transfer, error) {
	args := m.Called(ctx, userID, limit, offset)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.Transfer), args.Error(1)
}

func (m *TransferRepositoryMock) Update(transfer *model.Transfer) error {
	args := m.Called(transfer)
	return args.Error(0)
}

func (m *TransferRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TransferRepositoryMock) GetByNameAndUser(ctx context.Context, name, userID string) (*model.Beneficiary, error) {
	args := m.Called(ctx, name, userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Beneficiary), args.Error(1)
}