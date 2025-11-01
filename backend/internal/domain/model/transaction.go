package model

import (
	"time"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID              string    `json:"id" db:"id"`
	UserID          string    `json:"user_id" db:"user_id"`
	AccountID       string    `json:"account_id" db:"account_id"`
	TransferID      string    `json:"transfer_id,omitempty" db:"transfer_id"`
	Type            string    `json:"type" db:"type"`
	Amount          float64   `json:"amount" db:"amount"`
	Currency        string    `json:"currency" db:"currency"`
	SourceAccountID string    `json:"source_account_id,omitempty" db:"source_account_id"`
	DestAccountID   string    `json:"dest_account_id,omitempty" db:"dest_account_id"`
	Fee             float64   `json:"fee,omitempty" db:"fee"`
	ExchangeRate    float64   `json:"exchange_rate,omitempty" db:"exchange_rate"`
	Reference       string    `json:"reference" db:"reference"`
	ReferenceNumber string    `json:"reference_number" db:"reference_number"`
	Description     string    `json:"description" db:"description"`
	Status          string    `json:"status" db:"status"`
	ProcessedAt     time.Time `json:"processed_at" db:"processed_at"`
	FailureReason   string    `json:"failure_reason,omitempty" db:"failure_reason"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// TransactionType represents the type of transaction
type TransactionType string

const (
	TransactionDeposit    TransactionType = "DEPOSIT"
	TransactionWithdrawal TransactionType = "WITHDRAWAL"
	TransactionTransfer   TransactionType = "TRANSFER"
	TransactionFee        TransactionType = "FEE"
)

// TransactionStatus represents the status of a transaction
type TransactionStatus string

const (
	TransactionPending   TransactionStatus = "PENDING"
	TransactionCompleted TransactionStatus = "COMPLETED"
	TransactionFailed    TransactionStatus = "FAILED"
	TransactionCancelled TransactionStatus = "CANCELLED"
)

// NewTransaction creates a new transaction instance
func NewTransaction(accountID, transactionType string, amount float64, currency, description string) *Transaction {
	return &Transaction{
		AccountID:   accountID,
		Type:        transactionType,
		Amount:      amount,
		Currency:    currency,
		Description: description,
		Status:      string(TransactionPending),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}