package utils

import (
	"regexp"
	"strings"
)

// ValidateEmail validates an email address
func ValidateEmail(email string) bool {
	// Basic email validation regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword validates a password
func ValidatePassword(password string) bool {
	// Password should be at least 8 characters long
	return len(password) >= 8
}

// ValidatePhoneNumber validates a phone number
func ValidatePhoneNumber(phoneNumber string) bool {
	// Basic phone number validation (allows + and digits)
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return phoneRegex.MatchString(phoneNumber)
}

// ValidateCurrencyCode validates a currency code
func ValidateCurrencyCode(currencyCode string) bool {
	// Currency code should be 3 uppercase letters
	currencyRegex := regexp.MustCompile(`^[A-Z]{3}$`)
	return currencyRegex.MatchString(strings.ToUpper(currencyCode))
}

// ValidateCountryCode validates a country code
func ValidateCountryCode(countryCode string) bool {
	// Country code should be 2 uppercase letters
	countryRegex := regexp.MustCompile(`^[A-Z]{2}$`)
	return countryRegex.MatchString(strings.ToUpper(countryCode))
}