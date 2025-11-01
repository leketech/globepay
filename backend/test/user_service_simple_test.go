package test

import (
	"testing"

	"globepay/internal/domain"
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
	email := "test@example.com"
	firstName := "John"
	lastName := "Doe"

	// Set up mock expectations
	mockUserRepo.On("GetByEmail", email).Return(nil, nil) // User doesn't exist
	mockUserRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)

	// This is a simplified test - in reality, we would need to handle password hashing
	// but for now we're just testing the structure
	user := &domain.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	// Call the method under test
	err := userService.CreateAccount(1, user)

	// Assertions
	assert.NoError(t, err)

	// Verify mock expectations
	mockUserRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
}