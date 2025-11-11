package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"globepay/internal/domain/model"
	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/metrics"
	"globepay/internal/utils"

	"github.com/gin-gonic/gin"
)

// CreateTransferRequest represents the create transfer request body
type CreateTransferRequest struct {
	RecipientName      string  `json:"recipientName" binding:"required"`
	RecipientEmail     string  `json:"recipientEmail,omitempty"`
	RecipientCountry   string  `json:"recipientCountry" binding:"required,len=2"`
	RecipientBankName  string  `json:"recipientBankName" binding:"required"`
	RecipientAccountNo string  `json:"recipientAccountNumber" binding:"required"`
	RecipientSwiftCode string  `json:"recipientSwiftCode,omitempty"`
	SourceCurrency     string  `json:"sourceCurrency" binding:"required,len=3"`
	DestCurrency       string  `json:"destCurrency" binding:"required,len=3"`
	SourceAmount       float64 `json:"sourceAmount" binding:"required,gt=0"`
	Purpose            string  `json:"purpose" binding:"required"`
}

// GetTransfers handles getting user transfers
func GetTransfers(c *gin.Context, serviceFactory *service.ServiceFactory) {
	fmt.Println("GetTransfers called")

	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		fmt.Println("User ID not found in context")
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	fmt.Printf("User ID from context: %v\n", userID)
	fmt.Printf("User ID type: %T\n", userID)

	// Get pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Ensure reasonable limits
	if limit > 100 {
		limit = 100
	}
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	transferService := serviceFactory.GetTransferService()
	transfers, err := transferService.GetTransfersByUser(c.Request.Context(), userID.(string), limit, offset)
	if err != nil {
		utils.InternalServerError(c, "TRANSFERS_NOT_FOUND", "Failed to retrieve transfers")
		return
	}

	// In a real implementation, you would also return pagination info
	c.JSON(http.StatusOK, gin.H{
		"transfers": transfers,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": len(transfers), // In real implementation, this would be total count
		},
	})
}

// GetTransfer handles getting a specific transfer
func GetTransfer(c *gin.Context, serviceFactory *service.ServiceFactory) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	transferID := c.Param("id")
	if transferID == "" {
		utils.BadRequest(c, "MISSING_TRANSFER_ID", "Transfer ID is required")
		return
	}

	transferService := serviceFactory.GetTransferService()
	transfer, err := transferService.GetTransferByID(c.Request.Context(), transferID)
	if err != nil {
		utils.NotFound(c, "TRANSFER_NOT_FOUND", "Transfer not found")
		return
	}

	// Check if user owns this transfer
	if transfer.UserID != userID.(string) {
		utils.Forbidden(c, "ACCESS_DENIED", "You don't have access to this transfer")
		return
	}

	c.JSON(http.StatusOK, transfer)
}

// CreateTransfer handles creating a new transfer
func CreateTransfer(c *gin.Context, serviceFactory *service.ServiceFactory) {
	var req CreateTransferRequest
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

	transferService := serviceFactory.GetTransferService()

	// Create transfer model
	transfer := &model.Transfer{
		UserID:           userID.(string),
		RecipientName:    req.RecipientName,
		RecipientCountry: req.RecipientCountry,
		SourceCurrency:   req.SourceCurrency,
		DestCurrency:     req.DestCurrency,
		SourceAmount:     req.SourceAmount,
		Purpose:          req.Purpose,
	}

	// Create transfer
	if err := transferService.CreateTransfer(c.Request.Context(), transfer); err != nil {
		// Check for specific error types
		if _, ok := err.(*service.ValidationError); ok {
			utils.BadRequest(c, "VALIDATION_ERROR", err.Error())
			return
		}
		if _, ok := err.(*service.NotFoundError); ok {
			utils.NotFound(c, "ACCOUNT_NOT_FOUND", err.Error())
			return
		}
		if _, ok := err.(*service.InsufficientFundsError); ok {
			utils.BadRequest(c, "INSUFFICIENT_FUNDS", err.Error())
			return
		}
		utils.InternalServerError(c, "TRANSFER_CREATION_FAILED", "Failed to create transfer")
		return
	}

	// Increment transfer metrics
	if metricsInterface, exists := c.Get("metrics"); exists {
		if m, ok := metricsInterface.(*metrics.Metrics); ok {
			m.TransfersTotal.Inc()
			m.TransferAmountTotal.Add(transfer.SourceAmount)
		}
	}

	c.JSON(http.StatusCreated, transfer)
}

// CancelTransfer handles cancelling a transfer
func CancelTransfer(c *gin.Context, serviceFactory *service.ServiceFactory) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

	transferID := c.Param("id")
	if transferID == "" {
		utils.BadRequest(c, "MISSING_TRANSFER_ID", "Transfer ID is required")
		return
	}

	transferService := serviceFactory.GetTransferService()

	// Check if user owns this transfer
	transfer, err := transferService.GetTransferByID(c.Request.Context(), transferID)
	if err != nil {
		utils.NotFound(c, "TRANSFER_NOT_FOUND", "Transfer not found")
		return
	}

	if transfer.UserID != userID.(string) {
		utils.Forbidden(c, "ACCESS_DENIED", "You don't have access to this transfer")
		return
	}

	// Cancel transfer
	if err := transferService.CancelTransfer(c.Request.Context(), transferID); err != nil {
		if _, ok := err.(*service.ConflictError); ok {
			utils.BadRequest(c, "TRANSFER_CANNOT_CANCEL", err.Error())
			return
		}
		utils.InternalServerError(c, "TRANSFER_CANCEL_FAILED", "Failed to cancel transfer")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer cancelled successfully",
	})
}

// GetPublicExchangeRates handles getting exchange rates without authentication
func GetPublicExchangeRates(c *gin.Context, serviceFactory *service.ServiceFactory) {
	from := c.Query("from")
	to := c.Query("to")
	amountStr := c.Query("amount")

	if from == "" || to == "" || amountStr == "" {
		utils.BadRequest(c, "MISSING_PARAMETERS", "from, to, and amount parameters are required")
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		utils.BadRequest(c, "INVALID_AMOUNT", "Invalid amount parameter")
		return
	}

	currencyService := serviceFactory.GetCurrencyService()
	rateResp, err := currencyService.GetExchangeRate(c.Request.Context(), from, to, amount)
	if err != nil {
		utils.InternalServerError(c, "RATE_FETCH_FAILED", "Failed to fetch exchange rate")
		return
	}

	c.JSON(http.StatusOK, rateResp)
}

// GetExchangeRates handles getting exchange rates (protected endpoint)
func GetExchangeRates(c *gin.Context, serviceFactory *service.ServiceFactory) {
	from := c.Query("from")
	to := c.Query("to")
	amountStr := c.Query("amount")

	if from == "" || to == "" || amountStr == "" {
		utils.BadRequest(c, "MISSING_PARAMETERS", "from, to, and amount parameters are required")
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		utils.BadRequest(c, "INVALID_AMOUNT", "Invalid amount parameter")
		return
	}

	currencyService := serviceFactory.GetCurrencyService()
	rateResp, err := currencyService.GetExchangeRate(c.Request.Context(), from, to, amount)
	if err != nil {
		utils.InternalServerError(c, "RATE_FETCH_FAILED", "Failed to fetch exchange rate")
		return
	}

	c.JSON(http.StatusOK, rateResp)
}
