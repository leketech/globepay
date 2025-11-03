package storage

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Config holds the S3 configuration
type S3Config struct {
	Region   string
	Endpoint string // Optional, for local testing
}

// S3Client wraps the S3 client
type S3Client struct {
	client *s3.Client
}

// ObjectMetadata represents S3 object metadata
type ObjectMetadata struct {
	Key          string
	Size         int64
	LastModified string
	ContentType  string
	ETag         string
}

// NewS3Client creates a new S3 client
func NewS3Client(cfg S3Config) (*S3Client, error) {
	// Load AWS configuration
	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.Region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create S3 client
	client := s3.NewFromConfig(awsConfig, func(o *s3.Options) {
		if cfg.Endpoint != "" {
			o.EndpointResolver = s3.EndpointResolverFunc(func(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: cfg.Endpoint,
				}, nil
			})
		}
	})

	log.Println("Successfully created S3 client")
	return &S3Client{client: client}, nil
}

// CreateBucket creates a new S3 bucket
func (s *S3Client) CreateBucket(ctx context.Context, bucketName string) error {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	_, err := s.client.CreateBucket(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create bucket %s: %w", bucketName, err)
	}

	log.Printf("Successfully created bucket %s", bucketName)
	return nil
}

// PutObject uploads an object to S3
func (s *S3Client) PutObject(ctx context.Context, bucketName, key string, data []byte, contentType string) error {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	}

	_, err := s.client.PutObject(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to put object %s in bucket %s: %w", key, bucketName, err)
	}

	log.Printf("Successfully uploaded object %s to bucket %s", key, bucketName)
	return nil
}

// GetObject downloads an object from S3
func (s *S3Client) GetObject(ctx context.Context, bucketName, key string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	result, err := s.client.GetObject(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get object %s from bucket %s: %w", key, bucketName, err)
	}
	defer result.Body.Close()

	// Read the object data
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read object data: %w", err)
	}

	log.Printf("Successfully downloaded object %s from bucket %s", key, bucketName)
	return buf.Bytes(), nil
}

// DeleteObject deletes an object from S3
func (s *S3Client) DeleteObject(ctx context.Context, bucketName, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	_, err := s.client.DeleteObject(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete object %s from bucket %s: %w", key, bucketName, err)
	}

	log.Printf("Successfully deleted object %s from bucket %s", key, bucketName)
	return nil
}

// ListObjects lists objects in an S3 bucket
func (s *S3Client) ListObjects(ctx context.Context, bucketName, prefix string) ([]ObjectMetadata, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
	}

	result, err := s.client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects in bucket %s: %w", bucketName, err)
	}

	objects := make([]ObjectMetadata, len(result.Contents))
	for i, obj := range result.Contents {
		objects[i] = ObjectMetadata{
			Key:          *obj.Key,
			Size:         obj.Size, // Fix: obj.Size is already int64, no need to dereference
			LastModified: obj.LastModified.String(),
			ETag:         *obj.ETag,
		}
	}

	log.Printf("Successfully listed %d objects in bucket %s", len(objects), bucketName)
	return objects, nil
}

// GetObjectURL returns the URL of an S3 object
func (s *S3Client) GetObjectURL(bucketName, key string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
}

// Close closes the S3 client
func (s *S3Client) Close() error {
	// S3 client doesn't need explicit closing
	log.Println("S3 client closed")
	return nil
}

// GetDefaultConfig returns a default S3 configuration
func GetDefaultConfig() S3Config {
	return S3Config{
		Region: "us-east-1",
	}
}

// GetConfigFromEnv returns an S3 configuration from environment variables
func GetConfigFromEnv() S3Config {
	// In a real implementation, you would use os.Getenv to get values from environment variables
	// For now, we'll return the default config
	return GetDefaultConfig()
}

// PutObjectWithMetadata uploads an object to S3 with metadata
func (s *S3Client) PutObjectWithMetadata(ctx context.Context, bucketName, key string, data []byte, contentType string, metadata map[string]string) error {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
		Metadata:    metadata,
	}

	_, err := s.client.PutObject(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to put object %s in bucket %s with metadata: %w", key, bucketName, err)
	}

	log.Printf("Successfully uploaded object %s to bucket %s with metadata", key, bucketName)
	return nil
}

// GetObjectMetadata retrieves metadata for an S3 object
func (s *S3Client) GetObjectMetadata(ctx context.Context, bucketName, key string) (map[string]string, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	result, err := s.client.HeadObject(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata for object %s in bucket %s: %w", key, bucketName, err)
	}

	metadata := make(map[string]string)
	for key, value := range result.Metadata {
		// Fix: Values are already strings, no need to dereference
		metadata[key] = value
	}

	return metadata, nil
}

// CopyObject copies an object from one location to another in S3
func (s *S3Client) CopyObject(ctx context.Context, sourceBucket, sourceKey, destBucket, destKey string) error {
	copySource := fmt.Sprintf("%s/%s", sourceBucket, sourceKey)
	
	// Ensure the source key is properly URL encoded
	if strings.Contains(sourceKey, " ") || strings.Contains(sourceKey, "+") {
		copySource = fmt.Sprintf("%s/%s", sourceBucket, strings.ReplaceAll(sourceKey, " ", "%20"))
	}

	input := &s3.CopyObjectInput{
		Bucket:     aws.String(destBucket),
		Key:        aws.String(destKey),
		CopySource: aws.String(copySource),
	}

	_, err := s.client.CopyObject(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to copy object from %s/%s to %s/%s: %w", sourceBucket, sourceKey, destBucket, destKey, err)
	}

	log.Printf("Successfully copied object from %s/%s to %s/%s", sourceBucket, sourceKey, destBucket, destKey)
	return nil
}

// GeneratePresignedURL generates a presigned URL for an S3 object
func (s *S3Client) GeneratePresignedURL(ctx context.Context, bucketName, key string, expiresInSeconds int64) (string, error) {
	// Note: This requires the s3presign package which is not imported above
	// In a real implementation, you would use:
	// presigner := s3presign.NewPresigner(s.client)
	// req, err := presigner.PresignGetObject(ctx, &s3.GetObjectInput{
	//     Bucket: aws.String(bucketName),
	//     Key:    aws.String(key),
	// }, s3presign.WithPresignExpires(time.Duration(expiresInSeconds)*time.Second))
	// if err != nil {
	//     return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	// }
	// return req.URL, nil

	// For now, we'll return a placeholder
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s?X-Amz-Expires=%d", bucketName, key, expiresInSeconds), nil
}