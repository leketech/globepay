package mocks

import (
	"globepay/internal/domain/model"
	"globepay/internal/service"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) GetProfile(userID int64) (*model.User, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.User), args.Error(1)
}

func (m *UserServiceMock) UpdateProfile(userID int64, user *model.User) error {
	args := m.Called(userID, user)
	return args.Error(0)
}

func (m *UserServiceMock) GetVerificationStatus(userID int64) (*service.UserVerification, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*service.UserVerification), args.Error(1)
}

func (m *UserServiceMock) SubmitVerification(userID int64, verification *service.UserVerification) error {
	args := m.Called(userID, verification)
	return args.Error(0)
}

func (m *UserServiceMock) GetAccounts(userID int64) ([]model.Account, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]model.Account), args.Error(1)
}

func (m *UserServiceMock) CreateAccount(userID int64, account *model.Account) error {
	args := m.Called(userID, account)
	return args.Error(0)
}

func (m *UserServiceMock) GetUserByID(id int64) (*model.User, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.User), args.Error(1)
}
