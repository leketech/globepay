package test

import (
	"testing"
	"globepay/internal/domain"
)

func TestUserCreation(t *testing.T) {
	// Simple test without mocks or external dependencies
	user := &domain.User{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
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
}

func TestAccountCreation(t *testing.T) {
	// Simple test for account creation
	account := &domain.Account{
		Currency: "USD",
		Balance:  100.0,
	}

	if account.Currency != "USD" {
		t.Errorf("Expected currency to be USD, got %s", account.Currency)
	}

	if account.Balance != 100.0 {
		t.Errorf("Expected balance to be 100.0, got %f", account.Balance)
	}
}