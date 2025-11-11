package integration

import (
	"testing"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DatabaseTestSuite struct {
	suite.Suite
	db       *utils.TestDB
	userRepo repository.UserRepository
	testDB   *utils.TestDB
}

func (suite *DatabaseTestSuite) SetupSuite() {
	// Initialize test database
	suite.testDB = utils.NewTestDB()
	
	// Initialize repository
	if suite.testDB != nil {
		suite.userRepo = repository.NewUserRepository(suite.testDB.DB)
	} else {
		// In a real test, you would connect to a test database
		// For now, we'll skip the actual database connection
		suite.T().Skip("Skipping database tests - no test database configured")
	}
}

func (suite *DatabaseTestSuite) TearDownSuite() {
	if suite.testDB != nil {
		suite.testDB.Close()
	}
}

func (suite *DatabaseTestSuite) SetupTest() {
	// Clear test data before each test
	if suite.testDB != nil {
		suite.testDB.ClearTables()
	}
}

func (suite *DatabaseTestSuite) TestUserRepository_Create() {
	// Skip if no database connection
	if suite.testDB == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test user
	user := model.NewUser("test@example.com", "password123", "John", "Doe")
	
	// Test creating user
	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), user.ID)
}

func (suite *DatabaseTestSuite) TestUserRepository_GetByID() {
	// Skip if no database connection
	if suite.testDB == nil {
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

func (suite *DatabaseTestSuite) TestUserRepository_GetByEmail() {
	// Skip if no database connection
	if suite.testDB == nil {
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

func (suite *DatabaseTestSuite) TestUserRepository_Update() {
	// Skip if no database connection
	if suite.testDB == nil {
		suite.T().Skip("No database connection")
	}

	// First create a user
	user := model.NewUser("test4@example.com", "password123", "Alice", "Johnson")
	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)

	// Update user information
	user.FirstName = "Alicia"
	user.LastName = "Johnson-Smith"
	user.UpdatedAt = time.Now()

	// Save updated user
	err = suite.userRepo.Update(user)
	assert.NoError(suite.T(), err)

	// Retrieve updated user
	updatedUser, err := suite.userRepo.GetByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Alicia", updatedUser.FirstName)
	assert.Equal(suite.T(), "Johnson-Smith", updatedUser.LastName)
}

func (suite *DatabaseTestSuite) TestUserRepository_Delete() {
	// Skip if no database connection
	if suite.testDB == nil {
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

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}