package repository

import (
	"context"
	"database/sql"
	"time"
	"globepay/internal/domain/model"
)

// UserPreferencesRepo implements UserPreferencesRepository
type UserPreferencesRepo struct {
	db *sql.DB
}

// NewUserPreferencesRepository creates a new user preferences repository
func NewUserPreferencesRepository(db *sql.DB) UserPreferencesRepository {
	return &UserPreferencesRepo{db: db}
}

// GetUserPreferences retrieves user preferences by user ID
func (r *UserPreferencesRepo) GetUserPreferences(ctx context.Context, userID string) (*model.UserPreferences, error) {
	query := `
		SELECT id, user_id, email_notifications, push_notifications, sms_notifications,
		       transaction_alerts, security_alerts, marketing_emails, two_factor_enabled,
		       created_at, updated_at
		FROM user_preferences
		WHERE user_id = $1
	`

	preferences := &model.UserPreferences{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&preferences.ID, &preferences.UserID, &preferences.EmailNotifications,
		&preferences.PushNotifications, &preferences.SMSNotifications,
		&preferences.TransactionAlerts, &preferences.SecurityAlerts,
		&preferences.MarketingEmails, &preferences.TwoFactorEnabled,
		&preferences.CreatedAt, &preferences.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Create default preferences if none exist
			defaultPrefs := model.NewUserPreferences(userID)
			err = r.CreateUserPreferences(ctx, defaultPrefs)
			if err != nil {
				return nil, err
			}
			return defaultPrefs, nil
		}
		return nil, err
	}

	return preferences, nil
}

// CreateUserPreferences creates new user preferences
func (r *UserPreferencesRepo) CreateUserPreferences(ctx context.Context, preferences *model.UserPreferences) error {
	query := `
		INSERT INTO user_preferences (
			id, user_id, email_notifications, push_notifications, sms_notifications,
			transaction_alerts, security_alerts, marketing_emails, two_factor_enabled,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.ExecContext(
		ctx, query,
		preferences.ID, preferences.UserID, preferences.EmailNotifications,
		preferences.PushNotifications, preferences.SMSNotifications,
		preferences.TransactionAlerts, preferences.SecurityAlerts,
		preferences.MarketingEmails, preferences.TwoFactorEnabled,
		preferences.CreatedAt, preferences.UpdatedAt,
	)

	return err
}

// UpdateUserPreferences updates existing user preferences
func (r *UserPreferencesRepo) UpdateUserPreferences(ctx context.Context, preferences *model.UserPreferences) error {
	query := `
		UPDATE user_preferences
		SET email_notifications = $1, push_notifications = $2, sms_notifications = $3,
		    transaction_alerts = $4, security_alerts = $5, marketing_emails = $6,
		    two_factor_enabled = $7, updated_at = $8
		WHERE user_id = $9
	`

	preferences.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(
		ctx, query,
		preferences.EmailNotifications, preferences.PushNotifications,
		preferences.SMSNotifications, preferences.TransactionAlerts,
		preferences.SecurityAlerts, preferences.MarketingEmails,
		preferences.TwoFactorEnabled, preferences.UpdatedAt,
		preferences.UserID,
	)

	return err
}