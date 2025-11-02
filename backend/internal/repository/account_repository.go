package repository

import (
	"context"
	"database/sql"
	"time"

	"globepay/internal/domain/model"
)

// AccountRepo implements AccountRepository
type AccountRepo struct {
	db *sql.DB
}

// NewAccountRepository creates a new AccountRepo
func NewAccountRepository(db *sql.DB) AccountRepository {
	return &AccountRepo{db: db}
}

// Create inserts a new account into the database
func (r *AccountRepo) Create(account *model.Account) error {
	query := `
		INSERT INTO accounts (user_id, account_number, account_type, currency, balance, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	now := time.Now()
	account.CreatedAt = now
	account.UpdatedAt = now

	return r.db.QueryRowContext(context.Background(), query, account.UserID, account.AccountNumber, account.AccountType, account.Currency, account.Balance, account.Status, account.CreatedAt, account.UpdatedAt).Scan(&account.ID)
}

// GetByID retrieves an account by ID
func (r *AccountRepo) GetByID(id string) (*model.Account, error) {
	query := `
		SELECT id, user_id, account_number, account_type, currency, balance, status, created_at, updated_at
		FROM accounts
		WHERE id = $1
	`

	account := &model.Account{}
	err := r.db.QueryRowContext(context.Background(), query, id).Scan(
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

// GetByUser retrieves all accounts for a user
func (r *AccountRepo) GetByUser(ctx context.Context, userID string) ([]*model.Account, error) {
	query := `
		SELECT id, user_id, account_number, account_type, currency, balance, status, created_at, updated_at
		FROM accounts
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []*model.Account{}
	for rows.Next() {
		account := &model.Account{}
		err := rows.Scan(
			&account.ID, &account.UserID, &account.AccountNumber, &account.AccountType,
			&account.Currency, &account.Balance, &account.Status, &account.CreatedAt, &account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

// GetByNumber retrieves an account by account number
func (r *AccountRepo) GetByNumber(ctx context.Context, accountNumber string) (*model.Account, error) {
	query := `
		SELECT id, user_id, account_number, account_type, currency, balance, status, created_at, updated_at
		FROM accounts
		WHERE account_number = $1
	`

	account := &model.Account{}
	err := r.db.QueryRowContext(ctx, query, accountNumber).Scan(
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

// Update updates an account in the database
func (r *AccountRepo) Update(account *model.Account) error {
	query := `
		UPDATE accounts
		SET user_id = $1, account_number = $2, account_type = $3, currency = $4, balance = $5, status = $6, updated_at = $7
		WHERE id = $8
	`

	account.UpdatedAt = time.Now()
	result, err := r.db.ExecContext(context.Background(), query, account.UserID, account.AccountNumber, account.AccountType, account.Currency, account.Balance, account.Status, account.UpdatedAt, account.ID)
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

// Delete removes an account from the database
func (r *AccountRepo) Delete(id string) error {
	query := `DELETE FROM accounts WHERE id = $1`

	result, err := r.db.ExecContext(context.Background(), query, id)
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

// GetAll retrieves all accounts from the database
func (r *AccountRepo) GetAll() ([]model.Account, error) {
	query := `
		SELECT id, user_id, account_number, account_type, currency, balance, status, created_at, updated_at
		FROM accounts
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []model.Account{}
	for rows.Next() {
		var account model.Account
		err := rows.Scan(
			&account.ID, &account.UserID, &account.AccountNumber, &account.AccountType,
			&account.Currency, &account.Balance, &account.Status, &account.CreatedAt, &account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

// GetByStatus retrieves all accounts with a specific status
func (r *AccountRepo) GetByStatus(status string) ([]model.Account, error) {
	query := `
		SELECT id, user_id, account_number, account_type, currency, balance, status, created_at, updated_at
		FROM accounts
		WHERE status = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(context.Background(), query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []model.Account{}
	for rows.Next() {
		var account model.Account
		err := rows.Scan(
			&account.ID, &account.UserID, &account.AccountNumber, &account.AccountType,
			&account.Currency, &account.Balance, &account.Status, &account.CreatedAt, &account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

// UpdateBalance updates the balance of an account
func (r *AccountRepo) UpdateBalance(ctx context.Context, accountID string, newBalance float64) error {
	query := `
		UPDATE accounts
		SET balance = $1, updated_at = $2
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, newBalance, time.Now(), accountID)
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

// GetByUserAndCurrency retrieves an account by user ID and currency
func (r *AccountRepo) GetByUserAndCurrency(ctx context.Context, userID, currency string) (*model.Account, error) {
	query := `
		SELECT id, user_id, account_number, account_type, currency, balance, status, created_at, updated_at
		FROM accounts
		WHERE user_id = $1 AND currency = $2
	`

	account := &model.Account{}
	err := r.db.QueryRowContext(ctx, query, userID, currency).Scan(
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