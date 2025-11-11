package service

import (
	"fmt"
	"time"

	"globepay/internal/domain"
)

// KYCServiceInterface defines the interface for KYC service
type KYCServiceInterface interface {
	SubmitKYCApplication(userID int64, application *KYCApplication) error
	GetKYCStatus(userID int64) (*KYCStatus, error)
	VerifyIdentity(userID int64, document *IdentityDocument) error
	VerifyAddress(userID int64, document *AddressDocument) error
	VerifyIncome(userID int64, document *IncomeDocument) error
	UpdateKYCLevel(userID int64, level int) error
	GetKYCApplication(userID int64) (*KYCApplication, error)
}

// KYCApplication represents a KYC application
type KYCApplication struct {
	UserID      int64             `json:"user_id"`
	Level       int               `json:"level"`
	Status      string            `json:"status"`
	SubmittedAt time.Time         `json:"submitted_at"`
	ReviewedAt  time.Time         `json:"reviewed_at"`
	ReviewerID  *int64            `json:"reviewer_id"`
	ReviewNotes string            `json:"review_notes"`
	IdentityDoc *IdentityDocument `json:"identity_doc"`
	AddressDoc  *AddressDocument  `json:"address_doc"`
	IncomeDoc   *IncomeDocument   `json:"income_doc"`
}

// KYCStatus represents the KYC status of a user
type KYCStatus struct {
	UserID      int64     `json:"user_id"`
	Level       int       `json:"level"`
	Status      string    `json:"status"`
	LastChecked time.Time `json:"last_checked"`
}

// IdentityDocument represents an identity document
type IdentityDocument struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	DocumentType   string    `json:"document_type"` // passport, driver_license, id_card
	DocumentNumber string    `json:"document_number"`
	IssueDate      time.Time `json:"issue_date"`
	ExpiryDate     time.Time `json:"expiry_date"`
	FrontImageURL  string    `json:"front_image_url"`
	BackImageURL   string    `json:"back_image_url"`
	SelfieImageURL string    `json:"selfie_image_url"`
	Verified       bool      `json:"verified"`
	VerifiedAt     time.Time `json:"verified_at"`
	VerifiedBy     *int64    `json:"verified_by"`
}

// AddressDocument represents an address document
type AddressDocument struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	DocumentType string    `json:"document_type"` // utility_bill, bank_statement, lease_agreement
	DocumentDate time.Time `json:"document_date"`
	Address      string    `json:"address"`
	ImageURL     string    `json:"image_url"`
	Verified     bool      `json:"verified"`
	VerifiedAt   time.Time `json:"verified_at"`
	VerifiedBy   *int64    `json:"verified_by"`
}

// IncomeDocument represents an income document
type IncomeDocument struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	DocumentType string    `json:"document_type"` // payslip, tax_return, bank_statement
	DocumentDate time.Time `json:"document_date"`
	IncomeAmount float64   `json:"income_amount"`
	Currency     string    `json:"currency"`
	ImageURL     string    `json:"image_url"`
	Verified     bool      `json:"verified"`
	VerifiedAt   time.Time `json:"verified_at"`
	VerifiedBy   *int64    `json:"verified_by"`
}

// KYCService implements KYCServiceInterface
type KYCService struct {
	// In a real implementation, you would have repositories for KYC data
	// For now, we'll use in-memory storage
	kycApplications map[int64]*KYCApplication
	kycStatuses     map[int64]*KYCStatus
}

// NewKYCService creates a new KYCService
func NewKYCService() *KYCService {
	return &KYCService{
		kycApplications: make(map[int64]*KYCApplication),
		kycStatuses:     make(map[int64]*KYCStatus),
	}
}

// SubmitKYCApplication submits a KYC application
func (s *KYCService) SubmitKYCApplication(userID int64, application *KYCApplication) error {
	// Validate application
	if application.Level < 1 || application.Level > 3 {
		return fmt.Errorf("invalid KYC level: %d", application.Level)
	}

	// Set application details
	application.UserID = userID
	application.Status = "pending"
	application.SubmittedAt = time.Now()

	// Store application
	s.kycApplications[userID] = application

	// Update KYC status
	status := &KYCStatus{
		UserID:      userID,
		Level:       application.Level,
		Status:      "pending",
		LastChecked: time.Now(),
	}
	s.kycStatuses[userID] = status

	return nil
}

// GetKYCStatus retrieves the KYC status of a user
func (s *KYCService) GetKYCStatus(userID int64) (*KYCStatus, error) {
	status, exists := s.kycStatuses[userID]
	if !exists {
		return nil, domain.ErrKYCRequired
	}

	return status, nil
}

