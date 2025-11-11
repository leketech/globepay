package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"globepay/internal/domain/model"
)

type UserPreferencesRepositoryMock struct {
	mock.Mock
}

func (m *UserPreferencesRepositoryMock) GetUserPreferences(ctx context.Context, userID string) (*model.UserPreferences, error) {
	args := m.Called(ctx, userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.UserPreferences), args.Error(1)
}

func (m *UserPreferencesRepositoryMock) CreateUserPreferences(ctx context.Context, preferences *model.UserPreferences) error {
	args := m.Called(ctx, preferences)
	return args.Error(0)
}

func (m *UserPreferencesRepositoryMock) UpdateUserPreferences(ctx context.Context, preferences *model.UserPreferences) error {
	args := m.Called(ctx, preferences)
	return args.Error(0)
}
