package integration

import (
	"context"
	"testing"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/infrastructure/database"
	"globepay/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DatabaseTestSuite struct {
	suite.Suite
	db     *database.Database
	userRepo repository.UserRepository
}

func (suite *DatabaseTestSuite) SetupSuite() {
	// In a real test, you would connect to a test database
	// For now, we'll skip the actual database connection
	suite.T().Skip("Skipping database tests - no test database configured")
}

func (suite *DatabaseTestSuite) TestUserRepository_Create() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test user
	user := model.NewUser("test@example.com", "password123", "John", "Doe")
	
	// Test creating user
	err := suite.userRepo.Create(context.Background(), user)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), user.ID)
}

func (suite *DatabaseTestSuite) TestUserRepository_GetByID() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test2@example.com", "password123", "Jane", "Doe")
	err := suite.userRepo.Create(context.Background(), user)
	assert.NoError(suite.T(), err)

	// Then retrieve the user
	retrievedUser, err := suite.userRepo.GetByID(context.Background(), user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Email, retrievedUser.Email)
	assert.Equal(suite.T(), user.FirstName, retrievedUser.FirstName)
	assert.Equal(suite.T(), user.LastName, retrievedUser.LastName)
}

func (suite *DatabaseTestSuite) TestUserRepository_GetByEmail() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test3@example.com", "password123", "Bob", "Smith")
	err := suite.userRepo.Create(context.Background(), user)
	assert.NoError(suite.T(), err)

	// Then retrieve the user by email
	retrievedUser, err := suite.userRepo.GetByEmail(context.Background(), user.Email)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, retrievedUser.ID)
	assert.Equal(suite.T(), user.FirstName, retrievedUser.FirstName)
	assert.Equal(suite.T(), user.LastName, retrievedUser.LastName)
}

func (suite *DatabaseTestSuite) TestUserRepository_Update() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test4@example.com", "password123", "Alice", "Johnson")
	err := suite.userRepo.Create(context.Background(), user)
	assert.NoError(suite.T(), err)

	// Update user information
	user.FirstName = "Alicia"
	user.LastName = "Johnson-Smith"
	user.UpdatedAt = time.Now()

	// Save updated user
	err = suite.userRepo.Update(context.Background(), user)
	assert.NoError(suite.T(), err)

	// Retrieve updated user
	updatedUser, err := suite.userRepo.GetByID(context.Background(), user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Alicia", updatedUser.FirstName)
	assert.Equal(suite.T(), "Johnson-Smith", updatedUser.LastName)
}

func (suite *DatabaseTestSuite) TestUserRepository_Delete() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test5@example.com", "password123", "Charlie", "Brown")
	err := suite.userRepo.Create(context.Background(), user)
	assert.NoError(suite.T(), err)

	// Delete the user
	err = suite.userRepo.Delete(context.Background(), user.ID)
	assert.NoError(suite.T(), err)

	// Try to retrieve the deleted user (should fail)
	_, err = suite.userRepo.GetByID(context.Background(), user.ID)
	assert.Error(suite.T(), err)
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}