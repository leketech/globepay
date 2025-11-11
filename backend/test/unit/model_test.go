package unit

import (
	"globepay/internal/domain"
	"testing"
)

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

func TestAccountModel(t *testing.T) {
	// Test creating an account model
	account := &domain.Account{
		UserID:        1,
		Currency:      "USD",
		Balance:       1000.0,
		AccountNumber: "ACC001",
		AccountType:   "checking",
		Status:        "active",
	}

	if account.UserID != 1 {
		t.Errorf("Expected user ID to be 1, got %d", account.UserID)
	}

	if account.Currency != "USD" {
		t.Errorf("Expected currency to be USD, got %s", account.Currency)
	}

	if account.Balance != 1000.0 {
		t.Errorf("Expected balance to be 1000.0, got %f", account.Balance)
	}

	if account.AccountNumber != "ACC001" {
		t.Errorf("Expected account number to be ACC001, got %s", account.AccountNumber)
	}

	if account.AccountType != "checking" {
		t.Errorf("Expected account type to be checking, got %s", account.AccountType)
	}

	if account.Status != "active" {
		t.Errorf("Expected status to be active, got %s", account.Status)
	}
}

func TestTransferModel(t *testing.T) {
	// Test creating a transfer model
	transfer := &domain.Transfer{
		SenderID:          1,
		ReceiverID:        2,
		SenderAccountID:   1,
		ReceiverAccountID: 2,
		Amount:            100.0,
		Currency:          "USD",
		Fee:               1.0,
		Status:            "pending",
		ReferenceNumber:   "TRF001",
	}

	if transfer.SenderID != 1 {
		t.Errorf("Expected sender ID to be 1, got %d", transfer.SenderID)
	}

	if transfer.ReceiverID != 2 {
		t.Errorf("Expected receiver ID to be 2, got %d", transfer.ReceiverID)
	}

	if transfer.Amount != 100.0 {
		t.Errorf("Expected amount to be 100.0, got %f", transfer.Amount)
	}

	if transfer.Currency != "USD" {
		t.Errorf("Expected currency to be USD, got %s", transfer.Currency)
	}

	if transfer.Fee != 1.0 {
		t.Errorf("Expected fee to be 1.0, got %f", transfer.Fee)
	}

	if transfer.Status != "pending" {
		t.Errorf("Expected status to be pending, got %s", transfer.Status)
	}

	if transfer.ReferenceNumber != "TRF001" {
		t.Errorf("Expected reference number to be TRF001, got %s", transfer.ReferenceNumber)
	}
}
