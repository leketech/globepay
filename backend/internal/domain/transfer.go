package domain

import (
	"time"
)

// Transfer represents a money transfer between users
type Transfer struct {
	ID              int64     `json:"id" db:"id"`
	SenderID        int64     `json:"sender_id" db:"sender_id"`
	ReceiverID      int64     `json:"receiver_id" db:"receiver_id"`
	SenderAccountID int64     `json:"sender_account_id" db:"sender_account_id"`
	ReceiverAccountID int64   `json:"receiver_account_id" db:"receiver_account_id"`
	Amount          float64   `json:"amount" db:"amount"`
	Currency        string    `json:"currency" db:"currency"`
	ExchangeRate    float64   `json:"exchange_rate" db:"exchange_rate"`
	Fee             float64   `json:"fee" db:"fee"`
	Status          string    `json:"status" db:"status"`
	Description     string    `json:"description" db:"description"`
	ReferenceNumber string    `json:"reference_number" db:"reference_number"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	ProcessedAt     time.Time `json:"processed_at" db:"processed_at"`
}

// TransferStatus represents the status of a transfer
type TransferStatus string

const (
	TransferPending   TransferStatus = "pending"
	TransferProcessed TransferStatus = "processed"
	TransferFailed    TransferStatus = "failed"
	TransferCancelled TransferStatus = "cancelled"
)

// ExchangeRate represents currency exchange rates
type ExchangeRate struct {
	ID          int64     `json:"id" db:"id"`
	FromCurrency string    `json:"from_currency" db:"from_currency"`
	ToCurrency   string    `json:"to_currency" db:"to_currency"`
	Rate        float64   `json:"rate" db:"rate"`
	LastUpdated time.Time `json:"last_updated" db:"last_updated"`
}