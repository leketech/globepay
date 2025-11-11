package model

import (
	"time"

	"github.com/google/uuid"
)

// Beneficiary represents a transfer beneficiary
type Beneficiary struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	AccountNo   string    `json:"account_no" db:"account_no"`
	BankName    string    `json:"bank_name" db:"bank_name"`
	BankAddress string    `json:"bank_address,omitempty" db:"bank_address"`
	Country     string    `json:"country" db:"country"`
	Currency    string    `json:"currency" db:"currency"`
	SwiftCode   string    `json:"swift_code,omitempty" db:"swift_code"`
	Iban        string    `json:"iban,omitempty" db:"iban"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// NewBeneficiary creates a new beneficiary instance
func NewBeneficiary(
	userID, name, accountNo, bankName, bankAddress, country, currency, swiftCode, iban string,
) *Beneficiary {
	return &Beneficiary{
		ID:          uuid.New().String(),
		UserID:      userID,
		Name:        name,
		AccountNo:   accountNo,
		BankName:    bankName,
		BankAddress: bankAddress,
		Country:     country,
		Currency:    currency,
		SwiftCode:   swiftCode,
		Iban:        iban,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
