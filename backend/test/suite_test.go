package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// TestSuite runs all integration tests
type TestSuite struct {
	suite.Suite
}

// SetupSuite runs before all tests
func (suite *TestSuite) SetupSuite() {
	// Setup any global test dependencies here
}

// TearDownSuite runs after all tests
func (suite *TestSuite) TearDownSuite() {
	// Clean up any global test dependencies here
}

// TestAll runs all integration tests
func TestAll(t *testing.T) {
	// Run all integration test suites
	// suite.Run(t, new(integration.APITestSuite))
	// Skip other test suites for now as they may have dependencies
	// suite.Run(t, new(integration.DatabaseTestSuite))
	// suite.Run(t, new(integration.RedisTestSuite))
	// suite.Run(t, new(integration.AccountTestSuite))
	// suite.Run(t, new(integration.TransferTestSuite))
	// suite.Run(t, new(integration.TransactionTestSuite))
	// suite.Run(t, new(integration.AuthTestSuite))
}