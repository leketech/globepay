package integration

import (
	"context"
	"testing"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AccountTestSuite struct {
	suite.Suite
	accountRepo repository.AccountRepository
	db          *utils.TestDB
	redisClient *utils.TestRedis
}

func (suite *AccountTestSuite) SetupSuite() {
	// Initialize test database
	suite.db = utils.NewTestDB()
	suite.redisClient = utils.NewTestRedis()

	// Initialize repository
	if suite.db != nil {
		suite.accountRepo = repository.NewAccountRepository(suite.db.DB)
	}
}

func (suite *AccountTestSuite) TearDownSuite() {
	if suite.db != nil {
		suite.db.Close()
	}
	if suite.redisClient != nil {
		suite.redisClient.Close()
	}
}

func (suite *AccountTestSuite) SetupTest() {
	// Clear test data before each test
	if suite.db != nil {
		suite.db.ClearTables()
	}
}

func (suite *AccountTestSuite) TestAccountRepository_Create() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test account
	account := &model.Account{
		UserID:        "1",
		Currency:      "USD",
		Balance:       1000.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	// Test creating account
	err := suite.accountRepo.Create(account)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), account.ID)
}

func (suite *AccountTestSuite) TestAccountRepository_GetByID() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test account
	account := &model.Account{
		UserID:        "1",
		Currency:      "USD",
		Balance:       1000.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	err := suite.accountRepo.Create(account)
	assert.NoError(suite.T(), err)

	// Test getting account by ID
	retrievedAccount, err := suite.accountRepo.GetByID(account.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedAccount)
	assert.Equal(suite.T(), account.ID, retrievedAccount.ID)
	assert.Equal(suite.T(), account.UserID, retrievedAccount.UserID)
	assert.Equal(suite.T(), account.Currency, retrievedAccount.Currency)
	assert.Equal(suite.T(), account.Balance, retrievedAccount.Balance)
	assert.Equal(suite.T(), account.AccountNumber, retrievedAccount.AccountNumber)
	assert.Equal(suite.T(), account.Status, retrievedAccount.Status)
}

func (suite *AccountTestSuite) TestAccountRepository_GetByUser() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test accounts for the same user
	account1 := &model.Account{
		UserID:        "1",
		Currency:      "USD",
		Balance:       1000.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	account2 := &model.Account{
		UserID:        "1",
		Currency:      "EUR",
		Balance:       500.0,
		AccountNumber: "ACC002",
		Status:        "active",
	}

	err := suite.accountRepo.Create(account1)
	assert.NoError(suite.T(), err)

	err = suite.accountRepo.Create(account2)
	assert.NoError(suite.T(), err)

	// Test getting accounts by user ID
	accounts, err := suite.accountRepo.GetByUser(context.Background(), "1")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), accounts, 2)
}

func (suite *AccountTestSuite) TestAccountRepository_UpdateBalance() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test account
	account := &model.Account{
		UserID:        "1",
		Currency:      "USD",
		Balance:       1000.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	err := suite.accountRepo.Create(account)
	assert.NoError(suite.T(), err)

	// Test updating account balance
	newBalance := 1500.0
	err = suite.accountRepo.UpdateBalance(context.Background(), account.ID, newBalance)
	assert.NoError(suite.T(), err)

	// Verify the balance was updated
	updatedAccount, err := suite.accountRepo.GetByID(account.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), newBalance, updatedAccount.Balance)
}

func TestAccountRepository(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}