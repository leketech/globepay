package email

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

// SESClient wraps the AWS SES client
type SESClient struct {
	client *ses.Client
	config aws.Config
}

// NewSESClient creates a new SES client
func NewSESClient(config aws.Config) *SESClient {
	return &SESClient{
		client: ses.NewFromConfig(config),
		config: config,
	}
}

// SendEmail sends an email using AWS SES
func (s *SESClient) SendEmail(ctx context.Context, from, to, subject, body string) error {
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Message: &types.Message{
			Body: &types.Body{
				Text: &types.Content{
					Data: aws.String(body),
				},
			},
			Subject: &types.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(from),
	}

	_, err := s.client.SendEmail(ctx, input)
	return err
}

// SendRawEmail sends a raw email using AWS SES
func (s *SESClient) SendRawEmail(ctx context.Context, from string, to []string, rawMessage []byte) (string, error) {
	input := &ses.SendRawEmailInput{
		Source:       aws.String(from),
		Destinations: to,
		RawMessage: &types.RawMessage{
			Data: rawMessage,
		},
	}

	result, err := s.client.SendRawEmail(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to send raw email: %w", err)
	}

	log.Printf("Successfully sent raw email with message ID %s", *result.MessageId)
	return *result.MessageId, nil
}

// VerifyEmailIdentity verifies an email address for sending
func (s *SESClient) VerifyEmailIdentity(ctx context.Context, emailAddress string) error {
	input := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String(emailAddress),
	}

	_, err := s.client.VerifyEmailIdentity(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to verify email identity %s: %w", emailAddress, err)
	}

	log.Printf("Successfully verified email identity %s", emailAddress)
	return nil
}

// Close closes the SES client
func (s *SESClient) Close() error {
	// SES client doesn't need explicit closing
	log.Println("SES client closed")
	return nil
}

// SendTemplatedEmail sends an email using a template
func (s *SESClient) SendTemplatedEmail(ctx context.Context, from string, to []string, templateName string, templateData string) (string, error) {
	input := &ses.SendTemplatedEmailInput{
		Source:               aws.String(from),
		Destination:          &types.Destination{ToAddresses: to},
		Template:             aws.String(templateName),
		TemplateData:         aws.String(templateData),
	}

	result, err := s.client.SendTemplatedEmail(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to send templated email: %w", err)
	}

	log.Printf("Successfully sent templated email with message ID %s", *result.MessageId)
	return *result.MessageId, nil
}

// CreateTemplate creates an email template
func (s *SESClient) CreateTemplate(ctx context.Context, templateName, subjectPart, htmlPart, textPart string) error {
	template := &types.Template{
		TemplateName: aws.String(templateName),
		SubjectPart:  aws.String(subjectPart),
		HtmlPart:     aws.String(htmlPart),
		TextPart:     aws.String(textPart),
	}

	input := &ses.CreateTemplateInput{
		Template: template,
	}

	_, err := s.client.CreateTemplate(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create template %s: %w", templateName, err)
	}

	log.Printf("Successfully created template %s", templateName)
	return nil
}

// DeleteTemplate deletes an email template
func (s *SESClient) DeleteTemplate(ctx context.Context, templateName string) error {
	input := &ses.DeleteTemplateInput{
		TemplateName: aws.String(templateName),
	}

	_, err := s.client.DeleteTemplate(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete template %s: %w", templateName, err)
	}

	log.Printf("Successfully deleted template %s", templateName)
	return nil
}

// ListTemplates lists all email templates
func (s *SESClient) ListTemplates(ctx context.Context) ([]types.TemplateMetadata, error) {
	input := &ses.ListTemplatesInput{}

	result, err := s.client.ListTemplates(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to list templates: %w", err)
	}

	log.Printf("Successfully listed %d templates", len(result.TemplatesMetadata))
	return result.TemplatesMetadata, nil
}

// GetSendQuota retrieves the send quota for the SES account
func (s *SESClient) GetSendQuota(ctx context.Context) (*ses.GetSendQuotaOutput, error) {
	input := &ses.GetSendQuotaInput{}

	result, err := s.client.GetSendQuota(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get send quota: %w", err)
	}

	return result, nil
}