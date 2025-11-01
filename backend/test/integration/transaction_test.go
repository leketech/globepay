package integration

import (
	"testing"

	"globepay/internal/domain"
	"globepay/internal/repository"
	"globepay/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionTestSuite struct {
	suite.Suite
	transactionService *service.TransactionService
	transactionRepo    repository.TransactionRepoInterface
	accountRepo        repository.AccountRepoInterface
	db                 *TestDB
	redisClient        *TestRedis
}

func (suite *TransactionTestSuite) SetupSuite() {
	// Initialize test database
	suite.db = NewTestDB()
	suite.redisClient = NewTestRedis()

	// Initialize repositories
	suite.accountRepo = repository.NewAccountRepo(suite.db.DB)
	suite.transactionRepo = repository.NewTransactionRepo(suite.db.DB)

	// Initialize service
	suite.transactionService = service.NewTransactionService(suite.transactionRepo, suite.accountRepo)
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
	suite.db.ClearTables()
}

func (suite *TransactionTestSuite) TestTransactionService_CreateDepositTransaction() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test account
	account := &domain.Account{
		UserID:        1,
		Currency:      "USD",
		Balance:       500.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	err := suite.accountRepo.Create(account)
	assert.NoError(suite.T(), err)

	// Create deposit transaction
	transaction := &domain.Transaction{
		UserID:          1,
		AccountID:       account.ID,
		Type:            string(domain.TransactionDeposit),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             0.0,
		Description:     "Test deposit",
		ReferenceNumber: "DEP001",
	}

	err = suite.transactionService.CreateTransaction(transaction)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "processed", transaction.Status)

	// Verify account balance
	updatedAccount, err := suite.accountRepo.GetByID(account.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 600.0, updatedAccount.Balance) // 500 + 100
}

func (suite *TransactionTestSuite) TestTransactionService_CreateWithdrawalTransaction() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test account with sufficient funds
	account := &domain.Account{
		UserID:        1,
		Currency:      "USD",
		Balance:       500.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	err := suite.accountRepo.Create(account)
	assert.NoError(suite.T(), err)

	// Create withdrawal transaction
	transaction := &domain.Transaction{
		UserID:          1,
		AccountID:       account.ID,
		Type:            string(domain.TransactionWithdrawal),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             5.0,
		Description:     "Test withdrawal",
		ReferenceNumber: "WTH001",
	}

	err = suite.transactionService.CreateTransaction(transaction)
	assert.NoError(suite.T(), err)

	// Manually process the transaction (in real app, this would be done by a worker)
	err = suite.transactionService.UpdateTransactionStatus(transaction.ID, domain.TransactionProcessed)
	assert.NoError(suite.T(), err)

	// Verify account balance
	updatedAccount, err := suite.accountRepo.GetByID(account.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 395.0, updatedAccount.Balance) // 500 - 100 - 5
}

func (suite *TransactionTestSuite) TestTransactionService_GetTransactions() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test account
	account := &domain.Account{
		UserID:        1,
		Currency:      "USD",
		Balance:       500.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	err := suite.accountRepo.Create(account)
	assert.NoError(suite.T(), err)

	// Create test transactions
	deposit := &domain.Transaction{
		UserID:          1,
		AccountID:       account.ID,
		Type:            string(domain.TransactionDeposit),
		Status:          string(domain.TransactionProcessed),
		Amount:          100.0,
		Currency:        "USD",
		Fee:             0.0,
		Description:     "Test deposit",
		ReferenceNumber: "DEP001",
	}

	withdrawal := &domain.Transaction{
		UserID:          1,
		AccountID:       account.ID,
		Type:            string(domain.TransactionWithdrawal),
		Status:          string(domain.TransactionProcessed),
		Amount:          50.0,
		Currency:        "USD",
		Fee:             2.0,
		Description:     "Test withdrawal",
		ReferenceNumber: "WTH001",
	}

	err = suite.transactionRepo.Create(deposit)
	assert.NoError(suite.T(), err)

	err = suite.transactionRepo.Create(withdrawal)
	assert.NoError(suite.T(), err)

	// Get transactions for user
	transactions, err := suite.transactionService.GetTransactions(1)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), transactions, 2)
}

func TestTransactionService(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}