package domain

import (
	"fmt"
	"net/http"
)

// Error represents a domain error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface
func (e Error) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

// HTTPStatus returns the HTTP status code for the error
func (e Error) HTTPStatus() int {
	return e.Code
}

// Common domain errors
var (
	ErrUserNotFound = Error{
		Code:    http.StatusNotFound,
		Message: "User not found",
	}

	ErrUserAlreadyExists = Error{
		Code:    http.StatusConflict,
		Message: "User already exists",
	}

	ErrInvalidCredentials = Error{
		Code:    http.StatusUnauthorized,
		Message: "Invalid credentials",
	}

	ErrAccountNotFound = Error{
		Code:    http.StatusNotFound,
		Message: "Account not found",
	}

	ErrInsufficientFunds = Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Insufficient funds",
	}

	ErrTransferNotFound = Error{
		Code:    http.StatusNotFound,
		Message: "Transfer not found",
	}

	ErrTransactionNotFound = Error{
		Code:    http.StatusNotFound,
		Message: "Transaction not found",
	}

	ErrInvalidAmount = Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Invalid amount",
	}

	ErrInvalidCurrency = Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Invalid currency",
	}

	ErrAccountFrozen = Error{
		Code:    http.StatusForbidden,
		Message: "Account is frozen",
	}

	ErrAccountClosed = Error{
		Code:    http.StatusForbidden,
		Message: "Account is closed",
	}

	ErrTransferLimitExceeded = Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Transfer limit exceeded",
	}

	ErrInvalidReferenceNumber = Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Invalid reference number",
	}

	ErrKYCRequired = Error{
		Code:    http.StatusForbidden,
		Message: "KYC verification required",
	}

	ErrInvalidOTP = Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Invalid OTP",
	}
)