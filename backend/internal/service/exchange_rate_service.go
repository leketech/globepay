package service

import (
	"fmt"
	"time"

	"globepay/internal/domain"
)

// ExchangeRateServiceInterface defines the interface for exchange rate service
type ExchangeRateServiceInterface interface {
	GetExchangeRate(fromCurrency, toCurrency string) (float64, error)
	GetAllExchangeRates() ([]domain.ExchangeRate, error)
	UpdateExchangeRate(rate *domain.ExchangeRate) error
	ConvertAmount(amount float64, fromCurrency, toCurrency string) (float64, error)
	GetSupportedCurrencies() []string
}

// ExchangeRateService implements ExchangeRateServiceInterface
type ExchangeRateService struct {
	// In a real implementation, you would have a repository for exchange rates
	// For now, we'll use an in-memory store
	exchangeRates map[string]map[string]float64
}

// NewExchangeRateService creates a new ExchangeRateService
func NewExchangeRateService() *ExchangeRateService {
	// Initialize with some default exchange rates
	// In a real implementation, these would come from an external service or database
	service := &ExchangeRateService{
		exchangeRates: map[string]map[string]float64{
			"USD": {
				"EUR": 0.85,
				"GBP": 0.75,
				"JPY": 110.0,
				"CAD": 1.25,
				"AUD": 1.35,
			},
			"EUR": {
				"USD": 1.18,
				"GBP": 0.88,
				"JPY": 129.0,
				"CAD": 1.47,
				"AUD": 1.59,
			},
			"GBP": {
				"USD": 1.33,
				"EUR": 1.14,
				"JPY": 147.0,
				"CAD": 1.67,
				"AUD": 1.81,
			},
		},
	}

	return service
}

// GetExchangeRate retrieves the exchange rate between two currencies
func (s *ExchangeRateService) GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	// If both currencies are the same, return 1
	if fromCurrency == toCurrency {
		return 1.0, nil
	}

	// Check if we have the rate
	rates, exists := s.exchangeRates[fromCurrency]
	if !exists {
		return 0, fmt.Errorf("exchange rate not available for currency %s", fromCurrency)
	}

	rate, exists := rates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("exchange rate not available from %s to %s", fromCurrency, toCurrency)
	}

	return rate, nil
}

// GetAllExchangeRates retrieves all exchange rates
func (s *ExchangeRateService) GetAllExchangeRates() ([]domain.ExchangeRate, error) {
	var rates []domain.ExchangeRate
	now := time.Now()

	// Convert our in-memory store to the domain model
	for fromCurrency, toRates := range s.exchangeRates {
		for toCurrency, rate := range toRates {
			rates = append(rates, domain.ExchangeRate{
				ID:           int64(len(rates) + 1),
				FromCurrency: fromCurrency,
				ToCurrency:   toCurrency,
				Rate:         rate,
				LastUpdated:  now,
			})
		}
	}

	return rates, nil
}

// UpdateExchangeRate updates an exchange rate
func (s *ExchangeRateService) UpdateExchangeRate(rate *domain.ExchangeRate) error {
	// Initialize the map if it doesn't exist
	if s.exchangeRates[rate.FromCurrency] == nil {
		s.exchangeRates[rate.FromCurrency] = make(map[string]float64)
	}

	// Update the rate
	s.exchangeRates[rate.FromCurrency][rate.ToCurrency] = rate.Rate
	rate.LastUpdated = time.Now()

	return nil
}

// ConvertAmount converts an amount from one currency to another
func (s *ExchangeRateService) ConvertAmount(amount float64, fromCurrency, toCurrency string) (float64, error) {
	// If both currencies are the same, return the original amount
	if fromCurrency == toCurrency {
		return amount, nil
	}

	// Get the exchange rate
	rate, err := s.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return 0, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	// Convert the amount
	convertedAmount := amount * rate

	return convertedAmount, nil
}

// GetSupportedCurrencies retrieves all supported currencies
func (s *ExchangeRateService) GetSupportedCurrencies() []string {
	currencies := make([]string, 0, len(s.exchangeRates))

	// Add all source currencies
	for currency := range s.exchangeRates {
		currencies = append(currencies, currency)
	}

	// Add all target currencies
	for _, rates := range s.exchangeRates {
		for currency := range rates {
			// Check if currency is already in the slice
			found := false
			for _, c := range currencies {
				if c == currency {
					found = true
					break
				}
			}
			if !found {
				currencies = append(currencies, currency)
			}
		}
	}

	return currencies
}

// GetInverseRate calculates the inverse exchange rate
func (s *ExchangeRateService) GetInverseRate(fromCurrency, toCurrency string) (float64, error) {
	rate, err := s.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return 0, err
	}

	// Return the inverse rate
	return 1 / rate, nil
}

// GetCrossRate calculates a cross exchange rate through a common currency
func (s *ExchangeRateService) GetCrossRate(fromCurrency, toCurrency, viaCurrency string) (float64, error) {
	// Get rate from fromCurrency to viaCurrency
	rate1, err := s.GetExchangeRate(fromCurrency, viaCurrency)
	if err != nil {
		return 0, fmt.Errorf("failed to get rate from %s to %s: %w", fromCurrency, viaCurrency, err)
	}

	// Get rate from viaCurrency to toCurrency
	rate2, err := s.GetExchangeRate(viaCurrency, toCurrency)
	if err != nil {
		return 0, fmt.Errorf("failed to get rate from %s to %s: %w", viaCurrency, toCurrency, err)
	}

	// Calculate cross rate
	crossRate := rate1 * rate2

	return crossRate, nil
}

// UpdateRatesFromExternalSource updates exchange rates from an external source
func (s *ExchangeRateService) UpdateRatesFromExternalSource() error {
	// In a real implementation, you would fetch rates from an external API
	// For example, from the European Central Bank, Fixer.io, or similar services
	// This is a placeholder implementation

	// Example of how you might update rates:
	// rates, err := fetchRatesFromAPI()
	// if err != nil {
	//     return fmt.Errorf("failed to fetch rates: %w", err)
	// }
	//
	// for _, rate := range rates {
	//     s.exchangeRates[rate.FromCurrency][rate.ToCurrency] = rate.Rate
	// }

	return nil
}
