package integration

import (
	"context"
	"testing"
	"time"

	"globepay/internal/infrastructure/cache"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RedisTestSuite struct {
	suite.Suite
	redisClient *cache.RedisClient
}

func (suite *RedisTestSuite) SetupSuite() {
	// In a real test, you would connect to a test Redis instance
	// For now, we'll skip the actual Redis connection
	suite.T().Skip("Skipping Redis tests - no test Redis instance configured")
}

func (suite *RedisTestSuite) TestRedis_SetAndGet() {
	// Skip if no Redis connection
	if suite.redisClient == nil {
		suite.T().Skip("No Redis connection")
	}

	ctx := context.Background()
	key := "test_key"
	value := "test_value"
	expiration := 1 * time.Minute

	// Test setting a value
	err := suite.redisClient.Set(ctx, key, value, expiration)
	assert.NoError(suite.T(), err)

	// Test getting the value
	retrievedValue, err := suite.redisClient.Get(ctx, key)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), value, retrievedValue)
}

func (suite *RedisTestSuite) TestRedis_Delete() {
	// Skip if no Redis connection
	if suite.redisClient == nil {
		suite.T().Skip("No Redis connection")
	}

	ctx := context.Background()
	key := "test_key_delete"
	value := "test_value"
	expiration := 1 * time.Minute

	// First set a value
	err := suite.redisClient.Set(ctx, key, value, expiration)
	assert.NoError(suite.T(), err)

	// Delete the value
	err = suite.redisClient.Delete(ctx, key)
	assert.NoError(suite.T(), err)

	// Try to get the deleted value (should fail)
	_, err = suite.redisClient.Get(ctx, key)
	assert.Error(suite.T(), err)
}

func (suite *RedisTestSuite) TestRedis_ExpiredKey() {
	// Skip if no Redis connection
	if suite.redisClient == nil {
		suite.T().Skip("No Redis connection")
	}

	ctx := context.Background()
	key := "test_key_expired"
	value := "test_value"
	expiration := 1 * time.Second

	// Set a value with short expiration
	err := suite.redisClient.Set(ctx, key, value, expiration)
	assert.NoError(suite.T(), err)

	// Wait for expiration
	time.Sleep(2 * time.Second)

	// Try to get the expired value (should fail)
	_, err = suite.redisClient.Get(ctx, key)
	assert.Error(suite.T(), err)
}

func TestRedisTestSuite(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}