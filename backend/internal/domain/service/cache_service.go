package service

import (
	"context"
	"encoding/json"
	"time"

	"globepay/internal/infrastructure/cache"
)

// CacheService provides caching functionality
type CacheService struct {
	redisClient *cache.RedisClient
}

// NewCacheService creates a new cache service
func NewCacheService(redisClient *cache.RedisClient) *CacheService {
	return &CacheService{
		redisClient: redisClient,
	}
}

// Set stores a value in cache
func (s *CacheService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// Serialize the value to JSON
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return s.redisClient.Set(ctx, key, data, expiration)
}

// Get retrieves a value from cache
func (s *CacheService) Get(ctx context.Context, key string, dest interface{}) error {
	// Get the value from Redis
	data, err := s.redisClient.Get(ctx, key)
	if err != nil {
		return err
	}

	// Deserialize the JSON data
	return json.Unmarshal([]byte(data), dest)
}

// Delete removes a key from cache
func (s *CacheService) Delete(ctx context.Context, key string) error {
	return s.redisClient.Delete(ctx, key)
}

// SetUserSession stores user session data
func (s *CacheService) SetUserSession(ctx context.Context, userID, sessionID string, expiration time.Duration) error {
	key := "user_session:" + userID
	return s.Set(ctx, key, sessionID, expiration)
}

// GetUserSession retrieves user session data
func (s *CacheService) GetUserSession(ctx context.Context, userID string) (string, error) {
	key := "user_session:" + userID
	var sessionID string
	err := s.Get(ctx, key, &sessionID)
	return sessionID, err
}

// DeleteUserSession removes user session data
func (s *CacheService) DeleteUserSession(ctx context.Context, userID string) error {
	key := "user_session:" + userID
	return s.Delete(ctx, key)
}

// SetRateLimit stores rate limit data
func (s *CacheService) SetRateLimit(ctx context.Context, key string, requests []time.Time, expiration time.Duration) error {
	return s.Set(ctx, key, requests, expiration)
}

// GetRateLimit retrieves rate limit data
func (s *CacheService) GetRateLimit(ctx context.Context, key string) ([]time.Time, error) {
	var requests []time.Time
	err := s.Get(ctx, key, &requests)
	return requests, err
}

// SetTokenBlacklist blacklists a JWT token
func (s *CacheService) SetTokenBlacklist(ctx context.Context, token string, expiration time.Duration) error {
	key := "blacklisted_token:" + token
	return s.Set(ctx, key, true, expiration)
}

// IsTokenBlacklisted checks if a JWT token is blacklisted
func (s *CacheService) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	key := "blacklisted_token:" + token
	var blacklisted bool
	err := s.Get(ctx, key, &blacklisted)
	if err != nil {
		return false, nil // Token not found, not blacklisted
	}
	return blacklisted, nil
}
