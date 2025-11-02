package service

import (
	"context"
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
		return nil, &AuthenticationError{Message: "Invalid email or password"}
	}
	
	// Verify password
	match, err := utils.CheckPassword(password, user.PasswordHash)
	if err != nil || !match {
		return nil, &AuthenticationError{Message: "Invalid email or password"}
	}
	
	// Generate access token
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, err
	}
	
	// Generate refresh token (longer expiration)
	refreshToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}
	
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// Register creates a new user and returns JWT tokens
func (s *AuthService) Register(ctx context.Context, email, password, firstName, lastName string) (*LoginResponse, error) {
	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	
	// Create user model
	user := &model.User{
		Email:        email,
		PasswordHash: hashedPassword,
		FirstName:    firstName,
		LastName:     lastName,
		AccountStatus: "active",
		KYCStatus:    "pending",
	}
	
	// Create user
	err = s.userService.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	
	// Generate access token
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, err
	}
	
	// Generate refresh token
	refreshToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, 7*24*time.Hour)
	if err != nil {
		return nil, err
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