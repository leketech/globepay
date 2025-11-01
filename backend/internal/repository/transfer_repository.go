package repository

import (
	"database/sql"
	"time"

	"globepay/internal/domain/model"
)

// TransferRepo implements TransferRepository
type TransferRepo struct {
	db *sql.DB
}

// NewTransferRepository creates a new TransferRepo
func NewTransferRepository(db *sql.DB) TransferRepository {
	return &TransferRepo{db: db}
}

// Create inserts a new transfer into the database
func (r *TransferRepo) Create(transfer *model.Transfer) error {
	query := `
		INSERT INTO transfers (user_id, recipient_name, recipient_email, recipient_country, recipient_bank_name, recipient_account_number, recipient_swift_code, source_currency, destination_currency, source_amount, destination_amount, exchange_rate, fee_amount, purpose, status, reference_number, estimated_arrival, created_at, updated_at, processed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
		RETURNING id
	`

	now := time.Now()
	transfer.CreatedAt = now
	transfer.UpdatedAt = now

	return r.db.QueryRow(query, transfer.UserID, transfer.RecipientName, transfer.RecipientEmail, transfer.RecipientCountry, transfer.RecipientBankName, transfer.RecipientAccountNumber, transfer.RecipientSwiftCode, transfer.SourceCurrency, transfer.DestCurrency, transfer.SourceAmount, transfer.DestAmount, transfer.ExchangeRate, transfer.FeeAmount, transfer.Purpose, transfer.Status, transfer.ReferenceNumber, transfer.EstimatedArrival, transfer.CreatedAt, transfer.UpdatedAt, transfer.ProcessedAt).Scan(&transfer.ID)
}

