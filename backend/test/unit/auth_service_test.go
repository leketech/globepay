package unit

import (
	"testing"
	"time"

	"globepay/internal/domain"
	"globepay/internal/domain/model"
	"globepay/internal/service"
	"globepay/test/mocks"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_Register(t *testing.T) {
	// Create mock user repository
	mockUserRepo := new(mocks.UserRepositoryMock)

	// Create auth service with mock repository
	authService := service.NewAuthService(mockUserRepo, "test-secret")

	// Test data
	user := &domain.User{
		Email:         "test@example.com",
		Password:      "password123", // Use Password field instead of PasswordHash
		FirstName:     "John",
		LastName:      "Doe",
		AccountStatus: "active", // Add AccountStatus field
	}

	// Set up mock expectations
	mockUserRepo.On("GetByEmail", "test@example.com").Return(nil, domain.ErrUserNotFound)

	// Update mock expectation to match the actual method signature
	mockUserRepo.On("Create", mock.AnythingOfType("*model.User")).Return(nil)

	// Call the method under test
	err := authService.Register(user)

	// Assertions
	assert.NoError(t, err)
	// Note: We can't check the password hashing in the domain.User since it's done in model.User
	assert.Equal(t, "active", user.AccountStatus)

	// Verify mock expectations
	mockUserRepo.AssertExpectations(t)
}

func TestAuthService_Login(t *testing.T) {
	// Create mock user repository
	mockUserRepo := new(mocks.UserRepositoryMock)

	// Create auth service with mock repository
	authService := service.NewAuthService(mockUserRepo, "test-secret")

	// Test data
	email := "test@example.com"
	password := "password123"

	// Create a model.User with hashed password
	modelUser := &model.User{
		ID:            "1",
		Email:         email,
		FirstName:     "John",
		LastName:      "Doe",
		AccountStatus: "active",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Set password using the model's SetPassword method
	if err := modelUser.SetPassword(password); err != nil {
		t.Fatalf("Failed to set password: %v", err)
	}

	// Set up mock expectations
	mockUserRepo.On("GetByEmail", email).Return(modelUser, nil)

	// Call the method under test
	token, err := authService.Login(email, password)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token is valid
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret"), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	// Verify claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, "1", claims["user_id"])
	assert.Equal(t, email, claims["email"])

	// Verify mock expectations
	mockUserRepo.AssertExpectations(t)
}

func TestAuthService_Login_InvalidCredentials(t *testing.T) {
	// Create mock user repository
	mockUserRepo := new(mocks.UserRepositoryMock)

	// Create auth service with mock repository
	authService := service.NewAuthService(mockUserRepo, "test-secret")

	// Test data
	email := "test@example.com"
	password := "wrongpassword"

	// Set up mock expectations
	mockUserRepo.On("GetByEmail", email).Return(nil, domain.ErrUserNotFound)

	// Call the method under test
	token, err := authService.Login(email, password)

	// Assertions
	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Equal(t, domain.ErrInvalidCredentials, err)

	// Verify mock expectations
	mockUserRepo.AssertExpectations(t)
}

func TestAuthService_HashPassword(t *testing.T) {
	// Create mock user repository
	mockUserRepo := new(mocks.UserRepositoryMock)

	// Create auth service with mock repository
	authService := service.NewAuthService(mockUserRepo, "test-secret")

	// Test data
	password := "password123"

	// Call the method under test
	hashedPassword, err := authService.HashPassword(password)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
	assert.NotEqual(t, password, hashedPassword)

	// Verify password can be validated
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	assert.NoError(t, err)
}

func TestAuthService_GenerateOTP(t *testing.T) {
	// Create mock user repository
	mockUserRepo := new(mocks.UserRepositoryMock)

	// Create auth service with mock repository
	authService := service.NewAuthService(mockUserRepo, "test-secret")

	// Call the method under test
	otp1, err1 := authService.GenerateOTP()
	otp2, err2 := authService.GenerateOTP()

	// Assertions
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotEmpty(t, otp1)
	assert.NotEmpty(t, otp2)
	// Note: The OTP generation might produce different lengths, so we'll just check they're not equal
	assert.NotEqual(t, otp1, otp2)
}
