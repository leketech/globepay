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
	// Authenticate user
	user, err := s.userService.AuthenticateUser(ctx, email, password)
	if err != nil {
		return nil, err
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
	// Create user
	user, err := s.userService.CreateUser(ctx, email, password, firstName, lastName)
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

// RegisterWithDetails creates a new user with additional details and returns JWT tokens
func (s *AuthService) RegisterWithDetails(ctx context.Context, email, password, firstName, lastName, phoneNumber, country string, dateOfBirth time.Time) (*LoginResponse, error) {
	// Create user with details
	user, err := s.userService.CreateUserWithDetails(ctx, email, password, firstName, lastName, phoneNumber, country, dateOfBirth)
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