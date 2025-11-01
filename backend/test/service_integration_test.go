package test

import (
	"testing"
	"globepay/internal/domain"
	"globepay/internal/service"
)

func TestExchangeRateService(t *testing.T) {
	// Test the exchange rate service which doesn't require database connections
	exchangeRateService := service.NewExchangeRateService()

	// Test getting exchange rate
	rate, err := exchangeRateService.GetExchangeRate("USD", "EUR")
	if err != nil {
		t.Errorf("Failed to get exchange rate: %v", err)
	}

	if rate != 0.85 {
		t.Errorf("Expected rate to be 0.85, got %f", rate)
	}

	// Test converting amount
	amount := 100.0
	converted, err := exchangeRateService.ConvertAmount(amount, "USD", "EUR")
	if err != nil {
		t.Errorf("Failed to convert amount: %v", err)
	}

	expected := 85.0 // 100 * 0.85
	if converted != expected {
		t.Errorf("Expected converted amount to be %f, got %f", expected, converted)
	}

	// Test same currency conversion
	sameCurrency, err := exchangeRateService.ConvertAmount(amount, "USD", "USD")
	if err != nil {
		t.Errorf("Failed to convert same currency: %v", err)
	}

	if sameCurrency != amount {
		t.Errorf("Expected same currency conversion to return original amount %f, got %f", amount, sameCurrency)
	}
}

func TestUserModel(t *testing.T) {
	// Test creating a user model
	user := &domain.User{
		Email:         "test@example.com",
		FirstName:     "John",
		LastName:      "Doe",
		PhoneNumber:   "+1234567890",
		Country:       "US",
		KYCStatus:     "pending",
		AccountStatus: "active",
	}

	if user.Email != "test@example.com" {
		t.Errorf("Expected email to be test@example.com, got %s", user.Email)
	}

	if user.FirstName != "John" {
		t.Errorf("Expected first name to be John, got %s", user.FirstName)
	}

	if user.LastName != "Doe" {
		t.Errorf("Expected last name to be Doe, got %s", user.LastName)
	}

	if user.KYCStatus != "pending" {
		t.Errorf("Expected KYC status to be pending, got %s", user.KYCStatus)
	}

	if user.AccountStatus != "active" {
		t.Errorf("Expected account status to be active, got %s", user.AccountStatus)
	}
}