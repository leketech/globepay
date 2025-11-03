package integration

import (
	"testing"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransferTestSuite struct {
	suite.Suite
	transferRepo    repository.TransferRepository    // Changed from TransferRepoInterface
	accountRepo     repository.AccountRepository     // Changed from AccountRepoInterface
	transactionRepo repository.TransactionRepository // Changed from TransactionRepoInterface
	db              *TestDB
	redisClient     *TestRedis
}

func (suite *TransferTestSuite) SetupSuite() {
	// Initialize test database
	suite.db = NewTestDB()
	suite.redisClient = NewTestRedis()

	// Initialize repositories
	suite.transferRepo = repository.NewTransferRepository(suite.db.DB)       // Changed from NewTransferRepo
	suite.accountRepo = repository.NewAccountRepository(suite.db.DB)         // Changed from NewAccountRepo
	suite.transactionRepo = repository.NewTransactionRepository(suite.db.DB) // Changed from NewTransactionRepo
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

func (suite *TransferTestSuite) TestTransferRepository_Create() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transfer
	transfer := &model.Transfer{
		UserID:             "1",
		SourceAccountID:    "1",
		DestinationAccountID: "2",
		Amount:             100.0,
		Currency:           "USD",
		Fee:                1.0,
		ExchangeRate:       1.0,
		Description:        "Test transfer",
		ReferenceNumber:    "TRF001",
		Status:             "pending",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	// Test creating transfer
	err := suite.transferRepo.Create(transfer)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), transfer.ID)
}

func (suite *TransferTestSuite) TestTransferRepository_GetByID() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transfer
	transfer := &model.Transfer{
		UserID:             "1",
		SourceAccountID:    "1",
		DestinationAccountID: "2",
		Amount:             100.0,
		Currency:           "USD",
		Fee:                1.0,
		ExchangeRate:       1.0,
		Description:        "Test transfer",
		ReferenceNumber:    "TRF001",
		Status:             "pending",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	err := suite.transferRepo.Create(transfer)
	assert.NoError(suite.T(), err)

	// Test getting transfer by ID
	retrievedTransfer, err := suite.transferRepo.GetByID(transfer.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedTransfer)
	assert.Equal(suite.T(), transfer.ID, retrievedTransfer.ID)
	assert.Equal(suite.T(), transfer.UserID, retrievedTransfer.UserID)
	assert.Equal(suite.T(), transfer.SourceAccountID, retrievedTransfer.SourceAccountID)
	assert.Equal(suite.T(), transfer.DestinationAccountID, retrievedTransfer.DestinationAccountID)
	assert.Equal(suite.T(), transfer.Amount, retrievedTransfer.Amount)
	assert.Equal(suite.T(), transfer.Currency, retrievedTransfer.Currency)
	assert.Equal(suite.T(), transfer.Description, retrievedTransfer.Description)
	assert.Equal(suite.T(), transfer.ReferenceNumber, retrievedTransfer.ReferenceNumber)
	assert.Equal(suite.T(), transfer.Status, retrievedTransfer.Status)
}

func (suite *TransferTestSuite) TestTransferRepository_GetByUser() { // Changed from GetByUserID
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test transfers for the same user
	transfer1 := &model.Transfer{
		UserID:             "1",
		SourceAccountID:    "1",
		DestinationAccountID: "2",
		Amount:             100.0,
		Currency:           "USD",
		Fee:                1.0,
		ExchangeRate:       1.0,
		Description:        "Test transfer 1",
		ReferenceNumber:    "TRF001",
		Status:             "completed",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	transfer2 := &model.Transfer{
		UserID:             "1",
		SourceAccountID:    "1",
		DestinationAccountID: "3",
		Amount:             50.0,
		Currency:           "USD",
		Fee:                1.0,
		ExchangeRate:       1.0,
		Description:        "Test transfer 2",
		ReferenceNumber:    "TRF002",
		Status:             "completed",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	err := suite.transferRepo.Create(transfer1)
	assert.NoError(suite.T(), err)

	err = suite.transferRepo.Create(transfer2)
	assert.NoError(suite.T(), err)

	// Test getting transfers by user ID
	transfers, err := suite.transferRepo.GetByUser(nil, "1", 100, 0) // Added context and pagination
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), transfers, 2)
}

func (suite *TransferTestSuite) TestTransferRepository_Update() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transfer
	transfer := &model.Transfer{
		UserID:             "1",
		SourceAccountID:    "1",
		DestinationAccountID: "2",
		Amount:             100.0,
		Currency:           "USD",
		Fee:                1.0,
		ExchangeRate:       1.0,
		Description:        "Test transfer",
		ReferenceNumber:    "TRF001",
		Status:             "pending",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	err := suite.transferRepo.Create(transfer)
	assert.NoError(suite.T(), err)

	// Update transfer status
	transfer.Status = "completed"
	transfer.UpdatedAt = time.Now()

	// Save updated transfer
	err = suite.transferRepo.Update(transfer)
	assert.NoError(suite.T(), err)

	// Retrieve updated transfer
	updatedTransfer, err := suite.transferRepo.GetByID(transfer.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "completed", updatedTransfer.Status)
}

func (suite *TransferTestSuite) TestTransferRepository_Delete() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transfer
	transfer := &model.Transfer{
		UserID:             "1",
		SourceAccountID:    "1",
		DestinationAccountID: "2",
		Amount:             100.0,
		Currency:           "USD",
		Fee:                1.0,
		ExchangeRate:       1.0,
		Description:        "Test transfer",
		ReferenceNumber:    "TRF001",
		Status:             "pending",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	err := suite.transferRepo.Create(transfer)
	assert.NoError(suite.T(), err)

	// Delete the transfer
	err = suite.transferRepo.Delete(transfer.ID)
	assert.NoError(suite.T(), err)

	// Try to retrieve the deleted transfer (should fail)
	_, err = suite.transferRepo.GetByID(transfer.ID)
	assert.Error(suite.T(), err)
}

func TestTransferTestSuite(t *testing.T) {
	suite.Run(t, new(TransferTestSuite))
}