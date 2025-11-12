package service

import (
	"context"
	"fmt"

	"globepay/internal/domain/model"
	"globepay/internal/infrastructure/email"
	"globepay/internal/infrastructure/sms"
)

// NotificationService provides notification functionality
type NotificationService struct {
	emailClient *email.SESClient
	smsClient   *sms.SNSClient
	fromEmail   string
}

// NewNotificationService creates a new notification service
func NewNotificationService(
	emailClient *email.SESClient,
	smsClient *sms.SNSClient,
	fromEmail string,
) *NotificationService {
	return &NotificationService{
		emailClient: emailClient,
		smsClient:   smsClient,
		fromEmail:   fromEmail,
	}
}

// SendTransferNotification sends a notification about a transfer
func (s *NotificationService) SendTransferNotification(ctx context.Context, user *model.User, transfer *model.Transfer) error {
	// Send email notification
	subject := "Transfer Confirmation"
	body := fmt.Sprintf("Dear %s,\n\nYour transfer of %s %s to %s has been %s.\n\nTransfer Details:\n- Reference Number: %s\n- Amount: %s %s\n- Recipient: %s\n- Status: %s\n\nThank you for using Globepay.\n\nBest regards,\nThe Globepay Team",
		user.FirstName, formatAmount(transfer.SourceAmount), transfer.SourceCurrency,
		transfer.RecipientName, transfer.Status, transfer.ReferenceNumber,
		formatAmount(transfer.SourceAmount), transfer.SourceCurrency,
		transfer.RecipientName, transfer.Status)

	if user.Email != "" {
		if err := s.emailClient.SendEmail(ctx, s.fromEmail, user.Email, subject, body); err != nil {
			// Log error but don't fail the transfer
			fmt.Printf("Failed to send email notification: %v\n", err)
		}
	}

	// Send SMS notification if phone number is available
	if user.PhoneNumber != "" {
		message := fmt.Sprintf("Your transfer of %s %s to %s has been %s. Ref: %s",
			formatAmount(transfer.SourceAmount), transfer.SourceCurrency,
			transfer.RecipientName, transfer.Status, transfer.ReferenceNumber)

		if err := s.smsClient.SendSMS(ctx, user.PhoneNumber, message); err != nil {
			// Log error but don't fail the transfer
			fmt.Printf("Failed to send SMS notification: %v\n", err)
		}
	}

	return nil
}

// SendWelcomeEmail sends a welcome email to a new user
func (s *NotificationService) SendWelcomeEmail(ctx context.Context, user *model.User) error {
	subject := "Welcome to Globepay!"
	body := fmt.Sprintf("Welcome to Globepay, %s!\n\nThank you for joining our platform. You can now start sending money to over 190 countries.\n\nTo get started:\n1. Complete your profile\n2. Verify your identity\n3. Add beneficiaries\n4. Start transferring money\n\nIf you have any questions, please contact our support team.\n\nBest regards,\nThe Globepay Team",
		user.FirstName)

	if user.Email != "" {
		return s.emailClient.SendEmail(ctx, s.fromEmail, user.Email, subject, body)
	}

	return nil
}

// SendPasswordResetEmail sends a password reset email
func (s *NotificationService) SendPasswordResetEmail(ctx context.Context, user *model.User, resetToken string) error {
	subject := "Password Reset Request"
	body := fmt.Sprintf("Hello %s,\n\nWe received a request to reset your password. Click the link below to reset your password:\n\nhttps://globepay.com/reset-password?token=%s\n\nIf you didn't request this, please ignore this email.\n\nThis link will expire in 1 hour.\n\nBest regards,\nThe Globepay Team",
		user.FirstName, resetToken)

	if user.Email != "" {
		return s.emailClient.SendEmail(ctx, s.fromEmail, user.Email, subject, body)
	}

	return nil
}

// formatAmount formats a monetary amount for display
func formatAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}
