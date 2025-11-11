package integration

import (
	"context"
	"testing"
	"time"

	"globepay/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RedisTestSuite struct {
	suite.Suite
	redisClient *utils.TestRedis
}

func (suite *RedisTestSuite) SetupSuite() {
	// Initialize test Redis
	suite.redisClient = utils.NewTestRedis()

	// Skip if no Redis connection
	if suite.redisClient == nil {
		suite.T().Skip("Skipping Redis tests - no test Redis instance configured")
	}
}

func (suite *RedisTestSuite) TearDownSuite() {
	if suite.redisClient != nil {
		suite.redisClient.Close()
	}
}

func (suite *RedisTestSuite) SetupTest() {
	// Clear Redis data before each test
	if suite.redisClient != nil {
		suite.redisClient.Clear()
	}
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
	err := suite.redisClient.Client.Set(ctx, key, value, expiration).Err()
	assert.NoError(suite.T(), err)

	// Test getting the value
	retrievedValue, err := suite.redisClient.Client.Get(ctx, key).Result()
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
	err := suite.redisClient.Client.Set(ctx, key, value, expiration).Err()
	assert.NoError(suite.T(), err)

	// Delete the value
	err = suite.redisClient.Client.Del(ctx, key).Err()
	assert.NoError(suite.T(), err)

	// Try to get the deleted value (should fail)
	_, err = suite.redisClient.Client.Get(ctx, key).Result()
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
	err := suite.redisClient.Client.Set(ctx, key, value, expiration).Err()
	assert.NoError(suite.T(), err)

	// Wait for expiration
	time.Sleep(2 * time.Second)

	// Try to get the expired value (should fail)
	_, err = suite.redisClient.Client.Get(ctx, key).Result()
	assert.Error(suite.T(), err)
}

func TestRedisTestSuite(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}
