package integration

import (
	"testing"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
	userRepo    repository.UserRepository
	db          *utils.TestDB
	redisClient *utils.TestRedis
}

func (suite *AuthTestSuite) SetupSuite() {
	// Initialize test database
	suite.db = utils.NewTestDB()
	suite.redisClient = utils.NewTestRedis()

	// Initialize repository
	if suite.db != nil {
		suite.userRepo = repository.NewUserRepository(suite.db.DB)
	}
}

func (suite *AuthTestSuite) TearDownSuite() {
	if suite.db != nil {
		suite.db.Close()
	}
	if suite.redisClient != nil {
		suite.redisClient.Close()
	}
}

func (suite *AuthTestSuite) SetupTest() {
	// Clear test data before each test
	if suite.db != nil {
		suite.db.ClearTables()
	}
}

func (suite *AuthTestSuite) TestUserRepository_Create() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test user
	user := model.NewUser("test@example.com", "password123", "John", "Doe")

	// Test creating user
	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), user.ID)
}

func (suite *AuthTestSuite) TestUserRepository_GetByID() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test2@example.com", "password123", "Jane", "Doe")
	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)

	// Then retrieve the user
	retrievedUser, err := suite.userRepo.GetByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Email, retrievedUser.Email)
	assert.Equal(suite.T(), user.FirstName, retrievedUser.FirstName)
	assert.Equal(suite.T(), user.LastName, retrievedUser.LastName)
}

func (suite *AuthTestSuite) TestUserRepository_GetByEmail() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test3@example.com", "password123", "Bob", "Smith")
	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)

	// Then retrieve the user by email
	retrievedUser, err := suite.userRepo.GetByEmail(user.Email)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, retrievedUser.ID)
	assert.Equal(suite.T(), user.FirstName, retrievedUser.FirstName)
	assert.Equal(suite.T(), user.LastName, retrievedUser.LastName)
}

func (suite *AuthTestSuite) TestUserRepository_Update() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test4@example.com", "password123", "Alice", "Johnson")
	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)

	// Update user information
	user.FirstName = "Alicia"
	user.LastName = "Johnson-Smith"
	user.UpdatedAt = user.CreatedAt // Just to satisfy the field

	// Save updated user
	err = suite.userRepo.Update(user)
	assert.NoError(suite.T(), err)

	// Retrieve updated user
	updatedUser, err := suite.userRepo.GetByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Alicia", updatedUser.FirstName)
	assert.Equal(suite.T(), "Johnson-Smith", updatedUser.LastName)
}

func (suite *AuthTestSuite) TestUserRepository_Delete() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test5@example.com", "password123", "Charlie", "Brown")
	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)

	// Delete the user
	err = suite.userRepo.Delete(user.ID)
	assert.NoError(suite.T(), err)

	// Try to retrieve the deleted user (should fail)
	_, err = suite.userRepo.GetByID(user.ID)
	assert.Error(suite.T(), err)
}

func (suite *AuthTestSuite) TestUserRepository_Create_DuplicateEmail() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user1 := model.NewUser("duplicate@example.com", "password123", "John", "Doe")
	err := suite.userRepo.Create(user1)
	assert.NoError(suite.T(), err)

	// Try to create another user with the same email
	user2 := model.NewUser("duplicate@example.com", "password456", "Jane", "Smith")
	err = suite.userRepo.Create(user2)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "already exists")
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