// GetByID retrieves a transfer by ID
func (r *TransferRepo) GetByID(id string) (*model.Transfer, error) {
	query := `
		SELECT id, user_id, recipient_name, recipient_email, recipient_country, recipient_bank_name, recipient_account_number, recipient_swift_code, source_currency, destination_currency, source_amount, destination_amount, exchange_rate, fee_amount, purpose, status, reference_number, estimated_arrival, created_at, updated_at, processed_at
		FROM transfers
		WHERE id = $1
	`

	transfer := &model.Transfer{}
	err := r.db.QueryRow(query, id).Scan(
		&transfer.ID, &transfer.UserID, &transfer.RecipientName, &transfer.RecipientEmail, &transfer.RecipientCountry, &transfer.RecipientBankName, &transfer.RecipientAccountNumber, &transfer.RecipientSwiftCode, &transfer.SourceCurrency, &transfer.DestCurrency, &transfer.SourceAmount, &transfer.DestAmount, &transfer.ExchangeRate, &transfer.FeeAmount, &transfer.Purpose, &transfer.Status, &transfer.ReferenceNumber, &transfer.EstimatedArrival, &transfer.CreatedAt, &transfer.UpdatedAt, &transfer.ProcessedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return transfer, nil
}

// GetByUser retrieves all transfers for a user (as sender or receiver)
func (r *TransferRepo) GetByUser(ctx interface{}, userID string, limit, offset int) ([]*model.Transfer, error) {
	query := `
		SELECT id, user_id, recipient_name, recipient_email, recipient_country, recipient_bank_name, recipient_account_number, recipient_swift_code, source_currency, destination_currency, source_amount, destination_amount, exchange_rate, fee_amount, purpose, status, reference_number, estimated_arrival, created_at, updated_at, processed_at
		FROM transfers
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := []*model.Transfer{}
	for rows.Next() {
		transfer := &model.Transfer{}
		err := rows.Scan(
			&transfer.ID, &transfer.UserID, &transfer.RecipientName, &transfer.RecipientEmail, &transfer.RecipientCountry, &transfer.RecipientBankName, &transfer.RecipientAccountNumber, &transfer.RecipientSwiftCode, &transfer.SourceCurrency, &transfer.DestCurrency, &transfer.SourceAmount, &transfer.DestAmount, &transfer.ExchangeRate, &transfer.FeeAmount, &transfer.Purpose, &transfer.Status, &transfer.ReferenceNumber, &transfer.EstimatedArrival, &transfer.CreatedAt, &transfer.UpdatedAt, &transfer.ProcessedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}

// Update updates a transfer in the database
func (r *TransferRepo) Update(transfer *model.Transfer) error {
	query := `
		UPDATE transfers
		SET user_id = $1, recipient_name = $2, recipient_email = $3, recipient_country = $4, recipient_bank_name = $5, recipient_account_number = $6, recipient_swift_code = $7, source_currency = $8, destination_currency = $9, source_amount = $10, destination_amount = $11, exchange_rate = $12, fee_amount = $13, purpose = $14, status = $15, reference_number = $16, estimated_arrival = $17, updated_at = $18, processed_at = $19
		WHERE id = $20
	`

	transfer.UpdatedAt = time.Now()
	result, err := r.db.Exec(query, transfer.UserID, transfer.RecipientName, transfer.RecipientEmail, transfer.RecipientCountry, transfer.RecipientBankName, transfer.RecipientAccountNumber, transfer.RecipientSwiftCode, transfer.SourceCurrency, transfer.DestCurrency, transfer.SourceAmount, transfer.DestAmount, transfer.ExchangeRate, transfer.FeeAmount, transfer.Purpose, transfer.Status, transfer.ReferenceNumber, transfer.EstimatedArrival, transfer.UpdatedAt, transfer.ProcessedAt, transfer.ID)
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

// Delete removes a transfer from the database
func (r *TransferRepo) Delete(id string) error {
	query := `DELETE FROM transfers WHERE id = $1`

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

// GetAll retrieves all transfers from the database
func (r *TransferRepo) GetAll() ([]model.Transfer, error) {
	query := `
		SELECT id, user_id, recipient_name, recipient_email, recipient_country, recipient_bank_name, recipient_account_number, recipient_swift_code, source_currency, destination_currency, source_amount, destination_amount, exchange_rate, fee_amount, purpose, status, reference_number, estimated_arrival, created_at, updated_at, processed_at
		FROM transfers
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := []model.Transfer{}
	for rows.Next() {
		var transfer model.Transfer
		err := rows.Scan(
			&transfer.ID, &transfer.UserID, &transfer.RecipientName, &transfer.RecipientEmail, &transfer.RecipientCountry, &transfer.RecipientBankName, &transfer.RecipientAccountNumber, &transfer.RecipientSwiftCode, &transfer.SourceCurrency, &transfer.DestCurrency, &transfer.SourceAmount, &transfer.DestAmount, &transfer.ExchangeRate, &transfer.FeeAmount, &transfer.Purpose, &transfer.Status, &transfer.ReferenceNumber, &transfer.EstimatedArrival, &transfer.CreatedAt, &transfer.UpdatedAt, &transfer.ProcessedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}

// GetByStatus retrieves all transfers with a specific status
func (r *TransferRepo) GetByStatus(status string) ([]model.Transfer, error) {
	query := `
		SELECT id, user_id, recipient_name, recipient_email, recipient_country, recipient_bank_name, recipient_account_number, recipient_swift_code, source_currency, destination_currency, source_amount, destination_amount, exchange_rate, fee_amount, purpose, status, reference_number, estimated_arrival, created_at, updated_at, processed_at
		FROM transfers
		WHERE status = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := []model.Transfer{}
	for rows.Next() {
		var transfer model.Transfer
		err := rows.Scan(
			&transfer.ID, &transfer.UserID, &transfer.RecipientName, &transfer.RecipientEmail, &transfer.RecipientCountry, &transfer.RecipientBankName, &transfer.RecipientAccountNumber, &transfer.RecipientSwiftCode, &transfer.SourceCurrency, &transfer.DestCurrency, &transfer.SourceAmount, &transfer.DestAmount, &transfer.ExchangeRate, &transfer.FeeAmount, &transfer.Purpose, &transfer.Status, &transfer.ReferenceNumber, &transfer.EstimatedArrival, &transfer.CreatedAt, &transfer.UpdatedAt, &transfer.ProcessedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}

// GetByReferenceNumber retrieves a transfer by reference number
func (r *TransferRepo) GetByReferenceNumber(referenceNumber string) (*model.Transfer, error) {
	query := `
		SELECT id, user_id, recipient_name, recipient_email, recipient_country, recipient_bank_name, recipient_account_number, recipient_swift_code, source_currency, destination_currency, source_amount, destination_amount, exchange_rate, fee_amount, purpose, status, reference_number, estimated_arrival, created_at, updated_at, processed_at
		FROM transfers
		WHERE reference_number = $1
	`

	transfer := &model.Transfer{}
	err := r.db.QueryRow(query, referenceNumber).Scan(
		&transfer.ID, &transfer.UserID, &transfer.RecipientName, &transfer.RecipientEmail, &transfer.RecipientCountry, &transfer.RecipientBankName, &transfer.RecipientAccountNumber, &transfer.RecipientSwiftCode, &transfer.SourceCurrency, &transfer.DestCurrency, &transfer.SourceAmount, &transfer.DestAmount, &transfer.ExchangeRate, &transfer.FeeAmount, &transfer.Purpose, &transfer.Status, &transfer.ReferenceNumber, &transfer.EstimatedArrival, &transfer.CreatedAt, &transfer.UpdatedAt, &transfer.ProcessedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return transfer, nil
}

// GetByNameAndUser retrieves a beneficiary by name and user
func (r *TransferRepo) GetByNameAndUser(ctx interface{}, name, userID string) (*model.Beneficiary, error) {
	query := `
		SELECT id, user_id, name, account_no, bank_name, bank_address, country, currency, swift_code, iban, created_at, updated_at
		FROM beneficiaries
		WHERE name = $1 AND user_id = $2
	`

	beneficiary := &model.Beneficiary{}
	err := r.db.QueryRow(query, name, userID).Scan(
		&beneficiary.ID, &beneficiary.UserID, &beneficiary.Name, &beneficiary.AccountNo, &beneficiary.BankName, &beneficiary.BankAddress, &beneficiary.Country, &beneficiary.Currency, &beneficiary.SwiftCode, &beneficiary.Iban, &beneficiary.CreatedAt, &beneficiary.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return beneficiary, nil
}