package service

import (
	"fmt"
	"time"

	"globepay/internal/domain"
	"globepay/internal/repository"
	"globepay/internal/domain/model"
)

// UserServiceInterface defines the interface for user service
type UserServiceInterface interface {
	GetProfile(userID string) (*model.User, error) // Changed from int64 to string
	UpdateProfile(userID string, user *model.User) error // Changed from int64 to string
	GetVerificationStatus(userID string) (*UserVerification, error) // Changed from int64 to string
	SubmitVerification(userID string, verification *UserVerification) error // Changed from int64 to string
	GetAccounts(userID string) ([]model.Account, error) // Changed from int64 to string
	CreateAccount(userID string, account *model.Account) error // Changed from int64 to string
	GetUserByID(id string) (*model.User, error) // Changed from int64 to string
}

// UserVerification represents user verification status
type UserVerification struct {
	UserID          string    `json:"user_id" db:"user_id"` // Changed from int64 to string
	EmailVerified   bool      `json:"email_verified" db:"email_verified"`
	PhoneVerified   bool      `json:"phone_verified" db:"phone_verified"`
	IDVerified      bool      `json:"id_verified" db:"id_verified"`
	AddressVerified bool      `json:"address_verified" db:"address_verified"`
	KYCLevel        int       `json:"kyc_level" db:"kyc_level"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// UserService implements UserServiceInterface
type UserService struct {
	userRepo    repository.UserRepository
	accountRepo repository.AccountRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepo repository.UserRepository, accountRepo repository.AccountRepository) *UserService {
	return &UserService{
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
}

// GetProfile retrieves a user's profile
func (s *UserService) GetProfile(userID string) (*model.User, error) { // Changed from int64 to string
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// UpdateProfile updates a user's profile
func (s *UserService) UpdateProfile(userID string, user *model.User) error { // Changed from int64 to string
	// Get existing user
	existingUser, err := s.userRepo.GetByID(userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Update fields
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.PhoneNumber = user.PhoneNumber // Fixed field name
	existingUser.Country = user.Country
	existingUser.UpdatedAt = time.Now()

	// Update user
	if err := s.userRepo.Update(existingUser); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// GetVerificationStatus retrieves a user's verification status
func (s *UserService) GetVerificationStatus(userID string) (*UserVerification, error) { // Changed from int64 to string
	// In a real implementation, you would retrieve this from the database
	// For now, we'll return a placeholder
	verification := &UserVerification{
		UserID:          userID,
		EmailVerified:   true,
		PhoneVerified:   false,
		IDVerified:      false,
		AddressVerified: false,
		KYCLevel:        1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	return verification, nil
}

// SubmitVerification submits user verification documents
func (s *UserService) SubmitVerification(userID string, verification *UserVerification) error { // Changed from int64 to string
	// In a real implementation, you would save this to the database
	// and process the verification documents
	verification.UserID = userID
	verification.CreatedAt = time.Now()
	verification.UpdatedAt = time.Now()

	// Update user verification status
	// This would involve database operations in a real implementation

	return nil
}

// GetAccounts retrieves all accounts for a user
func (s *UserService) GetAccounts(userID string) ([]model.Account, error) { // Changed from int64 to string
	// This method doesn't exist in the repository interface, so we'll use GetByUser instead
	accountPtrs, err := s.accountRepo.GetByUser(nil, userID) // Pass nil context for now
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}
	
	// Convert []*model.Account to []model.Account
	accounts := make([]model.Account, len(accountPtrs))
	for i, accountPtr := range accountPtrs {
		if accountPtr != nil {
			accounts[i] = *accountPtr
		}
	}

	return accounts, nil
}

// CreateAccount creates a new account for a user
func (s *UserService) CreateAccount(userID string, account *model.Account) error { // Changed from int64 to string
	// Set user ID
	account.UserID = userID

	// Generate account number
	account.AccountNumber = s.generateAccountNumber()

	// Set default values
	account.Balance = 0.0
	account.Status = string(domain.AccountActive)
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	// Create account
	if err := s.accountRepo.Create(account); err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	return nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id string) (*model.User, error) { // Changed from int64 to string
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// generateAccountNumber generates a unique account number
func (s *UserService) generateAccountNumber() string {
	// In a real implementation, you would generate a unique account number
	// based on your business requirements
	// For now, we'll generate a simple placeholder
	return fmt.Sprintf("ACC%012d", time.Now().UnixNano()%1000000000000)
}

// UpdateUserStatus updates a user's status
func (s *UserService) UpdateUserStatus(userID string, status string) error { // Changed from int64 to string
	// Get existing user
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Update status
	user.AccountStatus = status // Fixed field name
	user.UpdatedAt = time.Now()

	// Update user
	if err := s.userRepo.Update(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}