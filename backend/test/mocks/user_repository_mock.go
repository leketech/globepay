package mocks

import (
	"globepay/internal/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetByID(id int64) (*domain.User, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetAll() ([]domain.User, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]domain.User), args.Error(1)
}