package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
)

// CurrencyService provides currency-related functionality
type CurrencyService struct {
	currencyRepo repository.CurrencyRepository
}

// NewCurrencyService creates a new currency service
func NewCurrencyService(currencyRepo repository.CurrencyRepository) *CurrencyService {
	return &CurrencyService{
		currencyRepo: currencyRepo,
	}
}

// GetSupportedCurrencies retrieves all supported currencies
func (s *CurrencyService) GetSupportedCurrencies(ctx context.Context) ([]*model.Currency, error) {
	return s.currencyRepo.GetAll(ctx)
}

// GetCurrencyByCode retrieves a currency by code
func (s *CurrencyService) GetCurrencyByCode(ctx context.Context, code string) (*model.Currency, error) {
	return s.currencyRepo.GetByCode(ctx, code)
}

// ExchangeRateAPIResponse represents the response from ExchangeRate.host API
type ExchangeRateAPIResponse struct {
	Success bool `json:"success"`
	Query struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	} `json:"query"`
	Info struct {
		Timestamp int64   `json:"timestamp"`
		Rate      float64 `json:"rate"`
	} `json:"info"`
	Result float64 `json:"result"`
}

// GetExchangeRate retrieves the exchange rate between two currencies from ExchangeRate-API.com
func (s *CurrencyService) GetExchangeRate(ctx context.Context, fromCurrency, toCurrency string, amount float64) (*ExchangeRateResponse, error) {
	// Try to get real exchange rate from ExchangeRate-API.com (free, no API key required)
	url := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/%s", fromCurrency)
	
	resp, err := http.Get(url)
	if err != nil {
		// Fallback to mock data if API call fails
		return s.getMockExchangeRate(fromCurrency, toCurrency, amount), nil
	}
	defer resp.Body.Close()

	// Check if response is successful
	if resp.StatusCode != http.StatusOK {
		// Fallback to mock data if API returns error
		return s.getMockExchangeRate(fromCurrency, toCurrency, amount), nil
	}

	// Parse the response from ExchangeRate-API.com
	var apiResponse struct {
		Base  string             `json:"base"`
		Date  string             `json:"date"`
		Rates map[string]float64 `json:"rates"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		// Fallback to mock data if JSON parsing fails
		return s.getMockExchangeRate(fromCurrency, toCurrency, amount), nil
	}

	// Get the rate for the target currency
	rate, exists := apiResponse.Rates[toCurrency]
	if !exists {
		// Fallback to mock data if rate not found
		return s.getMockExchangeRate(fromCurrency, toCurrency, amount), nil
	}

	// Calculate fee
	fee := calculateTransferFee(amount)
	
	// Calculate converted amount
	convertedAmount := (amount - fee) * rate
	
	// Parse timestamp
	timestamp, err := time.Parse("2006-01-02", apiResponse.Date)
	if err != nil {
		timestamp = time.Now()
	}
	
	return &ExchangeRateResponse{
		FromCurrency:    fromCurrency,
		ToCurrency:      toCurrency,
		Rate:           rate,
		Fee:            fee,
		Amount:         amount,
		ConvertedAmount: convertedAmount,
		Timestamp:      timestamp,
	}, nil
}

// getMockExchangeRate provides fallback mock exchange rates
func (s *CurrencyService) getMockExchangeRate(fromCurrency, toCurrency string, amount float64) *ExchangeRateResponse {
	// In a real implementation, this would fetch real exchange rates
	// from a financial service
	rates := map[string]map[string]float64{
		"USD": {"EUR": 0.85, "GBP": 0.75, "JPY": 110.0, "NGN": 1580.0},
		"EUR": {"USD": 1.18, "GBP": 0.88, "JPY": 129.0, "NGN": 1860.0},
		"GBP": {"USD": 1.33, "EUR": 1.14, "JPY": 147.0, "NGN": 2110.0},
		"JPY": {"USD": 0.0091, "EUR": 0.0078, "GBP": 0.0068, "NGN": 14.36},
		"NGN": {"USD": 0.00063, "EUR": 0.00054, "GBP": 0.00047, "JPY": 0.0696},
	}
	
	rate := 1.0
	if fromRates, ok := rates[fromCurrency]; ok {
		if r, ok := fromRates[toCurrency]; ok {
			rate = r
		}
	}
	
	fee := calculateTransferFee(amount)
	convertedAmount := (amount - fee) * rate
	
	return &ExchangeRateResponse{
		FromCurrency:    fromCurrency,
		ToCurrency:      toCurrency,
		Rate:           rate,
		Fee:            fee,
		Amount:         amount,
		ConvertedAmount: convertedAmount,
		Timestamp:      time.Now(),
	}
}

// calculateTransferFee calculates transfer fee (simplified)
func calculateTransferFee(amount float64) float64 {
	// In a real implementation, this would use a more complex fee structure
	// based on amount, currency, destination, etc.
	
	// 2.5% fee with minimum of $1
	fee := amount * 0.025
	if fee < 1.0 {
		fee = 1.0
	}
	
	return fee
}

// ExchangeRateResponse represents the response for exchange rate requests
type ExchangeRateResponse struct {
	FromCurrency    string    `json:"fromCurrency"`
	ToCurrency      string    `json:"toCurrency"`
	Rate            float64   `json:"rate"`
	Fee             float64   `json:"fee"`
	Amount          float64   `json:"amount"`
	ConvertedAmount float64   `json:"convertedAmount"`
	Timestamp       time.Time `json:"timestamp"`
}