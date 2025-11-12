package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_NewUser(t *testing.T) {
	user := NewUser("test@example.com", "password123", "John", "Doe")

	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "pending", user.KYCStatus)
	assert.Equal(t, "active", user.AccountStatus)
	assert.False(t, user.EmailVerified)
	assert.False(t, user.PhoneVerified)
	assert.WithinDuration(t, time.Now(), user.CreatedAt, 5*time.Second)
	assert.WithinDuration(t, time.Now(), user.UpdatedAt, 5*time.Second)
}

func TestUser_SetPassword(t *testing.T) {
	user := NewUser("test@example.com", "password123", "John", "Doe")

	err := user.SetPassword("newpassword123")
	assert.NoError(t, err)
	assert.NotEqual(t, "", user.PasswordHash)
}

func TestUser_CheckPassword(t *testing.T) {
	user := NewUser("test@example.com", "password123", "John", "Doe")

	// Set a password first
	err := user.SetPassword("password123")
	assert.NoError(t, err)

	// Test correct password
	match, err := user.CheckPassword("password123")
	assert.NoError(t, err)
	assert.True(t, match)

	// Test incorrect password
	match, err = user.CheckPassword("wrongpassword")
	assert.NoError(t, err)
	assert.False(t, match)
}
