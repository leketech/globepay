package integration

import (
	"context"
	"testing"

	"globepay/internal/domain"
	"globepay/internal/repository"
	"globepay/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransferTestSuite struct {
	suite.Suite
	transferService *service.TransferService
	transferRepo    repository.TransferRepoInterface
	accountRepo     repository.AccountRepoInterface
	transactionRepo repository.TransactionRepoInterface
	db              *TestDB
	redisClient     *TestRedis
}

func (suite *TransferTestSuite) SetupSuite() {
	// Initialize test database
	suite.db = NewTestDB()
	suite.redisClient = NewTestRedis()

	// Initialize repositories
	suite.accountRepo = repository.NewAccountRepo(suite.db.DB)
	suite.transferRepo = repository.NewTransferRepo(suite.db.DB)
	suite.transactionRepo = repository.NewTransactionRepo(suite.db.DB)

	// Initialize service
	suite.transferService = service.NewTransferService(suite.transferRepo, suite.accountRepo, suite.transactionRepo)
}

func (suite *TransferTestSuite) TearDownSuite() {
	if suite.db != nil {
		suite.db.Close()
	}
	if suite.redisClient != nil {
		suite.redisClient.Close()
	}
}

func (suite *TransferTestSuite) SetupTest() {
	// Clear test data before each test
	suite.db.ClearTables()
}

func (suite *TransferTestSuite) TestTransferService_CreateTransfer() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test accounts
	senderAccount := &domain.Account{
		UserID:        1,
		Currency:      "USD",
		Balance:       1000.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	receiverAccount := &domain.Account{
		UserID:        2,
		Currency:      "USD",
		Balance:       500.0,
		AccountNumber: "ACC002",
		Status:        "active",
	}

	err := suite.accountRepo.Create(senderAccount)
	assert.NoError(suite.T(), err)

	err = suite.accountRepo.Create(receiverAccount)
	assert.NoError(suite.T(), err)

	// Create transfer
	transfer := &domain.Transfer{
		SenderID:          1,
		ReceiverID:        2,
		SenderAccountID:   senderAccount.ID,
		ReceiverAccountID: receiverAccount.ID,
		Amount:            100.0,
		Currency:          "USD",
		Fee:               1.0,
	}

	err = suite.transferService.CreateTransfer(transfer)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), transfer.ReferenceNumber)
	assert.Equal(suite.T(), "processed", transfer.Status)

	// Verify sender account balance
	updatedSenderAccount, err := suite.accountRepo.GetByID(senderAccount.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 899.0, updatedSenderAccount.Balance) // 1000 - 100 - 1

	// Verify receiver account balance
	updatedReceiverAccount, err := suite.accountRepo.GetByID(receiverAccount.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 600.0, updatedReceiverAccount.Balance) // 500 + 100
}

func (suite *TransferTestSuite) TestTransferService_GetTransfers() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test accounts
	senderAccount := &domain.Account{
		UserID:        1,
		Currency:      "USD",
		Balance:       1000.0,
		AccountNumber: "ACC001",
		Status:        "active",
	}

	receiverAccount := &domain.Account{
		UserID:        2,
		Currency:      "USD",
		Balance:       500.0,
		AccountNumber: "ACC002",
		Status:        "active",
	}

	err := suite.accountRepo.Create(senderAccount)
	assert.NoError(suite.T(), err)

	err = suite.accountRepo.Create(receiverAccount)
	assert.NoError(suite.T(), err)

	// Create test transfers
	transfer1 := &domain.Transfer{
		SenderID:          1,
		ReceiverID:        2,
		SenderAccountID:   senderAccount.ID,
		ReceiverAccountID: receiverAccount.ID,
		Amount:            100.0,
		Currency:          "USD",
		Fee:               1.0,
		Status:            "processed",
		ReferenceNumber:   "TRF001",
	}

	transfer2 := &domain.Transfer{
		SenderID:          1,
		ReceiverID:        2,
		SenderAccountID:   senderAccount.ID,
		ReceiverAccountID: receiverAccount.ID,
		Amount:            200.0,
		Currency:          "USD",
		Fee:               2.0,
		Status:            "processed",
		ReferenceNumber:   "TRF002",
	}

	err = suite.transferRepo.Create(transfer1)
	assert.NoError(suite.T(), err)

	err = suite.transferRepo.Create(transfer2)
	assert.NoError(suite.T(), err)

	// Get transfers for user
	transfers, err := suite.transferService.GetTransfers(1)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), transfers, 2)
}

func TestTransferService(t *testing.T) {
	suite.Run(t, new(TransferTestSuite))
}