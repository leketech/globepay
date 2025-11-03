package service

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"globepay/internal/domain"
	"globepay/internal/domain/model"
	"globepay/internal/repository"
)

// AuthServiceInterface defines the interface for authentication service
type AuthServiceInterface interface {
	Login(email, password string) (string, error)
	Register(user *domain.User) error
	RefreshToken(tokenString string) (string, error)
	GenerateOTP() (string, error)
	ValidateOTP(otp, hashedOTP string) bool
	HashPassword(password string) (string, error)
}

// AuthService implements AuthServiceInterface
type AuthService struct {
	userRepo    repository.UserRepository
	jwtSecret   string
	tokenExpiry time.Duration
}

// NewAuthService creates a new AuthService
func NewAuthService(userRepo repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		jwtSecret:   jwtSecret,
		tokenExpiry: 24 * time.Hour, // 24 hours
	}
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(email, password string) (string, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	// Compare password using the model's CheckPassword method
	valid, err := user.CheckPassword(password)
	if err != nil || !valid {
		return "", domain.ErrInvalidCredentials
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(s.tokenExpiry).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

// Register creates a new user
func (s *AuthService) Register(user *domain.User) error {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return domain.ErrUserAlreadyExists
	}

	// Create a model.User from domain.User
	modelUser := &model.User{
		ID:            user.ID,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		PhoneNumber:   user.PhoneNumber,
		DateOfBirth:   user.DateOfBirth,
		Country:       user.Country,
		KYCStatus:     user.KYCStatus,
		AccountStatus: user.AccountStatus,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Set password using the model's SetPassword method
	if err := modelUser.SetPassword(user.PasswordHash); err != nil { // Using PasswordHash field
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Set default status
	modelUser.AccountStatus = "active"

	// Create user
	if err := s.userRepo.Create(modelUser); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// RefreshToken generates a new JWT token from a refresh token
func (s *AuthService) RefreshToken(tokenString string) (string, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	// Validate token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Get user ID from claims
	userID, ok := claims["user_id"].(string) // Changed from float64 to string
	if !ok {
		return "", fmt.Errorf("invalid user ID in token")
	}

	// Get user
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	// Generate new JWT token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(s.tokenExpiry).Unix(),
	})

	tokenString, err = newToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

// GenerateOTP generates a random OTP
func (s *AuthService) GenerateOTP() (string, error) {
	// Generate 6-digit random number
	otp := make([]byte, 3)
	if _, err := rand.Read(otp); err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}

	// Convert to 6-digit string
	otpString := fmt.Sprintf("%06d", uint32(otp[0])<<16|uint32(otp[1])<<8|uint32(otp[2]))

	return otpString, nil
}

// ValidateOTP validates an OTP against a hashed OTP
func (s *AuthService) ValidateOTP(otp, hashedOTP string) bool {
	// Hash the provided OTP
	hashedInput, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	// Compare with stored hash
	return subtle.ConstantTimeCompare(hashedInput, []byte(hashedOTP)) == 1
}

// HashPassword hashes a password using bcrypt
func (s *AuthService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// GeneratePasswordResetToken generates a password reset token
func (s *AuthService) GeneratePasswordResetToken(userID string) (string, error) { // Changed from int64 to string
	// Create a token with user ID and timestamp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(), // 1 hour expiry
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate password reset token: %w", err)
	}

	return base64.URLEncoding.EncodeToString([]byte(tokenString)), nil
}

// ValidatePasswordResetToken validates a password reset token
func (s *AuthService) ValidatePasswordResetToken(encodedToken string) (string, error) { // Changed from int64 to string
	// Decode the token
	tokenString, err := base64.URLEncoding.DecodeString(encodedToken)
	if err != nil {
		return "", fmt.Errorf("invalid token format")
	}

	// Parse token
	token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	// Validate token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Get user ID from claims
	userID, ok := claims["user_id"].(string) // Changed from float64 to string
	if !ok {
		return "", fmt.Errorf("invalid user ID in token")
	}

	// Check if token is expired
	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return "", fmt.Errorf("token expired")
	}

	return userID, nil // Changed from int64(userID) to userID
}