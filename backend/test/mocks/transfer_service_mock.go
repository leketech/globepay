package mocks

import (
	"github.com/stretchr/testify/mock"

	"globepay/internal/domain"
)

type TransferServiceMock struct {
	mock.Mock
}

func (m *TransferServiceMock) GetTransfers(userID int64) ([]domain.Transfer, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.Transfer), args.Error(1)
}

func (m *TransferServiceMock) GetTransferByID(transferID int64) (*domain.Transfer, error) {
	args := m.Called(transferID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.Transfer), args.Error(1)
}

func (m *TransferServiceMock) CreateTransfer(transfer *domain.Transfer) error {
	args := m.Called(transfer)
	return args.Error(0)
}

func (m *TransferServiceMock) GetExchangeRates() ([]domain.ExchangeRate, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.ExchangeRate), args.Error(1)
}

func (m *TransferServiceMock) CalculateTransferFee(amount float64, fromCurrency, toCurrency string) (float64, error) {
	args := m.Called(amount, fromCurrency, toCurrency)
	return args.Get(0).(float64), args.Error(1)
}

func (m *TransferServiceMock) GetTransferByReferenceNumber(referenceNumber string) (*domain.Transfer, error) {
	args := m.Called(referenceNumber)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.Transfer), args.Error(1)
}

func (m *TransferServiceMock) UpdateTransferStatus(transferID int64, status domain.TransferStatus) error {
	args := m.Called(transferID, status)
	return args.Error(0)
}
