package repository

import (
	"database/sql"
	"fmt"
	"time"

	"globepay/internal/domain/model"
)

// UserRepo implements UserRepository
type UserRepo struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepo
func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepo{db: db}
}

// Create inserts a new user into the database
func (r *UserRepo) Create(user *model.User) error {
	// First check if user with this email already exists
	exists, err := r.emailExists(user.Email)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	if exists {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	query := `
		INSERT INTO users (id, email, password_hash, first_name, last_name, phone_number, date_of_birth, country_code, kyc_status, account_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Handle nullable fields
	var dateOfBirth *time.Time
	if !user.DateOfBirth.IsZero() {
		dateOfBirth = &user.DateOfBirth
	}

	var countryCode *string
	if user.Country != "" {
		countryCode = &user.Country
	}

	return r.db.QueryRow(query, user.ID, user.Email, user.PasswordHash, user.FirstName, user.LastName, user.PhoneNumber, dateOfBirth, countryCode, user.KYCStatus, user.AccountStatus, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
}

// emailExists checks if a user with the given email already exists
func (r *UserRepo) emailExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// GetByID retrieves a user by ID
func (r *UserRepo) GetByID(id string) (*model.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, phone_number, date_of_birth, country_code, kyc_status, account_status, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &model.User{}
	var dateOfBirth *time.Time
	var countryCode *string
	
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName,
		&user.PhoneNumber, &dateOfBirth, &countryCode, &user.KYCStatus, &user.AccountStatus, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	// Handle nullable fields
	if dateOfBirth != nil {
		user.DateOfBirth = *dateOfBirth
	}
	
	if countryCode != nil {
		user.Country = *countryCode
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepo) GetByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, phone_number, date_of_birth, country_code, kyc_status, account_status, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &model.User{}
	var dateOfBirth *time.Time
	var countryCode *string
	
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName,
		&user.PhoneNumber, &dateOfBirth, &countryCode, &user.KYCStatus, &user.AccountStatus, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	// Handle nullable fields
	if dateOfBirth != nil {
		user.DateOfBirth = *dateOfBirth
	}
	
	if countryCode != nil {
		user.Country = *countryCode
	}

	return user, nil
}

// Update updates a user in the database
func (r *UserRepo) Update(user *model.User) error {
	query := `
		UPDATE users
		SET email = $1, password_hash = $2, first_name = $3, last_name = $4, phone_number = $5, date_of_birth = $6, country_code = $7, kyc_status = $8, account_status = $9, updated_at = $10
		WHERE id = $11
	`

	user.UpdatedAt = time.Now()
	
	// Handle nullable fields
	var dateOfBirth *time.Time
	if !user.DateOfBirth.IsZero() {
		dateOfBirth = &user.DateOfBirth
	}

	var countryCode *string
	if user.Country != "" {
		countryCode = &user.Country
	}

	result, err := r.db.Exec(query, user.Email, user.PasswordHash, user.FirstName, user.LastName, user.PhoneNumber, dateOfBirth, countryCode, user.KYCStatus, user.AccountStatus, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Delete removes a user from the database
func (r *UserRepo) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// GetAll retrieves all users from the database
func (r *UserRepo) GetAll() ([]model.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, phone_number, date_of_birth, country_code, kyc_status, account_status, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		var dateOfBirth *time.Time
		var countryCode *string
		
		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName,
			&user.PhoneNumber, &dateOfBirth, &countryCode, &user.KYCStatus, &user.AccountStatus, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Handle nullable fields
		if dateOfBirth != nil {
			user.DateOfBirth = *dateOfBirth
		}
		
		if countryCode != nil {
			user.Country = *countryCode
		}
		
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetByUserAndCurrency retrieves an account by user ID and currency
func (r *UserRepo) GetByUserAndCurrency(ctx interface{}, userID, currency string) (*model.Account, error) {
	query := `
		SELECT id, user_id, account_number, account_type, currency, balance, status, created_at, updated_at
		FROM accounts
		WHERE user_id = $1 AND currency = $2
	`

	account := &model.Account{}
	err := r.db.QueryRow(query, userID, currency).Scan(
		&account.ID, &account.UserID, &account.AccountNumber, &account.AccountType,
		&account.Currency, &account.Balance, &account.Status, &account.CreatedAt, &account.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return account, nil
}