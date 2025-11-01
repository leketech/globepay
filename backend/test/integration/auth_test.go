package integration

import (
	"testing"

	"globepay/internal/domain"
	"globepay/internal/repository"
	"globepay/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type AuthTestSuite struct {
	suite.Suite
	authService *service.AuthService
	userRepo    repository.UserRepoInterface
	db          *TestDB
	redisClient *TestRedis
}

func (suite *AuthTestSuite) SetupSuite() {
	// Initialize test database
	suite.db = NewTestDB()
	suite.redisClient = NewTestRedis()

	// Initialize repository
	suite.userRepo = repository.NewUserRepo(suite.db.DB)

	// Initialize service
	suite.authService = service.NewAuthService(suite.userRepo, "test-secret-key")
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
	suite.db.ClearTables()
}

func (suite *AuthTestSuite) TestAuthService_Register() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test user
	user := &domain.User{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
		Country:   "US",
		Currency:  "USD",
		Status:    "active",
	}

	// Test registering user
	err := suite.authService.Register(user)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), user.ID)

	// Verify user was created in database
	createdUser, err := suite.userRepo.GetByEmail("test@example.com")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), createdUser)
	assert.Equal(suite.T(), "test@example.com", createdUser.Email)
	assert.Equal(suite.T(), "John", createdUser.FirstName)
	assert.Equal(suite.T(), "Doe", createdUser.LastName)
	// Password should be hashed
	assert.NotEqual(suite.T(), "password123", createdUser.Password)
}

func (suite *AuthTestSuite) TestAuthService_Login() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Create a test user with hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &domain.User{
		Email:     "test@example.com",
		Password:  string(hashedPassword),
		FirstName: "John",
		LastName:  "Doe",
		Country:   "US",
		Currency:  "USD",
		Status:    "active",
	}

	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)

	// Test logging in
	token, err := suite.authService.Login("test@example.com", "password123")
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), token)

	// Verify token is valid JWT
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret-key"), nil
	})
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), parsedToken.Valid)

	// Verify claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), float64(user.ID), claims["user_id"])
	assert.Equal(suite.T(), "test@example.com", claims["email"])
}

func (suite *AuthTestSuite) TestAuthService_Login_InvalidCredentials() {
	// Skip if no database connection
	if suite.db == nil {
		suite.T().Skip("No database connection")
	}

	// Test logging in with invalid credentials
	token, err := suite.authService.Login("nonexistent@example.com", "password123")
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), token)
	assert.Equal(suite.T(), domain.ErrInvalidCredentials, err)

	// Create a user for testing wrong password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &domain.User{
		Email:     "test@example.com",
		Password:  string(hashedPassword),
		FirstName: "John",
		LastName:  "Doe",
	}

	err = suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)

	// Test logging in with wrong password
	token, err = suite.authService.Login("test@example.com", "wrongpassword")
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), token)
	assert.Equal(suite.T(), domain.ErrInvalidCredentials, err)
}

func (suite *AuthTestSuite) TestAuthService_HashPassword() {
	// Test hashing password
	password := "password123"
	hashedPassword, err := suite.authService.HashPassword(password)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), hashedPassword)
	assert.NotEqual(suite.T(), password, hashedPassword)

	// Verify password can be validated
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	assert.NoError(suite.T(), err)
}

func (suite *AuthTestSuite) TestAuthService_GenerateOTP() {
	// Test generating OTP
	otp1, err1 := suite.authService.GenerateOTP()
	otp2, err2 := suite.authService.GenerateOTP()

	assert.NoError(suite.T(), err1)
	assert.NoError(suite.T(), err2)
	assert.NotEmpty(suite.T(), otp1)
	assert.NotEmpty(suite.T(), otp2)
	assert.Len(suite.T(), otp1, 6)
	assert.Len(suite.T(), otp2, 6)
	// OTPs should be different (with very high probability)
	assert.NotEqual(suite.T(), otp1, otp2)
}

func TestAuthService(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}