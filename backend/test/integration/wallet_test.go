package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"globepay/internal/api/router"
	"globepay/internal/config"
	"globepay/internal/domain/model"
	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/metrics"
	"globepay/test/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestWalletAddMoney tests the /api/v1/wallet/add endpoint
func TestWalletAddMoney(t *testing.T) {
	// Skip this test in short mode
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	// Setup test database
	testDB := utils.NewTestDB()
	if testDB == nil {
		t.Skip("No test database connection")
	}
	defer testDB.Close()

	// Clear tables
	testDB.ClearTables()

	// Setup
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	// Use test database instead of the real one
	db := testDB.DB
	// Create a default AWS config for testing
	awsConfig := aws.Config{}
	serviceFactory := service.NewFactory(cfg, db, nil, awsConfig)
	metrics := metrics.NewMetrics()

	// Create test router
	gin.SetMode(gin.TestMode)
	r := gin.New()
	router.SetupRoutes(r, serviceFactory, metrics)

	// Create test user and account
	user := createTestUser(t, serviceFactory)
	_ = createTestAccount(t, serviceFactory, user.ID, "USD")

	// Prepare request body
	requestBody := map[string]interface{}{
		"amount":         100.0,
		"payment_method": "card",
		"card_number":    "4111111111111111",
		"expiry_date":    "12/25",
		"cvv":            "123",
	}

	// Convert to JSON
	jsonBody, err := json.Marshal(requestBody)
	assert.NoError(t, err)

	// Create request with auth token (in a real test, you'd generate a valid JWT)
	req, err := http.NewRequest("POST", "/api/v1/wallet/add", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	// In a real test, you would set a valid auth token here
	// req.Header.Set("Authorization", "Bearer "+validToken)

	// Create response recorder
	w := httptest.NewRecorder()

	// Perform request
	r.ServeHTTP(w, req)

	// Note: This test will fail because we don't have a valid auth token
	// In a real implementation, you would:
	// 1. Create a valid user and get a JWT token
	// 2. Set the Authorization header with the token
	// 3. Assert the response status and body

	fmt.Printf("Response Status: %d\n", w.Code)
	fmt.Printf("Response Body: %s\n", w.Body.String())
}

// TestWalletRequestMoney tests the /api/v1/wallet/request endpoint
func TestWalletRequestMoney(t *testing.T) {
	// Skip this test in short mode
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	// Setup test database
	testDB := utils.NewTestDB()
	if testDB == nil {
		t.Skip("No test database connection")
	}
	defer testDB.Close()

	// Clear tables
	testDB.ClearTables()

	// Setup
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	// Use test database instead of the real one
	db := testDB.DB
	// Create a default AWS config for testing
	awsConfig := aws.Config{}
	serviceFactory := service.NewFactory(cfg, db, nil, awsConfig)
	metrics := metrics.NewMetrics()

	// Create test router
	gin.SetMode(gin.TestMode)
	r := gin.New()
	router.SetupRoutes(r, serviceFactory, metrics)

	// Create test users and accounts
	requester := createTestUser(t, serviceFactory)
	_ = requester
	recipient := createTestUser(t, serviceFactory)
	_ = recipient

	// Prepare request body for user request
	requestBody := map[string]interface{}{
		"amount":       50.0,
		"recipient_id": recipient.ID,
		"description":  "Test money request",
		"is_link":      false,
	}

	// Convert to JSON
	jsonBody, err := json.Marshal(requestBody)
	assert.NoError(t, err)

	// Create request with auth token
	req, err := http.NewRequest("POST", "/api/v1/wallet/request", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	// In a real test, you would set a valid auth token here
	// req.Header.Set("Authorization", "Bearer "+validToken)

	// Create response recorder
	w := httptest.NewRecorder()

	// Perform request
	r.ServeHTTP(w, req)

	// Note: This test will fail because we don't have a valid auth token
	// In a real implementation, you would:
	// 1. Create a valid user and get a JWT token
	// 2. Set the Authorization header with the token
	// 3. Assert the response status and body

	fmt.Printf("Response Status: %d\n", w.Code)
	fmt.Printf("Response Body: %s\n", w.Body.String())
}

// Helper functions for creating test data
func createTestUser(t *testing.T, serviceFactory *service.Factory) *model.User {
	userService := serviceFactory.GetUserService()
	user := model.NewUser(
		fmt.Sprintf("test%d@example.com", time.Now().Unix()),
		"password123",
		"Test",
		"User",
	)
	err := userService.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	return user
}

func createTestAccount(t *testing.T, serviceFactory *service.Factory, userID, currency string) *model.Account {
	accountService := serviceFactory.GetAccountService()
	account, err := accountService.CreateAccount(context.Background(), userID, currency)
	assert.NoError(t, err)
	return account
}
