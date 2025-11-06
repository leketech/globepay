package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/utils"
)

// AuthService provides authentication-related functionality
type AuthService struct {
	userService *UserService
	jwtSecret   string
	jwtExpire   time.Duration
}

// NewAuthService creates a new authentication service
func NewAuthService(userService *UserService, jwtSecret string, jwtExpire time.Duration) *AuthService {
	return &AuthService{
		userService: userService,
		jwtSecret:   jwtSecret,
		jwtExpire:   jwtExpire,
	}
}

// Login authenticates a user and returns JWT tokens
func (s *AuthService) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	// Get user by email
	user, err := s.userService.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Printf("User not found for email %s: %v\n", email, err)
		return nil, &AuthenticationError{Message: "Invalid email or password"}
	}
	
	// Verify password
	match, err := utils.CheckPassword(password, user.PasswordHash)
	if err != nil {
		fmt.Printf("Error checking password for user %s: %v\n", email, err)
		return nil, &AuthenticationError{Message: "Invalid email or password"}
	}
	
	if !match {
		fmt.Printf("Invalid password for user %s\n", email)
		return nil, &AuthenticationError{Message: "Invalid email or password"}
	}
	
	// Generate access token
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, s.jwtExpire)
	if err != nil {
		fmt.Printf("Failed to generate access token for user %s: %v\n", email, err)
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}
	
	// Generate refresh token (longer expiration)
	refreshToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, 7*24*time.Hour)
	if err != nil {
		fmt.Printf("Failed to generate refresh token for user %s: %v\n", email, err)
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}
	
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// Register creates a new user and returns JWT tokens
func (s *AuthService) Register(ctx context.Context, email, password, firstName, lastName, phoneNumber, dateOfBirth, country string) (*LoginResponse, error) {
	fmt.Printf("Registering user: email=%s, firstName=%s, lastName=%s, phoneNumber=%s, dateOfBirth=%s, country=%s\n", email, firstName, lastName, phoneNumber, dateOfBirth, country)
	
	// Parse date of birth if provided
	var dob time.Time
	if dateOfBirth != "" {
		var err error
		dob, err = time.Parse("2006-01-02", dateOfBirth)
		if err != nil {
			fmt.Printf("Failed to parse date of birth: %v\n", err)
			// Continue without date of birth if parsing fails
		}
	}
	
	// Create user model with all details (password will be hashed in the model)
	user := model.NewUserWithDetails(email, password, firstName, lastName, phoneNumber, country, dob)
	
	fmt.Printf("Creating user in database: %+v\n", user)
	
	// Create user
	err := s.userService.CreateUser(ctx, user)
	if err != nil {
		fmt.Printf("Failed to create user in database: %v\n", err)
		// Check if it's a conflict error (user already exists)
		if strings.Contains(err.Error(), "user with email") && strings.Contains(err.Error(), "already exists") {
			return nil, &ConflictError{Message: "User with this email already exists"}
		}
		// Return a more generic error for other database issues
		return nil, fmt.Errorf("failed to register user: %w", err)
	}
	
	fmt.Printf("User created successfully: %+v\n", user)
	
	// Generate access token
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, s.jwtExpire)
	if err != nil {
		fmt.Printf("Failed to generate access token: %v\n", err)
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}
	
	// Generate refresh token
	refreshToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, 7*24*time.Hour)
	if err != nil {
		fmt.Printf("Failed to generate refresh token: %v\n", err)
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}
	
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// RefreshToken generates new tokens using a refresh token
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*RefreshResponse, error) {
	// Validate refresh token
	claims, err := utils.ValidateJWT(refreshToken, s.jwtSecret)
	if err != nil {
		return nil, &AuthenticationError{Message: "Invalid refresh token"}
	}
	
	// Get user
	user, err := s.userService.GetUserByID(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}
	
	// Generate new access token
	newAccessToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, err
	}
	
	// Generate new refresh token
	newRefreshToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}
	
	return &RefreshResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

// ValidateToken validates a JWT token and returns the user claims
func (s *AuthService) ValidateToken(ctx context.Context, tokenString string) (*utils.Claims, error) {
	claims, err := utils.ValidateJWT(tokenString, s.jwtSecret)
	if err != nil {
		return nil, &AuthenticationError{Message: "Invalid token"}
	}
	
	return claims, nil
}

// LoginResponse represents the response for login/registration
type LoginResponse struct {
	AccessToken  string      `json:"token"`
	RefreshToken string      `json:"refreshToken"`
	User         *model.User `json:"user"`
}

// RefreshResponse represents the response for token refresh
type RefreshResponse struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}