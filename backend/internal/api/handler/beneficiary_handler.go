package handler

import (
	"net/http"

	"globepay/internal/domain/model"
	"globepay/internal/domain/service"
	"globepay/internal/utils"

	"github.com/gin-gonic/gin"
)

// CreateBeneficiaryRequest represents the create beneficiary request body
type CreateBeneficiaryRequest struct {
	Name          string `json:"name" binding:"required"`
	Country       string `json:"country" binding:"required,len=2"`
	BankName      string `json:"bankName" binding:"required"`
	AccountNumber string `json:"accountNumber" binding:"required"`
	SwiftCode     string `json:"swiftCode,omitempty"`
	Iban          string `json:"iban,omitempty"`
	BankAddress   string `json:"bankAddress,omitempty"`
	Currency      string `json:"currency,omitempty"`
}

// UpdateBeneficiaryRequest represents the update beneficiary request body
type UpdateBeneficiaryRequest struct {
	Name          string `json:"name,omitempty"`
	Country       string `json:"country,omitempty,len=2"`
	BankName      string `json:"bankName,omitempty"`
	AccountNumber string `json:"accountNumber,omitempty"`
	SwiftCode     string `json:"swiftCode,omitempty"`
	Iban          string `json:"iban,omitempty"`
	BankAddress   string `json:"bankAddress,omitempty"`
	Currency      string `json:"currency,omitempty"`
}

// GetBeneficiaries handles getting user beneficiaries
func GetBeneficiaries(c *gin.Context, serviceFactory *service.Factory) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	beneficiaryService := serviceFactory.GetBeneficiaryService()
	beneficiaries, err := beneficiaryService.GetBeneficiariesByUser(c.Request.Context(), userID.(string))
	if err != nil {
		utils.InternalServerError(c, "BENEFICIARIES_NOT_FOUND", "Failed to retrieve beneficiaries")
		return
	}

	c.JSON(http.StatusOK, beneficiaries)
}

// CreateBeneficiary handles creating a new beneficiary
func CreateBeneficiary(c *gin.Context, serviceFactory *service.Factory) {
	var req CreateBeneficiaryRequest
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

	beneficiaryService := serviceFactory.GetBeneficiaryService()

	// Create beneficiary model
	beneficiary := &model.Beneficiary{
		UserID:      userID.(string),
		Name:        req.Name,
		Country:     req.Country,
		BankName:    req.BankName,
		AccountNo:   req.AccountNumber,
		SwiftCode:   req.SwiftCode,
		Iban:        req.Iban,
		BankAddress: req.BankAddress,
		Currency:    req.Currency,
	}

	// Create beneficiary
	if err := beneficiaryService.CreateBeneficiary(c.Request.Context(), beneficiary); err != nil {
		// Check for specific error types
		if _, ok := err.(*service.ValidationError); ok {
			utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
			return
		}
		if _, ok := err.(*service.ConflictError); ok {
			utils.BadRequest(c, "BENEFICIARY_EXISTS", err.Error())
			return
		}
		utils.InternalServerError(c, "BENEFICIARY_CREATION_FAILED", "Failed to create beneficiary")
		return
	}

	c.JSON(http.StatusCreated, beneficiary)
}

// UpdateBeneficiary handles updating a beneficiary
func UpdateBeneficiary(c *gin.Context, serviceFactory *service.Factory) {
	var req UpdateBeneficiaryRequest
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

	beneficiaryID := c.Param("id")
	if beneficiaryID == "" {
		utils.BadRequest(c, "MISSING_BENEFICIARY_ID", "Beneficiary ID is required")
		return
	}

	beneficiaryService := serviceFactory.GetBeneficiaryService()

	// Get existing beneficiary
	beneficiary, err := beneficiaryService.GetBeneficiaryByID(c.Request.Context(), beneficiaryID)
	if err != nil {
		utils.NotFound(c, "BENEFICIARY_NOT_FOUND", "Beneficiary not found")
		return
	}

	// Check if user owns this beneficiary
	if beneficiary.UserID != userID.(string) {
		utils.Forbidden(c, "ACCESS_DENIED", "You don't have access to this beneficiary")
		return
	}

	// Update beneficiary fields if provided
	if req.Name != "" {
		beneficiary.Name = req.Name
	}
	if req.Country != "" {
		beneficiary.Country = req.Country
	}
	if req.BankName != "" {
		beneficiary.BankName = req.BankName
	}
	if req.AccountNumber != "" {
		beneficiary.AccountNo = req.AccountNumber
	}
	if req.SwiftCode != "" {
		beneficiary.SwiftCode = req.SwiftCode
	}
	if req.Iban != "" {
		beneficiary.Iban = req.Iban
	}
	if req.BankAddress != "" {
		beneficiary.BankAddress = req.BankAddress
	}
	if req.Currency != "" {
		beneficiary.Currency = req.Currency
	}

	// Save updated beneficiary
	if err := beneficiaryService.UpdateBeneficiary(c.Request.Context(), beneficiary); err != nil {
		utils.InternalServerError(c, "BENEFICIARY_UPDATE_FAILED", "Failed to update beneficiary")
		return
	}

	c.JSON(http.StatusOK, beneficiary)
}

// DeleteBeneficiary handles deleting a beneficiary
func DeleteBeneficiary(c *gin.Context, serviceFactory *service.Factory) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	beneficiaryID := c.Param("id")
	if beneficiaryID == "" {
		utils.BadRequest(c, "MISSING_BENEFICIARY_ID", "Beneficiary ID is required")
		return
	}

	beneficiaryService := serviceFactory.GetBeneficiaryService()

	// Get existing beneficiary
	beneficiary, err := beneficiaryService.GetBeneficiaryByID(c.Request.Context(), beneficiaryID)
	if err != nil {
		utils.NotFound(c, "BENEFICIARY_NOT_FOUND", "Beneficiary not found")
		return
	}

	// Check if user owns this beneficiary
	if beneficiary.UserID != userID.(string) {
		utils.Forbidden(c, "ACCESS_DENIED", "You don't have access to this beneficiary")
		return
	}

	// Delete beneficiary
	if err := beneficiaryService.DeleteBeneficiary(c.Request.Context(), beneficiaryID); err != nil {
		utils.InternalServerError(c, "BENEFICIARY_DELETE_FAILED", "Failed to delete beneficiary")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Beneficiary deleted successfully",
	})
}
