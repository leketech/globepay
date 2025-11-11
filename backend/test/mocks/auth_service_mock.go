package mocks

import (
	"github.com/stretchr/testify/mock"
	"globepay/internal/domain"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) Register(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *AuthServiceMock) RefreshToken(tokenString string) (string, error) {
	args := m.Called(tokenString)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) GenerateOTP() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) ValidateOTP(otp, hashedOTP string) bool {
	args := m.Called(otp, hashedOTP)
	return args.Bool(0)
}

func (m *AuthServiceMock) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}
