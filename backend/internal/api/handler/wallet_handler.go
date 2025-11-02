package handler

import (
	"fmt"
	"globepay/internal/domain/model"
	"globepay/internal/domain/service"
	"globepay/internal/infrastructure/metrics"
	"globepay/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// WalletHandler handles wallet-related API requests
type WalletHandler struct {
	serviceFactory *service.ServiceFactory
	metrics        *metrics.Metrics
}

// NewWalletHandler creates a new wallet handler
func NewWalletHandler(serviceFactory *service.ServiceFactory, metrics *metrics.Metrics) *WalletHandler {
	return &WalletHandler{
		serviceFactory: serviceFactory,
		metrics:        metrics,
	}
}

// AddMoneyRequest represents the request body for adding money to wallet
type AddMoneyRequest struct {
	Amount       float64 `json:"amount" binding:"required,gt=0"`
	PaymentMethod string  `json:"payment_method" binding:"required,oneof=card bank"`
	CardNumber   string  `json:"card_number,omitempty"`
	ExpiryDate   string  `json:"expiry_date,omitempty"`
	CVV          string  `json:"cvv,omitempty"`
	AccountNumber string  `json:"account_number,omitempty"`
	RoutingNumber string  `json:"routing_number,omitempty"`
}

// AddMoneyResponse represents the response for adding money to wallet
type AddMoneyResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Balance float64 `json:"balance,omitempty"`
}

// AddMoney handles adding money to user's wallet
func (h *WalletHandler) AddMoney(c *gin.Context) {
	fmt.Println("Handling /api/v1/wallet/add request")
	
	var req AddMoneyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	
	// Get account service
	accountService := h.serviceFactory.GetAccountService()
	
	// In a real implementation, you would:
	// 1. Validate payment details
	// 2. Process payment with payment gateway or bank API
	// 3. Update user's account balance
	// 4. Record transaction
	
	// For now, we'll just simulate adding money to the account
	// Get user's account (assuming USD for demo)
	account, err := accountService.GetAccountByUserIDAndCurrency(c.Request.Context(), userID.(string), "USD")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get account"})
		return
	}
	
	// Update account balance
	newBalance := account.Balance + req.Amount
	if err := accountService.UpdateAccountBalance(c.Request.Context(), account.ID, newBalance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update account balance"})
		return
	}
	
	// Record transaction (simplified)
	transactionService := h.serviceFactory.GetTransactionService()
	
	// Create a transaction object
	transaction := &model.Transaction{
		AccountID: account.ID,
		Type:      "credit",
		Amount:    req.Amount,
		Currency:  "USD",
		ReferenceNumber: fmt.Sprintf("ADD-MONEY-%s", utils.GenerateUUID()[:8]),
		Description: fmt.Sprintf("Added money via %s", req.PaymentMethod),
	}
	
	err = transactionService.CreateTransaction(c.Request.Context(), transaction)
	if err != nil {
		// Log error but don't fail the request
		fmt.Printf("Failed to record transaction: %v\n", err)
	}
	
	// Update metrics
	h.metrics.TransfersTotal.Inc()
	h.metrics.TransferAmountTotal.Add(req.Amount)
	
	c.JSON(http.StatusOK, AddMoneyResponse{
		Success: true,
		Message: fmt.Sprintf("Successfully added $%.2f to your wallet", req.Amount),
		Balance: newBalance,
	})
}

// RequestMoneyRequest represents the request body for requesting money
type RequestMoneyRequest struct {
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	RecipientID string  `json:"recipient_id,omitempty"`
	Description string  `json:"description,omitempty"`
	IsLink      bool    `json:"is_link"`
}

// RequestMoneyResponse represents the response for requesting money
type RequestMoneyResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	RequestID   string `json:"request_id,omitempty"`
	PaymentLink string `json:"payment_link,omitempty"`
}

// RequestMoney handles requesting money from another user
func (h *WalletHandler) RequestMoney(c *gin.Context) {
	fmt.Println("Handling /api/v1/wallet/request request")
	
	var req RequestMoneyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Get user ID from context (set by auth middleware)
	requesterID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	
	// Get money request service
	moneyRequestService := h.serviceFactory.GetMoneyRequestService()
	
	var recipientID string
	if req.IsLink {
		// For payment links, recipient is the requester themselves
		recipientID = requesterID.(string)
	} else {
		// Validate recipient ID is provided
		if req.RecipientID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Recipient ID is required"})
			return
		}
		recipientID = req.RecipientID
	}
	
	// Create money request
	moneyRequest, err := moneyRequestService.CreateRequest(
		c.Request.Context(),
		requesterID.(string),
		recipientID,
		req.Amount,
		"USD", // Assuming USD for demo
		req.Description,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create money request"})
		return
	}
	
	var paymentLink string
	if req.IsLink {
		// Generate payment link
		paymentLink, err = moneyRequestService.CreatePaymentLink(c.Request.Context(), moneyRequest.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate payment link"})
			return
		}
	}
	
	c.JSON(http.StatusOK, RequestMoneyResponse{
		Success:     true,
		Message:     "Money request created successfully",
		RequestID:   moneyRequest.ID,
		PaymentLink: paymentLink,
	})
}

// GetMoneyRequestsResponse represents the response for getting money requests
type GetMoneyRequestsResponse struct {
	Success   bool                `json:"success"`
	Requests  []*MoneyRequestInfo `json:"requests"`
	Timestamp time.Time           `json:"timestamp"`
}

// MoneyRequestInfo represents a money request in the response
type MoneyRequestInfo struct {
	ID            string    `json:"id"`
	RequesterID   string    `json:"requester_id"`
	RecipientID   string    `json:"recipient_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	PaymentLink   string    `json:"payment_link,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	ExpiresAt     time.Time `json:"expires_at"`
	PaidAt        time.Time `json:"paid_at,omitempty"`
}

// GetMoneyRequests handles getting money requests for the user
func (h *WalletHandler) GetMoneyRequests(c *gin.Context) {
	fmt.Println("Handling /api/v1/wallet/requests request")
	
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	
	// Get money request service
	moneyRequestService := h.serviceFactory.GetMoneyRequestService()
	
	// Get requests where user is requester or recipient
	requestsAsRequester, err := moneyRequestService.GetRequestsByRequester(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get money requests"})
		return
	}
	
	requestsAsRecipient, err := moneyRequestService.GetRequestsByRecipient(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get money requests"})
		return
	}
	
	// Combine requests
	allRequests := append(requestsAsRequester, requestsAsRecipient...)
	
	// Convert to response format
	requestInfos := make([]*MoneyRequestInfo, len(allRequests))
	for i, req := range allRequests {
		requestInfos[i] = &MoneyRequestInfo{
			ID:          req.ID,
			RequesterID: req.RequesterID,
			RecipientID: req.RecipientID,
			Amount:      req.Amount,
			Currency:    req.Currency,
			Description: req.Description,
			Status:      req.Status,
			PaymentLink: req.PaymentLink,
			CreatedAt:   req.CreatedAt,
			ExpiresAt:   req.ExpiresAt,
			PaidAt:      req.PaidAt,
		}
	}
	
	c.JSON(http.StatusOK, GetMoneyRequestsResponse{
		Success:   true,
		Requests:  requestInfos,
		Timestamp: time.Now(),
	})
}