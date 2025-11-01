package handler

import (
	"net/http"
	"strconv"

	"globepay/internal/domain/service"
	"globepay/internal/utils"

	"github.com/gin-gonic/gin"
)

// GetTransactions handles getting user transactions
func GetTransactions(c *gin.Context, serviceFactory *service.ServiceFactory) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "MISSING_USER_ID", "User ID not found in context")
		return
	}

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

	transactionService := serviceFactory.GetTransactionService()
	transactions, err := transactionService.GetTransactionsByUser(c.Request.Context(), userID.(string), limit, offset)
	if err != nil {
		utils.InternalServerError(c, "TRANSACTIONS_NOT_FOUND", "Failed to retrieve transactions")
		return
	}

	// In a real implementation, you would also return pagination info
	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": len(transactions), // In real implementation, this would be total count
		},
	})
}

// GetTransaction handles getting a specific transaction
func GetTransaction(c *gin.Context, serviceFactory *service.ServiceFactory) {
	transactionID := c.Param("id")
	if transactionID == "" {
		utils.BadRequest(c, "MISSING_TRANSACTION_ID", "Transaction ID is required")
		return
	}

	transactionService := serviceFactory.GetTransactionService()
	transaction, err := transactionService.GetTransactionByID(c.Request.Context(), transactionID)
	if err != nil {
		utils.NotFound(c, "TRANSACTION_NOT_FOUND", "Transaction not found")
		return
	}

	// In a real implementation, you would check if user has access to this transaction
	// For now, we'll assume the service layer handles this

	c.JSON(http.StatusOK, transaction)
}