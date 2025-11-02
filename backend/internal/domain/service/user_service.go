package service

import (
	"context"
	"globepay/internal/domain/model"
	"globepay/internal/repository"
)

// UserService handles user-related business logic
type UserService struct {
	userRepo            repository.UserRepository
	userPreferencesRepo repository.UserPreferencesRepository
}

// NewUserService creates a new user service
func NewUserService(
	userRepo repository.UserRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// SetUserPreferencesRepo sets the user preferences repository (used for dependency injection)
func (s *UserService) SetUserPreferencesRepo(repo repository.UserPreferencesRepository) {
	s.userPreferencesRepo = repo
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.Create(user)
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepo.GetByEmail(email)
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.Update(user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.Delete(id)
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.userRepo.GetAll()
}

// GetUserPreferences retrieves user preferences
func (s *UserService) GetUserPreferences(ctx context.Context, userID string) (*model.UserPreferences, error) {
	if s.userPreferencesRepo == nil {
		return nil, &NotFoundError{Message: "User preferences repository not initialized"}
	}
	return s.userPreferencesRepo.GetUserPreferences(ctx, userID)
}

// UpdateUserPreferences updates user preferences
func (s *UserService) UpdateUserPreferences(ctx context.Context, preferences *model.UserPreferences) error {
	if s.userPreferencesRepo == nil {
		return &NotFoundError{Message: "User preferences repository not initialized"}
	}
	return s.userPreferencesRepo.UpdateUserPreferences(ctx, preferences)
}