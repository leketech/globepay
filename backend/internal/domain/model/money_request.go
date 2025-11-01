package model

import (
	"time"
)

// MoneyRequest represents a request for money from one user to another
type MoneyRequest struct {
	ID              string    `json:"id" db:"id"`
	RequesterID     string    `json:"requester_id" db:"requester_id"`
	RecipientID     string    `json:"recipient_id" db:"recipient_id"`
	Amount          float64   `json:"amount" db:"amount"`
	Currency        string    `json:"currency" db:"currency"`
	Description     string    `json:"description" db:"description"`
	Status          string    `json:"status" db:"status"`
	PaymentLink     string    `json:"payment_link" db:"payment_link"`
	ExpiresAt       time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	PaidAt          time.Time `json:"paid_at,omitempty" db:"paid_at"`
}

// MoneyRequestStatus represents the status of a money request
type MoneyRequestStatus string

const (
	MoneyRequestPending   MoneyRequestStatus = "pending"
	MoneyRequestPaid      MoneyRequestStatus = "paid"
	MoneyRequestCancelled MoneyRequestStatus = "cancelled"
	MoneyRequestExpired   MoneyRequestStatus = "expired"
)

// NewMoneyRequest creates a new money request instance
func NewMoneyRequest(
	requesterID, recipientID string,
	amount float64, currency, description string,
) *MoneyRequest {
	return &MoneyRequest{
		RequesterID: requesterID,
		RecipientID: recipientID,
		Amount:      amount,
		Currency:    currency,
		Description: description,
		Status:      string(MoneyRequestPending),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}