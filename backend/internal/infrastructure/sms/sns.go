package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// SNSClient wraps the AWS SNS client
type SNSClient struct {
	client *sns.Client
	config aws.Config
}

// NewSNSClient creates a new SNS client
func NewSNSClient(config aws.Config) *SNSClient {
	return &SNSClient{
		client: sns.NewFromConfig(config),
		config: config,
	}
}

// SendSMS sends an SMS using AWS SNS
func (s *SNSClient) SendSMS(ctx context.Context, phoneNumber, message string) error {
	input := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phoneNumber),
	}

	_, err := s.client.Publish(ctx, input)
	return err
}