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
	GetProfile(userID int64) (*model.User, error)
	UpdateProfile(userID int64, user *model.User) error
	GetVerificationStatus(userID int64) (*UserVerification, error)
	SubmitVerification(userID int64, verification *UserVerification) error
	GetAccounts(userID int64) ([]model.Account, error)
	CreateAccount(userID int64, account *model.Account) error
	GetUserByID(id int64) (*model.User, error)
}

// UserVerification represents user verification status
type UserVerification struct {
	UserID          int64     `json:"user_id" db:"user_id"`
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
func (s *UserService) GetProfile(userID int64) (*model.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// UpdateProfile updates a user's profile
func (s *UserService) UpdateProfile(userID int64, user *model.User) error {
	// Get existing user
	existingUser, err := s.userRepo.GetByID(userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Update fields
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Phone = user.Phone
	existingUser.Country = user.Country
	existingUser.Currency = user.Currency
	existingUser.UpdatedAt = time.Now()

	// Update user
	if err := s.userRepo.Update(existingUser); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// GetVerificationStatus retrieves a user's verification status
func (s *UserService) GetVerificationStatus(userID int64) (*UserVerification, error) {
	// In a real implementation, you would retrieve this from the database
	// For now, we'll return a placeholder
	verification := &UserVerification{
		UserID:        userID,
		EmailVerified: true,
		PhoneVerified: false,
		IDVerified:    false,
		AddressVerified: false,
		KYCLevel:      1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return verification, nil
}

// SubmitVerification submits user verification documents
func (s *UserService) SubmitVerification(userID int64, verification *UserVerification) error {
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
func (s *UserService) GetAccounts(userID int64) ([]model.Account, error) {
	accounts, err := s.accountRepo.GetByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}

	return accounts, nil
}

// CreateAccount creates a new account for a user
func (s *UserService) CreateAccount(userID int64, account *model.Account) error {
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
func (s *UserService) GetUserByID(id int64) (*model.User, error) {
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
func (s *UserService) UpdateUserStatus(userID int64, status string) error {
	// Get existing user
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Update status
	user.Status = status
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