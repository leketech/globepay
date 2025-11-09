package unit

import (
	"testing"

	"globepay/internal/domain/model"
	"globepay/internal/service"
	"globepay/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceSimple(t *testing.T) {
	// Create mock repositories
	mockUserRepo := new(mocks.UserRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	// Create user service with mock repositories
	userService := service.NewUserService(mockUserRepo, mockAccountRepo)

	// Test data
	userID := "1" // Changed from int64 to string

	// Create a test account
	account := &model.Account{
		UserID:   userID,
		Currency: "USD",
		Balance:  100.0,
	}

	// Set up mock expectations
	mockAccountRepo.On("Create", mock.AnythingOfType("*model.Account")).Return(nil)

	// Call the method under test
	err := userService.CreateAccount(userID, account) // Changed to match the correct method signature

	// Assertions
	assert.NoError(t, err)

	// Verify mock expectations
	mockAccountRepo.AssertExpectations(t)
}