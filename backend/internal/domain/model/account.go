package model

import (
	"time"

	"globepay/internal/utils"
)

// Account represents a user's financial account
type Account struct {
	ID            string    `json:"id" db:"id"`
	UserID        string    `json:"user_id" db:"user_id"`
	AccountNumber string    `json:"account_number" db:"account_number"`
	AccountType   string    `json:"account_type" db:"account_type"`
	Currency      string    `json:"currency" db:"currency"`
	Balance       float64   `json:"balance" db:"balance"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// AccountType represents the type of account
type AccountType string

const (
	AccountTypeChecking AccountType = "checking"
	AccountTypeSavings  AccountType = "savings"
	AccountTypeBusiness AccountType = "business"
)

// AccountStatus represents the status of an account
type AccountStatus string

const (
	AccountActive   AccountStatus = "active"
	AccountInactive AccountStatus = "inactive"
	AccountClosed   AccountStatus = "closed"
	AccountFrozen   AccountStatus = "frozen"
)

// NewAccount creates a new account instance
func NewAccount(userID, currency, accountNumber string) *Account {
	return &Account{
		ID:            utils.GenerateUUID(),
		UserID:        userID,
		AccountNumber: accountNumber,
		AccountType:   string(AccountTypeChecking),
		Currency:      currency,
		Balance:       0.0,
		Status:        string(AccountActive),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
