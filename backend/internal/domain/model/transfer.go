package model

import (
	"time"
)

// Transfer represents a money transfer
type Transfer struct {
	ID                     string    `json:"id" db:"id"`
	UserID                 string    `json:"userId" db:"user_id"`
	RecipientName          string    `json:"recipientName" db:"recipient_name"`
	RecipientEmail         string    `json:"recipientEmail" db:"recipient_email"`
	RecipientCountry       string    `json:"recipientCountry" db:"recipient_country"`
	RecipientBankName      string    `json:"recipientBankName" db:"recipient_bank_name"`
	RecipientAccountNumber string    `json:"recipientAccountNumber" db:"recipient_account_number"`
	RecipientSwiftCode     string    `json:"recipientSwiftCode" db:"recipient_swift_code"`
	SourceCurrency         string    `json:"sourceCurrency" db:"source_currency"`
	DestCurrency           string    `json:"destCurrency" db:"destination_currency"`
	SourceAmount           float64   `json:"sourceAmount" db:"source_amount"`
	DestAmount             float64   `json:"destAmount" db:"destination_amount"`
	ExchangeRate           float64   `json:"exchangeRate" db:"exchange_rate"`
	FeeAmount              float64   `json:"fee" db:"fee_amount"`
	Purpose                string    `json:"purpose" db:"purpose"`
	Status                 string    `json:"status" db:"status"`
	ReferenceNumber        string    `json:"referenceNumber" db:"reference_number"`
	EstimatedArrival       time.Time `json:"estimatedArrival" db:"estimated_arrival"`
	ProcessedAt            time.Time `json:"processedAt,omitempty" db:"processed_at"`
	FailureReason          string    `json:"failureReason,omitempty" db:"failure_reason"`
	CreatedAt              time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt              time.Time `json:"updatedAt" db:"updated_at"`
}

// TransferStatus represents the status of a transfer
type TransferStatus string

const (
	TransferPending   TransferStatus = "pending"
	TransferCompleted TransferStatus = "completed"
	TransferFailed    TransferStatus = "failed"
	TransferCancelled TransferStatus = "cancelled"
)

// NewTransfer creates a new transfer instance
func NewTransfer(
	userID, recipientName, recipientEmail, recipientCountry, recipientBankName, recipientAccountNumber, recipientSwiftCode, sourceCurrency, destCurrency string,
	sourceAmount, destAmount, exchangeRate, feeAmount float64,
	purpose, status, referenceNumber string,
	estimatedArrival, processedAt time.Time,
) *Transfer {
	return &Transfer{
		UserID:                 userID,
		RecipientName:          recipientName,
		RecipientEmail:         recipientEmail,
		RecipientCountry:       recipientCountry,
		RecipientBankName:      recipientBankName,
		RecipientAccountNumber: recipientAccountNumber,
		RecipientSwiftCode:     recipientSwiftCode,
		SourceCurrency:         sourceCurrency,
		DestCurrency:           destCurrency,
		SourceAmount:           sourceAmount,
		DestAmount:             destAmount,
		ExchangeRate:           exchangeRate,
		FeeAmount:              feeAmount,
		Purpose:                purpose,
		Status:                 status,
		ReferenceNumber:        referenceNumber,
		EstimatedArrival:       estimatedArrival,
		ProcessedAt:            processedAt,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}
}
