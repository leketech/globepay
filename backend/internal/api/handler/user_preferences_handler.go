package handler

import (
	"net/http"

	"globepay/internal/domain/model"
	"globepay/internal/domain/service"
	"globepay/internal/utils"

	"github.com/gin-gonic/gin"
)

// UserPreferencesRequest represents the user preferences request body
type UserPreferencesRequest struct {
	EmailNotifications bool `json:"email_notifications"`
	PushNotifications  bool `json:"push_notifications"`
	SMSNotifications   bool `json:"sms_notifications"`
	TransactionAlerts  bool `json:"transaction_alerts"`
	SecurityAlerts     bool `json:"security_alerts"`
	MarketingEmails    bool `json:"marketing_emails"`
	TwoFactorEnabled   bool `json:"two_factor_enabled"`
}

// GetUserPreferences handles getting user preferences
func GetUserPreferences(c *gin.Context, serviceFactory *service.Factory) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	userService := serviceFactory.GetUserService()
	preferences, err := userService.GetUserPreferences(c.Request.Context(), userID.(string))
	if err != nil {
		utils.InternalServerError(c, "PREFERENCES_NOT_FOUND", "Failed to retrieve user preferences")
		return
	}

	c.JSON(http.StatusOK, preferences)
}

// UpdateUserPreferences handles updating user preferences
func UpdateUserPreferences(c *gin.Context, serviceFactory *service.Factory) {
	var req UserPreferencesRequest
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

	// Get existing preferences or create new ones
	preferences, err := userService.GetUserPreferences(c.Request.Context(), userID.(string))
	if err != nil {
		// Create new preferences if they don't exist
		preferences = model.NewUserPreferences(userID.(string))
	}

	// Update preferences with request data
	preferences.EmailNotifications = req.EmailNotifications
	preferences.PushNotifications = req.PushNotifications
	preferences.SMSNotifications = req.SMSNotifications
	preferences.TransactionAlerts = req.TransactionAlerts
	preferences.SecurityAlerts = req.SecurityAlerts
	preferences.MarketingEmails = req.MarketingEmails
	preferences.TwoFactorEnabled = req.TwoFactorEnabled

	// Save updated preferences
	if err := userService.UpdateUserPreferences(c.Request.Context(), preferences); err != nil {
		utils.InternalServerError(c, "UPDATE_FAILED", "Failed to update user preferences")
		return
	}

	c.JSON(http.StatusOK, preferences)
}
