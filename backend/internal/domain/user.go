package domain

import (
	"time"
)

type User struct {
	ID            string    `json:"id" db:"id"`
	Email         string    `json:"email" db:"email"`
	Password      string    `json:"-" db:"password"` // Add this line
	PasswordHash  string    `json:"-" db:"password_hash"`
	FirstName     string    `json:"first_name" db:"first_name"`
	LastName      string    `json:"last_name" db:"last_name"`
	PhoneNumber   string    `json:"phone_number" db:"phone_number"`
	DateOfBirth   time.Time `json:"date_of_birth" db:"date_of_birth"`
	Country       string    `json:"country" db:"country"`
	KYCStatus     string    `json:"kyc_status" db:"kyc_status"`
	AccountStatus string    `json:"account_status" db:"account_status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}