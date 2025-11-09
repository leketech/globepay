package unit

import (
	"testing"

	"globepay/internal/domain/model"
	"globepay/internal/service"
	"globepay/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateAccount(t *testing.T) {
	// Create a mock account repository
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	// Create user service with mock repository
	userService := service.NewUserService(nil, mockAccountRepo)

	// Test data
	userID := "1"

	// Create a test account
	account := &model.Account{
		UserID:   userID,
		Currency: "USD",
		Balance:  100.0,
	}

	// Set up mock expectations
	mockAccountRepo.On("Create", mock.AnythingOfType("*model.Account")).Return(nil)

	// Call the method under test
	err := userService.CreateAccount(userID, account)

	// Assertions
	assert.NoError(t, err)

	// Verify mock expectations
	mockAccountRepo.AssertExpectations(t)
}

func TestUserService_GetUserByEmail(t *testing.T) {
	// Create a mock user repository
	mockUserRepo := new(mocks.UserRepositoryMock)

	// Create user service with mock repository
	userService := service.NewUserService(mockUserRepo, nil)

	// Test data
	email := "test@example.com"

	// Create a test user
	user := &model.User{
		ID:        "1",
		Email:     email,
		FirstName: "John",
		LastName:  "Doe",
	}

	// Set up mock expectations
	mockUserRepo.On("GetByEmail", email).Return(user, nil)

	// Call the method under test
	authUser, err := userService.GetUserByEmail(email)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, authUser)
	assert.Equal(t, user.ID, authUser.ID)
	assert.Equal(t, user.Email, authUser.Email)

	// Verify mock expectations
	mockUserRepo.AssertExpectations(t)
}