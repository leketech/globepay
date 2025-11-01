package model

import (
	"time"
)

// Currency represents a currency
type Currency struct {
	Code      string    `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	Symbol    string    `json:"symbol" db:"symbol"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// NewCurrency creates a new currency instance
func NewCurrency(code, name, symbol string) *Currency {
	return &Currency{
		Code:      code,
		Name:      name,
		Symbol:    symbol,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}