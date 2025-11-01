package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"globepay/internal/domain"
)

// LoadUsersFixture loads user data from the users.json fixture file
func LoadUsersFixture() ([]domain.User, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Construct the path to the fixtures file
	fixturePath := filepath.Join(wd, "fixtures", "users.json")

	// Read the file
	data, err := os.ReadFile(fixturePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data
	var users []domain.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// LoadAccountsFixture loads account data from the accounts.json fixture file
func LoadAccountsFixture() ([]domain.Account, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Construct the path to the fixtures file
	fixturePath := filepath.Join(wd, "fixtures", "accounts.json")

	// Read the file
	data, err := os.ReadFile(fixturePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data
	var accounts []domain.Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// LoadTransfersFixture loads transfer data from the transfers.json fixture file
func LoadTransfersFixture() ([]domain.Transfer, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Construct the path to the fixtures file
	fixturePath := filepath.Join(wd, "fixtures", "transfers.json")

	// Read the file
	data, err := os.ReadFile(fixturePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data
	var transfers []domain.Transfer
	err = json.Unmarshal(data, &transfers)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}

// GetTestUser returns a specific user from the fixture data by index
func GetTestUser(index int) (*domain.User, error) {
	users, err := LoadUsersFixture()
	if err != nil {
		return nil, err
	}

	if index >= len(users) {
		return nil, nil
	}

	return &users[index], nil
}

// GetTestAccount returns a specific account from the fixture data by index
func GetTestAccount(index int) (*domain.Account, error) {
	accounts, err := LoadAccountsFixture()
	if err != nil {
		return nil, err
	}

	if index >= len(accounts) {
		return nil, nil
	}

	return &accounts[index], nil
}