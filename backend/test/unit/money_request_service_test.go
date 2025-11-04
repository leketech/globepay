package unit

import (
	"context"
	"testing"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/domain/service"
	"globepay/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMoneyRequestService_CreateRequest(t *testing.T) {
	// Create mock repositories
	mockMoneyRequestRepo := new(mocks.MoneyRequestRepository)
	mockAccountRepo := new(mocks.AccountRepository)

	// Create money request service with mock repositories
	moneyRequestService := service.NewMoneyRequestService(mockMoneyRequestRepo, mockAccountRepo)

	// Test data
	requesterID := "requester-123"
	recipientID := "recipient-456"
	amount := 100.0
	currency := "USD"
	description := "Test money request"

	// Set up mock expectations
	mockMoneyRequestRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.MoneyRequest")).Return(nil)

	// Call the method under test
	request, err := moneyRequestService.CreateRequest(context.Background(), requesterID, recipientID, amount, currency, description)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, request)
	assert.Equal(t, requesterID, request.RequesterID)
	assert.Equal(t, recipientID, request.RecipientID)
	assert.Equal(t, amount, request.Amount)
	assert.Equal(t, currency, request.Currency)
	assert.Equal(t, description, request.Description)
	assert.Equal(t, "pending", request.Status)
	assert.NotEmpty(t, request.ID)
	assert.WithinDuration(t, time.Now().Add(30*24*time.Hour), request.ExpiresAt, time.Minute)

	// Verify mock expectations
	mockMoneyRequestRepo.AssertExpectations(t)
}

func TestMoneyRequestService_CreatePaymentLink(t *testing.T) {
	// Create mock repositories
	mockMoneyRequestRepo := new(mocks.MoneyRequestRepository)

	// Create money request service with mock repositories
	// We'll use nil for accountRepo since we're not testing account functionality
	moneyRequestService := service.NewMoneyRequestService(mockMoneyRequestRepo, nil)

	// Test data
	requestID := "request-123"
	expectedPaymentLink := "/pay/request-123/some-token"

	// Set up mock expectations
	mockMoneyRequestRepo.On("UpdatePaymentLink", mock.Anything, requestID, mock.AnythingOfType("string")).Return(nil)

	// Call the method under test
	paymentLink, err := moneyRequestService.CreatePaymentLink(context.Background(), requestID)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentLink)
	assert.Contains(t, paymentLink, "/pay/"+requestID)

	// Verify mock expectations
	mockMoneyRequestRepo.AssertExpectations(t)
}