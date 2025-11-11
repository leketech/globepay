package domain

import (
	"time"
)

// Account represents a user's financial account
type Account struct {
	ID            int64     `json:"id" db:"id"`
	UserID        int64     `json:"user_id" db:"user_id"`
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

// AccountBalanceHistory represents the history of account balance changes
type AccountBalanceHistory struct {
	ID           int64     `json:"id" db:"id"`
	AccountID    int64     `json:"account_id" db:"account_id"`
	OldBalance   float64   `json:"old_balance" db:"old_balance"`
	NewBalance   float64   `json:"new_balance" db:"new_balance"`
	ChangeAmount float64   `json:"change_amount" db:"change_amount"`
	Reason       string    `json:"reason" db:"reason"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
