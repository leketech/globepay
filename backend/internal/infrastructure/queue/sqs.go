package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"errors"
)

// SQSConfig holds the SQS configuration
type SQSConfig struct {
	Region   string
	Endpoint string // Optional, for local testing
}

// SQSClient wraps the SQS client
type SQSClient struct {
	client *sqs.Client
}

// Message represents an SQS message
type Message struct {
	ID            string
	Body          string
	ReceiptHandle string
	Attributes    map[string]string
}

// NewSQSClient creates a new SQS client
func NewSQSClient(cfg SQSConfig) (*SQSClient, error) {
	// Load AWS configuration
	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.Region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create SQS client
	client := sqs.NewFromConfig(awsConfig, func(o *sqs.Options) {
		if cfg.Endpoint != "" {
			o.EndpointResolver = sqs.EndpointResolverFunc(func(region string, options sqs.EndpointResolverOptions) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: cfg.Endpoint,
				}, nil
			})
		}
	})

	log.Println("Successfully created SQS client")
	return &SQSClient{client: client}, nil
}

// CreateQueue creates a new SQS queue
func (s *SQSClient) CreateQueue(ctx context.Context, queueName string) (string, error) {
	input := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
		Attributes: map[string]string{
			"DelaySeconds":      "0",
			"VisibilityTimeout": "30",
		},
	}

	result, err := s.client.CreateQueue(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to create queue %s: %w", queueName, err)
	}

	log.Printf("Successfully created queue %s", queueName)
	return *result.QueueUrl, nil
}

// SendMessage sends a message to an SQS queue
func (s *SQSClient) SendMessage(ctx context.Context, queueURL, messageBody string) (string, error) {
	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(messageBody),
	}

	result, err := s.client.SendMessage(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to send message: %w", err)
	}

	log.Printf("Successfully sent message to queue %s", queueURL)
	return *result.MessageId, nil
}

// ReceiveMessages receives messages from an SQS queue
func (s *SQSClient) ReceiveMessages(ctx context.Context, queueURL string, maxMessages int32) ([]Message, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: maxMessages,
		WaitTimeSeconds:     20, // Enable long polling
	}

	result, err := s.client.ReceiveMessage(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to receive messages: %w", err)
	}

	messages := make([]Message, len(result.Messages))
	for i, msg := range result.Messages {
		attributes := make(map[string]string)
		for key, value := range msg.Attributes {
			if value != nil {
				attributes[key] = *value
			} else {
				attributes[key] = ""
			}
		}

		messages[i] = Message{
			ID:            *msg.MessageId,
			Body:          *msg.Body,
			ReceiptHandle: *msg.ReceiptHandle,
			Attributes:    attributes,
		}
	}

	log.Printf("Successfully received %d messages from queue %s", len(messages), queueURL)
	return messages, nil
}

// DeleteMessage deletes a message from an SQS queue
func (s *SQSClient) DeleteMessage(ctx context.Context, queueURL, receiptHandle string) error {
	input := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	}

	_, err := s.client.DeleteMessage(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	log.Printf("Successfully deleted message from queue %s", queueURL)
	return nil
}

// GetQueueURL gets the URL of an SQS queue
func (s *SQSClient) GetQueueURL(ctx context.Context, queueName string) (string, error) {
	input := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}

	result, err := s.client.GetQueueUrl(ctx, input)
	if err != nil {
		var queueDoesNotExist *types.QueueDoesNotExist
		if errors.As(err, &queueDoesNotExist) {
			return "", fmt.Errorf("queue %s does not exist", queueName)
		}
		return "", fmt.Errorf("failed to get queue URL: %w", err)
	}

	return *result.QueueUrl, nil
}

// Close closes the SQS client
func (s *SQSClient) Close() error {
	// SQS client doesn't need explicit closing
	log.Println("SQS client closed")
	return nil
}

// GetDefaultConfig returns a default SQS configuration
func GetDefaultConfig() SQSConfig {
	return SQSConfig{
		Region: "us-east-1",
	}
}

// GetConfigFromEnv returns an SQS configuration from environment variables
func GetConfigFromEnv() SQSConfig {
	// In a real implementation, you would use os.Getenv to get values from environment variables
	// For now, we'll return the default config
	return GetDefaultConfig()
}

// SendMessageBatch sends a batch of messages to an SQS queue
func (s *SQSClient) SendMessageBatch(ctx context.Context, queueURL string, messages []string) ([]string, error) {
	entries := make([]types.SendMessageBatchRequestEntry, len(messages))
	for i, msg := range messages {
		entries[i] = types.SendMessageBatchRequestEntry{
			Id:          aws.String(fmt.Sprintf("msg-%d", i)),
			MessageBody: aws.String(msg),
		}
	}

	input := &sqs.SendMessageBatchInput{
		QueueUrl: aws.String(queueURL),
		Entries:  entries,
	}

	result, err := s.client.SendMessageBatch(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to send message batch: %w", err)
	}

	messageIDs := make([]string, len(result.Successful))
	for i, success := range result.Successful {
		messageIDs[i] = *success.MessageId
	}

	log.Printf("Successfully sent batch of %d messages to queue %s", len(messageIDs), queueURL)
	return messageIDs, nil
}

// ChangeMessageVisibility changes the visibility timeout of a message
func (s *SQSClient) ChangeMessageVisibility(ctx context.Context, queueURL, receiptHandle string, visibilityTimeout int32) error {
	input := &sqs.ChangeMessageVisibilityInput{
		QueueUrl:          aws.String(queueURL),
		ReceiptHandle:     aws.String(receiptHandle),
		VisibilityTimeout: visibilityTimeout,
	}

	_, err := s.client.ChangeMessageVisibility(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to change message visibility: %w", err)
	}

	log.Printf("Successfully changed message visibility for queue %s", queueURL)
	return nil
}

// PurgeQueue deletes all messages in a queue
func (s *SQSClient) PurgeQueue(ctx context.Context, queueURL string) error {
	input := &sqs.PurgeQueueInput{
		QueueUrl: aws.String(queueURL),
	}

	_, err := s.client.PurgeQueue(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to purge queue: %w", err)
	}

	log.Printf("Successfully purged queue %s", queueURL)
	return nil
}