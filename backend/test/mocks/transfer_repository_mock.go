package mocks

import (
	"globepay/internal/domain"
	"github.com/stretchr/testify/mock"
)

type TransferRepositoryMock struct {
	mock.Mock
}

func (m *TransferRepositoryMock) Create(transfer *domain.Transfer) error {
	args := m.Called(transfer)
	return args.Error(0)
}

func (m *TransferRepositoryMock) GetByID(id int64) (*domain.Transfer, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.Transfer), args.Error(1)
}

func (m *TransferRepositoryMock) GetByUserID(userID int64) ([]domain.Transfer, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.Transfer), args.Error(1)
}

func (m *TransferRepositoryMock) GetByReferenceNumber(referenceNumber string) (*domain.Transfer, error) {
	args := m.Called(referenceNumber)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.Transfer), args.Error(1)
}

func (m *TransferRepositoryMock) Update(transfer *domain.Transfer) error {
	args := m.Called(transfer)
	return args.Error(0)
}

func (m *TransferRepositoryMock) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}