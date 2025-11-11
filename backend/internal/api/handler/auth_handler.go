package handler

import (
	"fmt"
	"net/http"

	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/metrics"
	"globepay/internal/utils"

	"github.com/gin-gonic/gin"
)

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the registration request body
type RegisterRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	Country     string `json:"country,omitempty"`
}

// RefreshTokenRequest represents the refresh token request body
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// ForgotPasswordRequest represents the forgot password request body
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ResetPasswordRequest represents the reset password request body
type ResetPasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

// Login handles user login
func Login(c *gin.Context, serviceFactory *service.ServiceFactory) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("JSON binding error in login: %v\n", err)
		utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	authService := serviceFactory.GetAuthService()
	resp, err := authService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		// Log the error for debugging
		fmt.Printf("Login error for user %s: %v\n", req.Email, err)

		// Check if it's an authentication error
		if _, ok := err.(*service.AuthenticationError); ok {
			utils.Unauthorized(c, "INVALID_CREDENTIALS", "Invalid email or password")
			return
		}
		// Handle other errors as internal server errors
		utils.InternalServerError(c, "LOGIN_FAILED", "Failed to login")
		return
	}

	// Increment successful login metric
	if metricsInterface, exists := c.Get("metrics"); exists {
		if m, ok := metricsInterface.(*metrics.Metrics); ok {
			m.LoginAttemptsTotal.Inc()
		}
	}

	c.JSON(http.StatusOK, resp)
}

// Register handles user registration
func Register(c *gin.Context, serviceFactory *service.ServiceFactory) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("JSON binding error in register: %v\n", err)
		utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	authService := serviceFactory.GetAuthService()

	// Pass all user data to the service, not just basic fields
	resp, err := authService.Register(c.Request.Context(), req.Email, req.Password, req.FirstName, req.LastName, req.PhoneNumber, req.DateOfBirth, req.Country)
	if err != nil {
		// Log the error for debugging
		fmt.Printf("Registration error for user %s: %v\n", req.Email, err)

		// Check if it's a conflict error (user already exists)
		if _, ok := err.(*service.ConflictError); ok {
			utils.BadRequest(c, "USER_EXISTS", "User with this email already exists")
			return
		}
		// Check if it's a validation error
		if _, ok := err.(*service.ValidationError); ok {
			utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
			return
		}
		// Handle other errors as internal server errors
		utils.InternalServerError(c, "REGISTRATION_FAILED", "Failed to register user")
		return
	}

	// Increment successful registration metric
	if metricsInterface, exists := c.Get("metrics"); exists {
		if m, ok := metricsInterface.(*metrics.Metrics); ok {
			m.UsersRegisteredTotal.Inc()
		}
	}

	c.JSON(http.StatusCreated, resp)
}

// RefreshToken handles token refresh
func RefreshToken(c *gin.Context, serviceFactory *service.ServiceFactory) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	authService := serviceFactory.GetAuthService()
	resp, err := authService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		// Log the error for debugging
		fmt.Printf("Refresh token error: %v\n", err)

		// Check if it's an authentication error
		if _, ok := err.(*service.AuthenticationError); ok {
			utils.Unauthorized(c, "INVALID_TOKEN", "Invalid refresh token")
			return
		}
		utils.InternalServerError(c, "REFRESH_FAILED", "Failed to refresh token")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ForgotPassword handles forgot password requests
func ForgotPassword(c *gin.Context, serviceFactory *service.ServiceFactory) {
	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	// In a real implementation, you would:
	// 1. Check if user exists
	// 2. Generate a reset token
	// 3. Store the token with expiration
	// 4. Send email with reset link

	c.JSON(http.StatusOK, gin.H{
		"message": "If the email exists, a password reset link has been sent",
	})
}

// ResetPassword handles password reset
func ResetPassword(c *gin.Context, serviceFactory *service.ServiceFactory) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	// In a real implementation, you would:
	// 1. Validate the reset token
	// 2. Check token expiration
	// 3. Update user's password
	// 4. Invalidate the token

	c.JSON(http.StatusOK, gin.H{
		"message": "Password has been reset successfully",
	})
}
