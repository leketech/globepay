package cache

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates a Redis client with proper URL handling + retry
func NewRedisClient(redisURL string) (*RedisClient, error) {
	if redisURL == "" {
		return nil, fmt.Errorf("redis URL cannot be empty")
	}

	// Normalize input:
	// Accept "redis://user:pass@host:6379/0"
	// Accept "host:6379"
	// Accept "host"
	if !strings.HasPrefix(redisURL, "redis://") {
		// If host only, add port
		if !strings.Contains(redisURL, ":") {
			redisURL = redisURL + ":6379"
		}
		redisURL = "redis://" + redisURL
	}

	// Parse redis:// URL
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("invalid redis URL: %w", err)
	}

	client := redis.NewClient(opt)
	ctx := context.Background()

	// Retry logic — 5 retries
	for i := 1; i <= 5; i++ {
		_, err := client.Ping(ctx).Result()
		if err == nil {
			log.Println("Redis connected successfully")
			return &RedisClient{client: client}, nil
		}

		log.Printf("Redis connection failed (%d/5): %v", i, err)
		time.Sleep(time.Duration(i) * time.Second) // exponential-ish backoff
	}

	return nil, fmt.Errorf("could not connect to Redis after retries")
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key %s does not exist", key)
		}
		return "", fmt.Errorf("failed to get key %s: %w", key, err)
	}
	return val, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := r.client.Set(ctx, key, value, expiration).Err(); err != nil {
		return fmt.Errorf("failed to set key %s: %w", key, err)
	}
	return nil
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete key %s: %w", key, err)
	}
	return nil
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

func (r *RedisClient) Ping(ctx context.Context) error {
	_, err := r.client.Ping(ctx).Result()
	return err
}

func (r *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check if key %s exists: %w", key, err)
	}
	return exists > 0, nil
}

func (r *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	if err := r.client.Expire(ctx, key, expiration).Err(); err != nil {
		return fmt.Errorf("failed to set expiration for key %s: %w", key, err)
	}
	return nil
}

func (r *RedisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get TTL for key %s: %w", key, err)
	}
	return ttl, nil
}

func (r *RedisClient) FlushDB(ctx context.Context) error {
	if err := r.client.FlushDB(ctx).Err(); err != nil {
		return fmt.Errorf("failed to flush database: %w", err)
	}
	log.Println("Redis DB flushed successfully")
	return nil
}

func (r *RedisClient) Increment(ctx context.Context, key string) (int64, error) {
	val, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to increment key %s: %w", key, err)
	}
	return val, nil
}

func (r *RedisClient) Decrement(ctx context.Context, key string) (int64, error) {
	val, err := r.client.Decr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to decrement key %s: %w", key, err)
	}
	return val, nil
}

// HashSet replaces HMSet (deprecated)
func (r *RedisClient) HashSet(ctx context.Context, key string, field string, value interface{}) error {
	if err := r.client.HSet(ctx, key, field, value).Err(); err != nil {
		return fmt.Errorf("failed to set hash field %s in key %s: %w", field, key, err)
	}
	return nil
}

func (r *RedisClient) HashGet(ctx context.Context, key string, field string) (string, error) {
	val, err := r.client.HGet(ctx, key, field).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("field %s does not exist in hash %s", field, key)
		}
		return "", fmt.Errorf("failed to get hash field %s from key %s: %w", field, key, err)
	}
	return val, nil
}

// HashSetMultiple using HSet for maps
func (r *RedisClient) HashSetMultiple(ctx context.Context, key string, values map[string]interface{}) error {
	if err := r.client.HSet(ctx, key, values).Err(); err != nil {
		return fmt.Errorf("failed to set multiple hash fields in key %s: %w", key, err)
	}
	return nil
}

func (r *RedisClient) HashGetAll(ctx context.Context, key string) (map[string]string, error) {
	values, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get all hash fields from key %s: %w", key, err)
	}
	return values, nil
}