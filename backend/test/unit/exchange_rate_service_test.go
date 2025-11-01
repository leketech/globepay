package unit

import (
	"testing"

	"globepay/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestExchangeRateService_GetExchangeRate(t *testing.T) {
	// Create exchange rate service
	exchangeRateService := service.NewExchangeRateService()

	// Test cases
	tests := []struct {
		name          string
		fromCurrency  string
		toCurrency    string
		expectedRate  float64
		expectError   bool
	}{
		{
			name:         "USD to EUR",
			fromCurrency: "USD",
			toCurrency:   "EUR",
			expectedRate: 0.85,
			expectError:  false,
		},
		{
			name:         "Same currency",
			fromCurrency: "USD",
			toCurrency:   "USD",
			expectedRate: 1.0,
			expectError:  false,
		},
		{
			name:         "Unsupported currency pair",
			fromCurrency: "USD",
			toCurrency:   "XYZ",
			expectedRate: 0,
			expectError:  true,
		},
		{
			name:         "Unsupported source currency",
			fromCurrency: "XYZ",
			toCurrency:   "USD",
			expectedRate: 0,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rate, err := exchangeRateService.GetExchangeRate(tt.fromCurrency, tt.toCurrency)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRate, rate)
			}
		})
	}
}

func TestExchangeRateService_ConvertAmount(t *testing.T) {
	// Create exchange rate service
	exchangeRateService := service.NewExchangeRateService()

	// Test cases
	tests := []struct {
		name          string
		amount        float64
		fromCurrency  string
		toCurrency    string
		expectedAmount float64
		expectError   bool
	}{
		{
			name:          "USD to EUR conversion",
			amount:        100.0,
			fromCurrency:  "USD",
			toCurrency:    "EUR",
			expectedAmount: 85.0, // 100 * 0.85
			expectError:   false,
		},
		{
			name:          "Same currency conversion",
			amount:        100.0,
			fromCurrency:  "USD",
			toCurrency:    "USD",
			expectedAmount: 100.0, // Same currency
			expectError:   false,
		},
		{
			name:          "Unsupported conversion",
			amount:        100.0,
			fromCurrency:  "USD",
			toCurrency:    "XYZ",
			expectedAmount: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convertedAmount, err := exchangeRateService.ConvertAmount(tt.amount, tt.fromCurrency, tt.toCurrency)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedAmount, convertedAmount)
			}
		})
	}
}

func TestExchangeRateService_GetAllExchangeRates(t *testing.T) {
	// Create exchange rate service
	exchangeRateService := service.NewExchangeRateService()

	// Get all exchange rates
	rates, err := exchangeRateService.GetAllExchangeRates()

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, rates)
	// Should have at least the default rates we initialized
	assert.True(t, len(rates) >= 15) // 3 source currencies * 5 target currencies each
}

func TestExchangeRateService_GetSupportedCurrencies(t *testing.T) {
	// Create exchange rate service
	exchangeRateService := service.NewExchangeRateService()

	// Get supported currencies
	currencies := exchangeRateService.GetSupportedCurrencies()

	// Assertions
	assert.NotEmpty(t, currencies)
	// Should contain at least our default currencies
	assert.Contains(t, currencies, "USD")
	assert.Contains(t, currencies, "EUR")
	assert.Contains(t, currencies, "GBP")
}