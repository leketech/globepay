package mocks

import (
	"context"
	"globepay/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetByID(id string) (*model.User, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetAll() ([]model.User, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]model.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByUserAndCurrency(ctx context.Context, userID, currency string) (*model.Account, error) {
	args := m.Called(ctx, userID, currency)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Account), args.Error(1)
}