package domain

import (
	"time"
)

type Transaction struct {
	ID                 string    `json:"id" db:"id"`
	UserID             string    `json:"user_id" db:"user_id"`
	Type               string    `json:"type" db:"type"` // TRANSFER, DEPOSIT, WITHDRAWAL
	Status             string    `json:"status" db:"status"`
	Amount             float64   `json:"amount" db:"amount"`
	Currency           string    `json:"currency" db:"currency"`
	SourceAccountID    string    `json:"source_account_id" db:"source_account_id"`
	DestAccountID      string    `json:"dest_account_id" db:"dest_account_id"`
	Fee                float64   `json:"fee" db:"fee"`
	ExchangeRate       float64   `json:"exchange_rate" db:"exchange_rate"`
	Description        string    `json:"description" db:"description"`
	Reference          string    `json:"reference" db:"reference"`
	ProcessedAt        *time.Time `json:"processed_at" db:"processed_at"`
	FailureReason      string    `json:"failure_reason,omitempty" db:"failure_reason"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}