// VerifyIdentity verifies a user's identity document
func (s *KYCService) VerifyIdentity(userID int64, document *IdentityDocument) error {
	// Get existing application
	application, exists := s.kycApplications[userID]
	if !exists {
		return fmt.Errorf("no KYC application found for user %d", userID)
	}

	// Validate document
	if document.DocumentType == "" || document.DocumentNumber == "" {
		return fmt.Errorf("invalid identity document")
	}

	// Set document details
	document.UserID = userID
	document.Verified = true
	document.VerifiedAt = time.Now()

	// Update application
	application.IdentityDoc = document
	application.ReviewedAt = time.Now()

	// If all required documents are verified, update status
	if s.isApplicationComplete(application) {
		application.Status = "approved"
		status := s.kycStatuses[userID]
		status.Status = "approved"
		status.LastChecked = time.Now()
	}

	return nil
}

// VerifyAddress verifies a user's address document
func (s *KYCService) VerifyAddress(userID int64, document *AddressDocument) error {
	// Get existing application
	application, exists := s.kycApplications[userID]
	if !exists {
		return fmt.Errorf("no KYC application found for user %d", userID)
	}

	// Validate document
	if document.DocumentType == "" || document.Address == "" {
		return fmt.Errorf("invalid address document")
	}

	// Set document details
	document.UserID = userID
	document.Verified = true
	document.VerifiedAt = time.Now()

	// Update application
	application.AddressDoc = document
	application.ReviewedAt = time.Now()

	// If all required documents are verified, update status
	if s.isApplicationComplete(application) {
		application.Status = "approved"
		status := s.kycStatuses[userID]
		status.Status = "approved"
		status.LastChecked = time.Now()
	}

	return nil
}

// VerifyIncome verifies a user's income document
func (s *KYCService) VerifyIncome(userID int64, document *IncomeDocument) error {
	// Get existing application
	application, exists := s.kycApplications[userID]
	if !exists {
		return fmt.Errorf("no KYC application found for user %d", userID)
	}

	// Validate document
	if document.DocumentType == "" || document.IncomeAmount <= 0 {
		return fmt.Errorf("invalid income document")
	}

	// Set document details
	document.UserID = userID
	document.Verified = true
	document.VerifiedAt = time.Now()

	// Update application
	application.IncomeDoc = document
	application.ReviewedAt = time.Now()

	// If all required documents are verified, update status
	if s.isApplicationComplete(application) {
		application.Status = "approved"
		status := s.kycStatuses[userID]
		status.Status = "approved"
		status.LastChecked = time.Now()
	}

	return nil
}

// UpdateKYCLevel updates a user's KYC level
func (s *KYCService) UpdateKYCLevel(userID int64, level int) error {
	// Validate level
	if level < 0 || level > 3 {
		return fmt.Errorf("invalid KYC level: %d", level)
	}

	// Update status
	status, exists := s.kycStatuses[userID]
	if !exists {
		status = &KYCStatus{
			UserID:      userID,
			Level:       level,
			Status:      "not_submitted",
			LastChecked: time.Now(),
		}
		s.kycStatuses[userID] = status
	} else {
		status.Level = level
		status.LastChecked = time.Now()
	}

	return nil
}

// GetKYCApplication retrieves a user's KYC application
func (s *KYCService) GetKYCApplication(userID int64) (*KYCApplication, error) {
	application, exists := s.kycApplications[userID]
	if !exists {
		return nil, fmt.Errorf("no KYC application found for user %d", userID)
	}

	return application, nil
}

// isApplicationComplete checks if a KYC application has all required documents verified
func (s *KYCService) isApplicationComplete(application *KYCApplication) bool {
	// Level 1 requires identity verification
	if application.Level >= 1 && application.IdentityDoc == nil {
		return false
	}

	// Level 2 requires address verification
	if application.Level >= 2 && application.AddressDoc == nil {
		return false
	}

	// Level 3 requires income verification
	if application.Level >= 3 && application.IncomeDoc == nil {
		return false
	}

	// Check if all submitted documents are verified
	if application.IdentityDoc != nil && !application.IdentityDoc.Verified {
		return false
	}

	if application.AddressDoc != nil && !application.AddressDoc.Verified {
		return false
	}

	if application.IncomeDoc != nil && !application.IncomeDoc.Verified {
		return false
	}

	return true
}

// RejectKYCApplication rejects a KYC application
func (s *KYCService) RejectKYCApplication(userID int64, notes string) error {
	// Get existing application
	application, exists := s.kycApplications[userID]
	if !exists {
		return fmt.Errorf("no KYC application found for user %d", userID)
	}

	// Update application status
	application.Status = "rejected"
	application.ReviewedAt = time.Now()
	application.ReviewNotes = notes

	// Update KYC status
	status := s.kycStatuses[userID]
	status.Status = "rejected"
	status.LastChecked = time.Now()

	return nil
}

// GetPendingApplications retrieves all pending KYC applications
func (s *KYCService) GetPendingApplications() ([]*KYCApplication, error) {
	var pending []*KYCApplication

	for _, application := range s.kycApplications {
		if application.Status == "pending" {
			pending = append(pending, application)
		}
	}

	return pending, nil
}
