package unit

import (
	"context"
	"testing"

	"globepay/internal/domain/model"
	"globepay/internal/domain/service"
	"globepay/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateUser(t *testing.T) {
	// Create a mock user repository
	mockUserRepo := new(mocks.UserRepository)

	// Create user service with mock repository
	userService := service.NewUserService(mockUserRepo)

	// Test data
	email := "test@example.com"
	password := "password123"
	firstName := "John"
	lastName := "Doe"

	// Set up mock expectations
	mockUserRepo.On("GetByEmail", mock.Anything, email).Return(nil, nil)
	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)

	// Call the method under test
	user, err := userService.CreateUser(context.Background(), email, password, firstName, lastName)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, lastName, user.LastName)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "pending", user.KYCStatus)
	assert.Equal(t, "active", user.AccountStatus)

	// Verify mock expectations
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_InvalidEmail(t *testing.T) {
	// Create a mock user repository
	mockUserRepo := new(mocks.UserRepository)

	// Create user service with mock repository
	userService := service.NewUserService(mockUserRepo)

	// Test data with invalid email
	email := "invalid-email"
	password := "password123"
	firstName := "John"
	lastName := "Doe"

	// Call the method under test
	user, err := userService.CreateUser(context.Background(), email, password, firstName, lastName)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "Invalid email format")

	// Verify mock was not called
	mockUserRepo.AssertNotCalled(t, "GetByEmail")
	mockUserRepo.AssertNotCalled(t, "Create")
}

func TestUserService_CreateUser_ShortPassword(t *testing.T) {
	// Create a mock user repository
	mockUserRepo := new(mocks.UserRepository)

	// Create user service with mock repository
	userService := service.NewUserService(mockUserRepo)

	// Test data with short password
	email := "test@example.com"
	password := "123"
	firstName := "John"
	lastName := "Doe"

	// Call the method under test
	user, err := userService.CreateUser(context.Background(), email, password, firstName, lastName)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "Password must be at least 8 characters")

	// Verify mock was not called
	mockUserRepo.AssertNotCalled(t, "GetByEmail")
	mockUserRepo.AssertNotCalled(t, "Create")
}

func TestUserService_CreateUser_UserExists(t *testing.T) {
	// Create a mock user repository
	mockUserRepo := new(mocks.UserRepository)

	// Create user service with mock repository
	userService := service.NewUserService(mockUserRepo)

	// Test data
	email := "existing@example.com"
	password := "password123"
	firstName := "John"
	lastName := "Doe"

	// Set up mock expectations - user already exists
	existingUser := &model.User{Email: email}
	mockUserRepo.On("GetByEmail", mock.Anything, email).Return(existingUser, nil)

	// Call the method under test
	user, err := userService.CreateUser(context.Background(), email, password, firstName, lastName)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "User with this email already exists")

	// Verify mock expectations
	mockUserRepo.AssertExpectations(t)
	mockUserRepo.AssertNotCalled(t, "Create")
}

func TestUserService_AuthenticateUser(t *testing.T) {
	// Create a mock user repository
	mockUserRepo := new(mocks.UserRepository)

	// Create user service with mock repository
	userService := service.NewUserService(mockUserRepo)

	// Test data
	email := "test@example.com"
	password := "password123"

	// Create a test user with hashed password
	user := model.NewUser(email, password, "John", "Doe")

	// Set up mock expectations
	mockUserRepo.On("GetByEmail", mock.Anything, email).Return(user, nil)

	// Call the method under test
	authUser, err := userService.AuthenticateUser(context.Background(), email, password)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, authUser)
	assert.Equal(t, user.ID, authUser.ID)
	assert.Equal(t, user.Email, authUser.Email)

	// Verify mock expectations
	mockUserRepo.AssertExpectations(t)
}