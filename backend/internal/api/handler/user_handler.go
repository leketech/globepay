package handler

import (
	"net/http"

	"globepay/internal/domain/service"
	"globepay/internal/utils"

	"github.com/gin-gonic/gin"
)

// UpdateUserProfileRequest represents the update user profile request body
type UpdateUserProfileRequest struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
}

// CreateUserAccountRequest represents the create account request body
type CreateUserAccountRequest struct {
	Currency string `json:"currency" binding:"required,len=3"`
}

// GetUserProfile handles getting user profile
func GetUserProfile(c *gin.Context, serviceFactory *service.ServiceFactory) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	userService := serviceFactory.GetUserService()
	user, err := userService.GetUserByID(c.Request.Context(), userID.(string))
	if err != nil {
		utils.InternalServerError(c, "USER_NOT_FOUND", "Failed to retrieve user profile")
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile handles updating user profile
func UpdateUserProfile(c *gin.Context, serviceFactory *service.ServiceFactory) {
	var req UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	userService := serviceFactory.GetUserService()
	user, err := userService.GetUserByID(c.Request.Context(), userID.(string))
	if err != nil {
		utils.InternalServerError(c, "USER_NOT_FOUND", "Failed to retrieve user profile")
		return
	}

	// Update user fields if provided
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}
	// Note: DateOfBirth update would require additional validation

	// Save updated user
	if err := userService.UpdateUser(c.Request.Context(), user); err != nil {
		utils.InternalServerError(c, "UPDATE_FAILED", "Failed to update user profile")
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserAccounts handles getting user accounts
func GetUserAccounts(c *gin.Context, serviceFactory *service.ServiceFactory) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	accountService := serviceFactory.GetAccountService()
	accounts, err := accountService.GetAccountsByUser(c.Request.Context(), userID.(string))
	if err != nil {
		utils.InternalServerError(c, "ACCOUNTS_NOT_FOUND", "Failed to retrieve accounts")
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// CreateUserAccount handles creating a new user account
func CreateUserAccount(c *gin.Context, serviceFactory *service.ServiceFactory) {
	var req CreateUserAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	accountService := serviceFactory.GetAccountService()
	account, err := accountService.CreateAccount(c.Request.Context(), userID.(string), req.Currency)
	if err != nil {
		// Check for specific error types
		if _, ok := err.(*service.ValidationError); ok {
			utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
			return
		}
		if _, ok := err.(*service.ConflictError); ok {
			utils.BadRequest(c, "ACCOUNT_EXISTS", err.Error())
			return
		}
		if _, ok := err.(*service.NotFoundError); ok {
			utils.NotFound(c, "USER_NOT_FOUND", err.Error())
			return
		}
		utils.InternalServerError(c, "ACCOUNT_CREATION_FAILED", "Failed to create account")
		return
	}

	c.JSON(http.StatusCreated, account)
}
