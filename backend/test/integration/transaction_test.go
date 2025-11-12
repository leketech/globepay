package integration

import (
	"context"
	"testing"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
	"globepay/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionTestSuite struct {
	suite.Suite
	transactionRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
	db              *utils.TestDB
	redisClient     *utils.TestRedis
}

func (suite *TransactionTestSuite) SetupSuite() {
	// Initialize test database
	suite.db = utils.NewTestDB()
	suite.redisClient = utils.NewTestRedis()

	// Initialize repositories
	if suite.db != nil {
		suite.transactionRepo = repository.NewTransactionRepository(suite.db.DB)
		suite.accountRepo = repository.NewAccountRepository(suite.db.DB)
	}
}

func (suite *TransactionTestSuite) TearDownSuite() {
	if suite.db != nil {
		suite.db.Close()
	}
	if suite.redisClient != nil {
		suite.redisClient.Close()
	}
}

func (suite *TransactionTestSuite) SetupTest() {
	// Clear test data before each test
	if suite.db != nil {
		suite.db.ClearTables()
	}
}

func (suite *TransactionTestSuite) TestTransactionRepository_Create() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transaction
	transaction := &model.Transaction{
		UserID:          "1",
		AccountID:       "1",
		Type:            string(model.TransactionDeposit),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             0.0,
		Description:     "Test deposit",
		ReferenceNumber: "DEP001",
		Status:          "pending",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Test creating transaction
	err := suite.transactionRepo.Create(transaction)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), transaction.ID)
}

func (suite *TransactionTestSuite) TestTransactionRepository_GetByID() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transaction
	transaction := &model.Transaction{
		UserID:          "1",
		AccountID:       "1",
		Type:            string(model.TransactionDeposit),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             0.0,
		Description:     "Test deposit",
		ReferenceNumber: "DEP001",
		Status:          "pending",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := suite.transactionRepo.Create(transaction)
	assert.NoError(suite.T(), err)

	// Test getting transaction by ID
	retrievedTransaction, err := suite.transactionRepo.GetByID(transaction.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedTransaction)
	assert.Equal(suite.T(), transaction.ID, retrievedTransaction.ID)
	assert.Equal(suite.T(), transaction.UserID, retrievedTransaction.UserID)
	assert.Equal(suite.T(), transaction.AccountID, retrievedTransaction.AccountID)
	assert.Equal(suite.T(), transaction.Type, retrievedTransaction.Type)
	assert.Equal(suite.T(), transaction.Amount, retrievedTransaction.Amount)
	assert.Equal(suite.T(), transaction.Currency, retrievedTransaction.Currency)
	assert.Equal(suite.T(), transaction.Description, retrievedTransaction.Description)
	assert.Equal(suite.T(), transaction.ReferenceNumber, retrievedTransaction.ReferenceNumber)
	assert.Equal(suite.T(), transaction.Status, retrievedTransaction.Status)
}

func (suite *TransactionTestSuite) TestTransactionRepository_GetByUser() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test transactions for the same user
	transaction1 := &model.Transaction{
		UserID:          "1",
		AccountID:       "1",
		Type:            string(model.TransactionDeposit),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             0.0,
		Description:     "Test deposit 1",
		ReferenceNumber: "DEP001",
		Status:          "completed",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	transaction2 := &model.Transaction{
		UserID:          "1",
		AccountID:       "1",
		Type:            string(model.TransactionWithdrawal),
		Amount:          50.0,
		Currency:        "USD",
		Fee:             1.0,
		Description:     "Test withdrawal",
		ReferenceNumber: "WIT001",
		Status:          "completed",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := suite.transactionRepo.Create(transaction1)
	assert.NoError(suite.T(), err)

	err = suite.transactionRepo.Create(transaction2)
	assert.NoError(suite.T(), err)

	// Test getting transactions by user ID
	transactions, err := suite.transactionRepo.GetByUser(context.TODO(), "1", 100, 0)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), transactions, 2)
}

func (suite *TransactionTestSuite) TestTransactionRepository_Update() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transaction
	transaction := &model.Transaction{
		UserID:          "1",
		AccountID:       "1",
		Type:            string(model.TransactionDeposit),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             0.0,
		Description:     "Test deposit",
		ReferenceNumber: "DEP001",
		Status:          "pending",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := suite.transactionRepo.Create(transaction)
	assert.NoError(suite.T(), err)

	// Update transaction status
	transaction.Status = "completed"
	transaction.UpdatedAt = time.Now()

	// Save updated transaction
	err = suite.transactionRepo.Update(transaction)
	assert.NoError(suite.T(), err)

	// Retrieve updated transaction
	updatedTransaction, err := suite.transactionRepo.GetByID(transaction.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "completed", updatedTransaction.Status)
}

func (suite *TransactionTestSuite) TestTransactionRepository_Delete() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transaction
	transaction := &model.Transaction{
		UserID:          "1",
		AccountID:       "1",
		Type:            string(model.TransactionDeposit),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             0.0,
		Description:     "Test deposit",
		ReferenceNumber: "DEP001",
		Status:          "pending",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := suite.transactionRepo.Create(transaction)
	assert.NoError(suite.T(), err)

	// Delete the transaction
	err = suite.transactionRepo.Delete(transaction.ID)
	assert.NoError(suite.T(), err)

	// Try to retrieve the deleted transaction (should fail)
	_, err = suite.transactionRepo.GetByID(transaction.ID)
	assert.Error(suite.T(), err)
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}
