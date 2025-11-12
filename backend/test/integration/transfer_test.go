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

type TransferTestSuite struct {
	suite.Suite
	transferRepo    repository.TransferRepository
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
	db              *utils.TestDB
	redisClient     *utils.TestRedis
}

func (suite *TransferTestSuite) SetupSuite() {
	// Initialize test database
	suite.db = utils.NewTestDB()
	suite.redisClient = utils.NewTestRedis()

	// Initialize repositories
	if suite.db != nil {
		suite.transferRepo = repository.NewTransferRepository(suite.db.DB)
		suite.accountRepo = repository.NewAccountRepository(suite.db.DB)
		suite.transactionRepo = repository.NewTransactionRepository(suite.db.DB)
	}
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
	if suite.db != nil {
		suite.db.ClearTables()
	}
}

func (suite *TransferTestSuite) TestTransferRepository_Create() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test transfer
	transfer := &model.Transfer{
		UserID:                 "1",
		RecipientName:          "John Doe",
		RecipientEmail:         "john@example.com",
		RecipientCountry:       "US",
		RecipientBankName:      "Test Bank",
		RecipientAccountNumber: "123456789",
		RecipientSwiftCode:     "TESTUS01",
		SourceCurrency:         "USD",
		DestCurrency:           "USD",
		SourceAmount:           100.0,
		DestAmount:             100.0,
		ExchangeRate:           1.0,
		FeeAmount:              1.0,
		Purpose:                "Test transfer",
		Status:                 "pending",
		ReferenceNumber:        "TRF001",
		EstimatedArrival:       time.Now().Add(24 * time.Hour),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
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
		UserID:                 "1",
		RecipientName:          "John Doe",
		RecipientEmail:         "john@example.com",
		RecipientCountry:       "US",
		RecipientBankName:      "Test Bank",
		RecipientAccountNumber: "123456789",
		RecipientSwiftCode:     "TESTUS01",
		SourceCurrency:         "USD",
		DestCurrency:           "USD",
		SourceAmount:           100.0,
		DestAmount:             100.0,
		ExchangeRate:           1.0,
		FeeAmount:              1.0,
		Purpose:                "Test transfer",
		Status:                 "pending",
		ReferenceNumber:        "TRF001",
		EstimatedArrival:       time.Now().Add(24 * time.Hour),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	err := suite.transferRepo.Create(transfer)
	assert.NoError(suite.T(), err)

	// Test getting transfer by ID
	retrievedTransfer, err := suite.transferRepo.GetByID(transfer.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedTransfer)
	assert.Equal(suite.T(), transfer.ID, retrievedTransfer.ID)
	assert.Equal(suite.T(), transfer.UserID, retrievedTransfer.UserID)
	assert.Equal(suite.T(), transfer.RecipientName, retrievedTransfer.RecipientName)
	assert.Equal(suite.T(), transfer.RecipientEmail, retrievedTransfer.RecipientEmail)
	assert.Equal(suite.T(), transfer.RecipientCountry, retrievedTransfer.RecipientCountry)
	assert.Equal(suite.T(), transfer.RecipientBankName, retrievedTransfer.RecipientBankName)
	assert.Equal(suite.T(), transfer.RecipientAccountNumber, retrievedTransfer.RecipientAccountNumber)
	assert.Equal(suite.T(), transfer.RecipientSwiftCode, retrievedTransfer.RecipientSwiftCode)
	assert.Equal(suite.T(), transfer.SourceCurrency, retrievedTransfer.SourceCurrency)
	assert.Equal(suite.T(), transfer.DestCurrency, retrievedTransfer.DestCurrency)
	assert.Equal(suite.T(), transfer.SourceAmount, retrievedTransfer.SourceAmount)
	assert.Equal(suite.T(), transfer.DestAmount, retrievedTransfer.DestAmount)
	assert.Equal(suite.T(), transfer.ExchangeRate, retrievedTransfer.ExchangeRate)
	assert.Equal(suite.T(), transfer.FeeAmount, retrievedTransfer.FeeAmount)
	assert.Equal(suite.T(), transfer.Purpose, retrievedTransfer.Purpose)
	assert.Equal(suite.T(), transfer.Status, retrievedTransfer.Status)
	assert.Equal(suite.T(), transfer.ReferenceNumber, retrievedTransfer.ReferenceNumber)
}

func (suite *TransferTestSuite) TestTransferRepository_GetByUser() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create test transfers for the same user
	transfer1 := &model.Transfer{
		UserID:                 "1",
		RecipientName:          "John Doe",
		RecipientEmail:         "john@example.com",
		RecipientCountry:       "US",
		RecipientBankName:      "Test Bank",
		RecipientAccountNumber: "123456789",
		RecipientSwiftCode:     "TESTUS01",
		SourceCurrency:         "USD",
		DestCurrency:           "USD",
		SourceAmount:           100.0,
		DestAmount:             100.0,
		ExchangeRate:           1.0,
		FeeAmount:              1.0,
		Purpose:                "Test transfer 1",
		Status:                 "completed",
		ReferenceNumber:        "TRF001",
		EstimatedArrival:       time.Now().Add(24 * time.Hour),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	transfer2 := &model.Transfer{
		UserID:                 "1",
		RecipientName:          "Jane Smith",
		RecipientEmail:         "jane@example.com",
		RecipientCountry:       "UK",
		RecipientBankName:      "Another Bank",
		RecipientAccountNumber: "987654321",
		RecipientSwiftCode:     "ANOTUK01",
		SourceCurrency:         "USD",
		DestCurrency:           "GBP",
		SourceAmount:           50.0,
		DestAmount:             40.0,
		ExchangeRate:           0.8,
		FeeAmount:              1.0,
		Purpose:                "Test transfer 2",
		Status:                 "completed",
		ReferenceNumber:        "TRF002",
		EstimatedArrival:       time.Now().Add(24 * time.Hour),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	err := suite.transferRepo.Create(transfer1)
	assert.NoError(suite.T(), err)

	err = suite.transferRepo.Create(transfer2)
	assert.NoError(suite.T(), err)

	// Test getting transfers by user ID
	transfers, err := suite.transferRepo.GetByUser(context.TODO(), "1", 100, 0)
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
		UserID:                 "1",
		RecipientName:          "John Doe",
		RecipientEmail:         "john@example.com",
		RecipientCountry:       "US",
		RecipientBankName:      "Test Bank",
		RecipientAccountNumber: "123456789",
		RecipientSwiftCode:     "TESTUS01",
		SourceCurrency:         "USD",
		DestCurrency:           "USD",
		SourceAmount:           100.0,
		DestAmount:             100.0,
		ExchangeRate:           1.0,
		FeeAmount:              1.0,
		Purpose:                "Test transfer",
		Status:                 "pending",
		ReferenceNumber:        "TRF001",
		EstimatedArrival:       time.Now().Add(24 * time.Hour),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
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
		UserID:                 "1",
		RecipientName:          "John Doe",
		RecipientEmail:         "john@example.com",
		RecipientCountry:       "US",
		RecipientBankName:      "Test Bank",
		RecipientAccountNumber: "123456789",
		RecipientSwiftCode:     "TESTUS01",
		SourceCurrency:         "USD",
		DestCurrency:           "USD",
		SourceAmount:           100.0,
		DestAmount:             100.0,
		ExchangeRate:           1.0,
		FeeAmount:              1.0,
		Purpose:                "Test transfer",
		Status:                 "pending",
		ReferenceNumber:        "TRF001",
		EstimatedArrival:       time.Now().Add(24 * time.Hour),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
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
