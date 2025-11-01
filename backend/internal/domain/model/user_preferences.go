package model

import (
	"time"

	"github.com/google/uuid"
)

// UserPreferences represents user notification and security preferences
type UserPreferences struct {
	ID               string    `json:"id" db:"id"`
	UserID           string    `json:"user_id" db:"user_id"`
	EmailNotifications bool    `json:"email_notifications" db:"email_notifications"`
	PushNotifications  bool    `json:"push_notifications" db:"push_notifications"`
	SMSNotifications   bool    `json:"sms_notifications" db:"sms_notifications"`
	TransactionAlerts  bool    `json:"transaction_alerts" db:"transaction_alerts"`
	SecurityAlerts     bool    `json:"security_alerts" db:"security_alerts"`
	MarketingEmails    bool    `json:"marketing_emails" db:"marketing_emails"`
	TwoFactorEnabled   bool    `json:"two_factor_enabled" db:"two_factor_enabled"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

// NewUserPreferences creates a new user preferences instance
func NewUserPreferences(userID string) *UserPreferences {
	return &UserPreferences{
		ID:                 uuid.New().String(),
		UserID:             userID,
		EmailNotifications: true,
		PushNotifications:  false,
		SMSNotifications:   false,
		TransactionAlerts:  true,
		SecurityAlerts:     true,
		MarketingEmails:    false,
		TwoFactorEnabled:   false,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}