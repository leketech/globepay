package repository

import (
	"context"
	"database/sql"
	"time"

	"globepay/internal/domain/model"
)

// TransactionRepo implements TransactionRepository
type TransactionRepo struct {
	db *sql.DB
}

// NewTransactionRepository creates a new TransactionRepo
func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &TransactionRepo{db: db}
}

// Create inserts a new transaction into the database
func (r *TransactionRepo) Create(transaction *model.Transaction) error {
	query := `
		INSERT INTO transactions (id, user_id, type, status, amount, currency, source_account_id, dest_account_id, fee, exchange_rate, description, reference, processed_at, failure_reason, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		RETURNING id
	`

	now := time.Now()
	transaction.CreatedAt = now
	transaction.UpdatedAt = now

	return r.db.QueryRowContext(context.Background(), query, transaction.ID, transaction.UserID, transaction.Type, transaction.Status, transaction.Amount, transaction.Currency, transaction.SourceAccountID, transaction.DestAccountID, transaction.Fee, transaction.ExchangeRate, transaction.Description, transaction.Reference, transaction.ProcessedAt, transaction.FailureReason, transaction.CreatedAt, transaction.UpdatedAt).Scan(&transaction.ID)
}

// GetByID retrieves a transaction by ID
func (r *TransactionRepo) GetByID(id string) (*model.Transaction, error) {
	query := `
		SELECT id, user_id, type, status, amount, currency, source_account_id, dest_account_id, fee, exchange_rate, description, reference, processed_at, failure_reason, created_at, updated_at
		FROM transactions
		WHERE id = $1
	`

	transaction := &model.Transaction{}
	err := r.db.QueryRowContext(context.Background(), query, id).Scan(
		&transaction.ID, &transaction.UserID, &transaction.Type, &transaction.Status,
		&transaction.Amount, &transaction.Currency, &transaction.SourceAccountID, &transaction.DestAccountID,
		&transaction.Fee, &transaction.ExchangeRate, &transaction.Description, &transaction.Reference,
		&transaction.ProcessedAt, &transaction.FailureReason,
		&transaction.CreatedAt, &transaction.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return transaction, nil
}

// GetByUser retrieves all transactions for a user
func (r *TransactionRepo) GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Transaction, error) {
	query := `
		SELECT id, user_id, type, status, amount, currency, source_account_id, dest_account_id, fee, exchange_rate, description, reference, processed_at, failure_reason, created_at, updated_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []*model.Transaction{}
	for rows.Next() {
		transaction := &model.Transaction{}
		err := rows.Scan(
			&transaction.ID, &transaction.UserID, &transaction.Type, &transaction.Status,
			&transaction.Amount, &transaction.Currency, &transaction.SourceAccountID, &transaction.DestAccountID,
			&transaction.Fee, &transaction.ExchangeRate, &transaction.Description, &transaction.Reference,
			&transaction.ProcessedAt, &transaction.FailureReason,
			&transaction.CreatedAt, &transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

// Update updates a transaction in the database
func (r *TransactionRepo) Update(transaction *model.Transaction) error {
	query := `
		UPDATE transactions
		SET user_id = $1, type = $2, status = $3, amount = $4, currency = $5, source_account_id = $6, dest_account_id = $7, fee = $8, exchange_rate = $9, description = $10, reference = $11, processed_at = $12, failure_reason = $13, updated_at = $14
		WHERE id = $15
	`

	transaction.UpdatedAt = time.Now()
	result, err := r.db.ExecContext(context.Background(), query, transaction.UserID, transaction.Type, transaction.Status, transaction.Amount, transaction.Currency, transaction.SourceAccountID, transaction.DestAccountID, transaction.Fee, transaction.ExchangeRate, transaction.Description, transaction.Reference, transaction.ProcessedAt, transaction.FailureReason, transaction.UpdatedAt, transaction.ID)
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

// Delete removes a transaction from the database
func (r *TransactionRepo) Delete(id string) error {
	query := `DELETE FROM transactions WHERE id = $1`

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

// GetAll retrieves all transactions from the database
func (r *TransactionRepo) GetAll() ([]model.Transaction, error) {
	query := `
		SELECT id, user_id, type, status, amount, currency, source_account_id, dest_account_id, fee, exchange_rate, description, reference, processed_at, failure_reason, created_at, updated_at
		FROM transactions
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(
			&transaction.ID, &transaction.UserID, &transaction.Type, &transaction.Status,
			&transaction.Amount, &transaction.Currency, &transaction.SourceAccountID, &transaction.DestAccountID,
			&transaction.Fee, &transaction.ExchangeRate, &transaction.Description, &transaction.Reference,
			&transaction.ProcessedAt, &transaction.FailureReason,
			&transaction.CreatedAt, &transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetByStatus retrieves all transactions with a specific status
func (r *TransactionRepo) GetByStatus(status string) ([]model.Transaction, error) {
	query := `
		SELECT id, user_id, type, status, amount, currency, source_account_id, dest_account_id, fee, exchange_rate, description, reference, processed_at, failure_reason, created_at, updated_at
		FROM transactions
		WHERE status = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(context.Background(), query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(
			&transaction.ID, &transaction.UserID, &transaction.Type, &transaction.Status,
			&transaction.Amount, &transaction.Currency, &transaction.SourceAccountID, &transaction.DestAccountID,
			&transaction.Fee, &transaction.ExchangeRate, &transaction.Description, &transaction.Reference,
			&transaction.ProcessedAt, &transaction.FailureReason,
			&transaction.CreatedAt, &transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetByAccount retrieves all transactions for an account
func (r *TransactionRepo) GetByAccount(ctx context.Context, accountID string, limit, offset int) ([]*model.Transaction, error) {
	query := `
		SELECT id, user_id, type, status, amount, currency, source_account_id, dest_account_id, fee, exchange_rate, description, reference, processed_at, failure_reason, created_at, updated_at
		FROM transactions
		WHERE source_account_id = $1 OR dest_account_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, accountID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []*model.Transaction{}
	for rows.Next() {
		transaction := &model.Transaction{}
		err := rows.Scan(
			&transaction.ID, &transaction.UserID, &transaction.Type, &transaction.Status,
			&transaction.Amount, &transaction.Currency, &transaction.SourceAccountID, &transaction.DestAccountID,
			&transaction.Fee, &transaction.ExchangeRate, &transaction.Description, &transaction.Reference,
			&transaction.ProcessedAt, &transaction.FailureReason,
			&transaction.CreatedAt, &transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetByTransfer retrieves all transactions for a transfer
func (r *TransactionRepo) GetByTransfer(ctx context.Context, transferID string) ([]*model.Transaction, error) {
	query := `
		SELECT id, user_id, type, status, amount, currency, source_account_id, dest_account_id, fee, exchange_rate, description, reference, processed_at, failure_reason, created_at, updated_at
		FROM transactions
		WHERE transfer_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, transferID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []*model.Transaction{}
	for rows.Next() {
		transaction := &model.Transaction{}
		err := rows.Scan(
			&transaction.ID, &transaction.UserID, &transaction.Type, &transaction.Status,
			&transaction.Amount, &transaction.Currency, &transaction.SourceAccountID, &transaction.DestAccountID,
			&transaction.Fee, &transaction.ExchangeRate, &transaction.Description, &transaction.Reference,
			&transaction.ProcessedAt, &transaction.FailureReason,
			&transaction.CreatedAt, &transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
