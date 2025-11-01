package service

import (
	"context"
	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/internal/utils"
	"time"
)

// MoneyRequestService handles business logic for money requests
type MoneyRequestService struct {
	moneyRequestRepo *repository.MoneyRequestRepository
	accountRepo      *repository.AccountRepository
}

// NewMoneyRequestService creates a new money request service
func NewMoneyRequestService(
	moneyRequestRepo *repository.MoneyRequestRepository,
	accountRepo *repository.AccountRepository,
) *MoneyRequestService {
	return &MoneyRequestService{
		moneyRequestRepo: moneyRequestRepo,
		accountRepo:      accountRepo,
	}
}

// CreateRequest creates a new money request
func (s *MoneyRequestService) CreateRequest(
	ctx context.Context,
	requesterID, recipientID string,
	amount float64,
	currency, description string,
) (*model.MoneyRequest, error) {
	// Create a new money request
	request := model.NewMoneyRequest(requesterID, recipientID, amount, currency, description)
	
	// Set expiration to 30 days from now
	request.ExpiresAt = time.Now().Add(30 * 24 * time.Hour)
	
	// Save to database
	if err := s.moneyRequestRepo.Create(ctx, request); err != nil {
		return nil, err
	}
	
	return request, nil
}

// CreatePaymentLink creates a payment link for a money request
func (s *MoneyRequestService) CreatePaymentLink(
	ctx context.Context,
	requestID string,
) (string, error) {
	// Generate a unique payment link
	paymentLink := utils.GeneratePaymentLink(requestID)
	
	// Update the request with the payment link
	if err := s.moneyRequestRepo.UpdatePaymentLink(ctx, requestID, paymentLink); err != nil {
		return "", err
	}
	
	return paymentLink, nil
}

// GetRequest retrieves a money request by ID
func (s *MoneyRequestService) GetRequest(ctx context.Context, id string) (*model.MoneyRequest, error) {
	return s.moneyRequestRepo.GetByID(ctx, id)
}

// GetRequestsByRequester retrieves all money requests made by a user
func (s *MoneyRequestService) GetRequestsByRequester(ctx context.Context, requesterID string) ([]*model.MoneyRequest, error) {
	return s.moneyRequestRepo.GetByRequester(ctx, requesterID)
}

// GetRequestsByRecipient retrieves all money requests for a recipient
func (s *MoneyRequestService) GetRequestsByRecipient(ctx context.Context, recipientID string) ([]*model.MoneyRequest, error) {
	return s.moneyRequestRepo.GetByRecipient(ctx, recipientID)
}

// PayRequest processes payment for a money request
func (s *MoneyRequestService) PayRequest(
	ctx context.Context,
	requestID, payerID string,
) error {
	// Get the money request
	request, err := s.moneyRequestRepo.GetByID(ctx, requestID)
	if err != nil {
		return err
	}
	
	// Check if request is still pending
	if request.Status != string(model.MoneyRequestPending) {
		return &ValidationError{Message: "Request is not in pending status"}
	}
	
	// Check if request has expired
	if request.ExpiresAt.Before(time.Now()) {
		return &ValidationError{Message: "Request has expired"}
	}
	
	// Get payer's account
	payerAccount, err := s.accountRepo.GetByUserIDAndCurrency(ctx, payerID, request.Currency)
	if err != nil {
		return err
	}
	
	// Check if payer has sufficient balance
	if payerAccount.Balance < request.Amount {
		return &ValidationError{Message: "Insufficient balance"}
	}
	
	// Get requester's account
	requesterAccount, err := s.accountRepo.GetByUserIDAndCurrency(ctx, request.RequesterID, request.Currency)
	if err != nil {
		return err
	}
	
	// Begin transaction
	// In a real implementation, you would use a database transaction here
	
	// Deduct amount from payer's account
	payerAccount.Balance -= request.Amount
	if err := s.accountRepo.UpdateBalance(ctx, payerAccount.ID, payerAccount.Balance); err != nil {
		return err
	}
	
	// Add amount to requester's account
	requesterAccount.Balance += request.Amount
	if err := s.accountRepo.UpdateBalance(ctx, requesterAccount.ID, requesterAccount.Balance); err != nil {
		// Rollback payer's account update
		payerAccount.Balance += request.Amount
		s.accountRepo.UpdateBalance(ctx, payerAccount.ID, payerAccount.Balance)
		return err
	}
	
	// Update request status to paid
	now := time.Now()
	if err := s.moneyRequestRepo.UpdateStatus(ctx, requestID, string(model.MoneyRequestPaid), &now); err != nil {
		// Rollback both account updates
		payerAccount.Balance += request.Amount
		s.accountRepo.UpdateBalance(ctx, payerAccount.ID, payerAccount.Balance)
		requesterAccount.Balance -= request.Amount
		s.accountRepo.UpdateBalance(ctx, requesterAccount.ID, requesterAccount.Balance)
		return err
	}
	
	return nil
}