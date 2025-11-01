package mocks

import (
	"globepay/internal/domain"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) Create(account *domain.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountRepositoryMock) GetByID(id int64) (*domain.Account, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.Account), args.Error(1)
}

func (m *AccountRepositoryMock) GetByUserID(userID int64) ([]domain.Account, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.Account), args.Error(1)
}

func (m *AccountRepositoryMock) Update(account *domain.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountRepositoryMock) UpdateBalance(id int64, balance float64) error {
	args := m.Called(id, balance)
	return args.Error(0)
}

func (m *AccountRepositoryMock) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}