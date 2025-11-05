package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
	ID            string    `json:"id" db:"id"`
	Email         string    `json:"email" db:"email"`
	PasswordHash  string    `json:"-" db:"password_hash"`
	FirstName     string    `json:"first_name" db:"first_name"`
	LastName      string    `json:"last_name" db:"last_name"`
	PhoneNumber   string    `json:"phone_number" db:"phone_number"`
	DateOfBirth   time.Time `json:"date_of_birth" db:"date_of_birth"`
	Country       string    `json:"country" db:"country_code"`
	KYCStatus     string    `json:"kyc_status" db:"kyc_status"`
	AccountStatus string    `json:"account_status" db:"account_status"`
	EmailVerified bool      `json:"email_verified" db:"email_verified"`
	PhoneVerified bool      `json:"phone_verified" db:"phone_verified"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// NewUser creates a new user instance
func NewUser(email, password, firstName, lastName string) *User {
	user := &User{
		ID:            uuid.New().String(),
		Email:         email,
		FirstName:     firstName,
		LastName:      lastName,
		KYCStatus:     "pending",
		AccountStatus: "active",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	
	// Hash the password
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err == nil {
			user.PasswordHash = string(hashedPassword)
		}
	}
	
	return user
}

// NewUserWithDetails creates a new user instance with additional details
func NewUserWithDetails(email, password, firstName, lastName, phoneNumber, country string, dateOfBirth time.Time) *User {
	user := NewUser(email, password, firstName, lastName)
	user.PhoneNumber = phoneNumber
	user.Country = country
	user.DateOfBirth = dateOfBirth
	return user
}

// SetPassword hashes and sets the user's password
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

// CheckPassword verifies a password against the stored hash
func (u *User) CheckPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}