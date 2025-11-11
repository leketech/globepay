package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"globepay/internal/domain/model"
)

type BeneficiaryRepositoryMock struct {
	mock.Mock
}

func (m *BeneficiaryRepositoryMock) Create(beneficiary *model.Beneficiary) error {
	args := m.Called(beneficiary)
	return args.Error(0)
}

func (m *BeneficiaryRepositoryMock) GetByID(id string) (*model.Beneficiary, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Beneficiary), args.Error(1)
}

func (m *BeneficiaryRepositoryMock) GetByUser(ctx context.Context, userID string) ([]*model.Beneficiary, error) {
	args := m.Called(ctx, userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.Beneficiary), args.Error(1)
}

func (m *BeneficiaryRepositoryMock) GetByNameAndUser(ctx context.Context, name, userID string) (*model.Beneficiary, error) {
	args := m.Called(ctx, name, userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Beneficiary), args.Error(1)
}

func (m *BeneficiaryRepositoryMock) Update(beneficiary *model.Beneficiary) error {
	args := m.Called(beneficiary)
	return args.Error(0)
}

func (m *BeneficiaryRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
