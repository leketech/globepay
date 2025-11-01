package utils

import (
	"fmt"
	"math/big"
	"strings"
)

// Currency represents a monetary amount with currency code
type Currency struct {
	Amount   *big.Float
	Currency string
}

// NewCurrency creates a new Currency instance
func NewCurrency(amount float64, currency string) *Currency {
	return &Currency{
		Amount:   big.NewFloat(amount),
		Currency: strings.ToUpper(currency),
	}
}

// Add adds two currency amounts
func (c *Currency) Add(other *Currency) (*Currency, error) {
	if c.Currency != other.Currency {
		return nil, fmt.Errorf("cannot add different currencies: %s and %s", c.Currency, other.Currency)
	}

	result := new(big.Float).Add(c.Amount, other.Amount)
	return &Currency{
		Amount:   result,
		Currency: c.Currency,
	}, nil
}

// Subtract subtracts one currency amount from another
func (c *Currency) Subtract(other *Currency) (*Currency, error) {
	if c.Currency != other.Currency {
		return nil, fmt.Errorf("cannot subtract different currencies: %s and %s", c.Currency, other.Currency)
	}

	result := new(big.Float).Sub(c.Amount, other.Amount)
	return &Currency{
		Amount:   result,
		Currency: c.Currency,
	}, nil
}

// Multiply multiplies a currency amount by a factor
func (c *Currency) Multiply(factor float64) *Currency {
	factorBig := big.NewFloat(factor)
	result := new(big.Float).Mul(c.Amount, factorBig)
	return &Currency{
		Amount:   result,
		Currency: c.Currency,
	}
}

// String returns a string representation of the currency
func (c *Currency) String() string {
	return fmt.Sprintf("%s %s", c.Currency, c.Amount.String())
}