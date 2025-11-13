package service

import (
	"context"
	"database/sql"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/internal/utils"
)

// BeneficiaryService provides beneficiary-related functionality
type BeneficiaryService struct {
	beneficiaryRepo repository.BeneficiaryRepository
}

// NewBeneficiaryService creates a new beneficiary service
func NewBeneficiaryService(beneficiaryRepo repository.BeneficiaryRepository) *BeneficiaryService {
	return &BeneficiaryService{
		beneficiaryRepo: beneficiaryRepo,
	}
}

// CreateBeneficiary creates a new beneficiary
func (s *BeneficiaryService) CreateBeneficiary(ctx context.Context, beneficiary *model.Beneficiary) error {
	// Validate country code
	if !utils.ValidateCountryCode(beneficiary.Country) {
		return &ValidationError{Field: "country", Message: "Invalid country code"}
	}

	// Check if beneficiary already exists for this user
	existingBeneficiary, err := s.beneficiaryRepo.GetByNameAndUser(ctx, beneficiary.Name, beneficiary.UserID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existingBeneficiary != nil {
		return &ConflictError{Message: "Beneficiary with this name already exists"}
	}

	// Save beneficiary
	if err := s.beneficiaryRepo.Create(beneficiary); err != nil {
		return err
	}

	return nil
}

// GetBeneficiariesByUser retrieves beneficiaries for a user
func (s *BeneficiaryService) GetBeneficiariesByUser(ctx context.Context, userID string) ([]*model.Beneficiary, error) {
	return s.beneficiaryRepo.GetByUser(ctx, userID)
}

// GetBeneficiaryByID retrieves a beneficiary by ID
func (s *BeneficiaryService) GetBeneficiaryByID(_ context.Context, beneficiaryID string) (*model.Beneficiary, error) {
	return s.beneficiaryRepo.GetByID(beneficiaryID)
}

// UpdateBeneficiary updates beneficiary information
func (s *BeneficiaryService) UpdateBeneficiary(_ context.Context, beneficiary *model.Beneficiary) error {
	beneficiary.UpdatedAt = time.Now()
	return s.beneficiaryRepo.Update(beneficiary)
}

// DeleteBeneficiary deletes a beneficiary
func (s *BeneficiaryService) DeleteBeneficiary(_ context.Context, beneficiaryID string) error {
	return s.beneficiaryRepo.Delete(beneficiaryID)
}
