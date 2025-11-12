package mocks

import (
	"github.com/stretchr/testify/mock"
	"globepay/internal/domain"
)

type ExchangeRateServiceMock struct {
	mock.Mock
}

func (m *ExchangeRateServiceMock) GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	args := m.Called(fromCurrency, toCurrency)
	return args.Get(0).(float64), args.Error(1)
}

func (m *ExchangeRateServiceMock) GetAllExchangeRates() ([]domain.ExchangeRate, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.ExchangeRate), args.Error(1)
}

func (m *ExchangeRateServiceMock) UpdateExchangeRate(rate *domain.ExchangeRate) error {
	args := m.Called(rate)
	return args.Error(0)
}

func (m *ExchangeRateServiceMock) ConvertAmount(amount float64, fromCurrency, toCurrency string) (float64, error) {
	args := m.Called(amount, fromCurrency, toCurrency)
	return args.Get(0).(float64), args.Error(1)
}

func (m *ExchangeRateServiceMock) GetSupportedCurrencies() []string {
	args := m.Called()
	return args.Get(0).([]string)
}
