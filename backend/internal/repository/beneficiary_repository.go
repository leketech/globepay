package repository

import (
	"context"
	"database/sql"
	"time"

	"globepay/internal/domain/model"
)

// BeneficiaryRepo implements BeneficiaryRepository
type BeneficiaryRepo struct {
	db *sql.DB
}

// NewBeneficiaryRepository creates a new BeneficiaryRepo
func NewBeneficiaryRepository(db *sql.DB) BeneficiaryRepository {
	return &BeneficiaryRepo{db: db}
}

// Create inserts a new beneficiary into the database
func (r *BeneficiaryRepo) Create(beneficiary *model.Beneficiary) error {
	query := `
		INSERT INTO beneficiaries (id, user_id, name, account_no, bank_name, bank_address, country, currency, swift_code, iban, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`

	now := time.Now()
	beneficiary.CreatedAt = now
	beneficiary.UpdatedAt = now

	return r.db.QueryRowContext(context.Background(), query, beneficiary.ID, beneficiary.UserID, beneficiary.Name, beneficiary.AccountNo, beneficiary.BankName, beneficiary.BankAddress, beneficiary.Country, beneficiary.Currency, beneficiary.SwiftCode, beneficiary.Iban, beneficiary.CreatedAt, beneficiary.UpdatedAt).Scan(&beneficiary.ID)
}

// GetByID retrieves a beneficiary by ID
func (r *BeneficiaryRepo) GetByID(id string) (*model.Beneficiary, error) {
	query := `
		SELECT id, user_id, name, account_no, bank_name, bank_address, country, currency, swift_code, iban, created_at, updated_at
		FROM beneficiaries
		WHERE id = $1
	`

	beneficiary := &model.Beneficiary{}
	err := r.db.QueryRowContext(context.Background(), query, id).Scan(
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

// GetByUser retrieves beneficiaries for a user
func (r *BeneficiaryRepo) GetByUser(ctx context.Context, userID string) ([]*model.Beneficiary, error) {
	query := `
		SELECT id, user_id, name, account_no, bank_name, bank_address, country, currency, swift_code, iban, created_at, updated_at
		FROM beneficiaries
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	beneficiaries := []*model.Beneficiary{}
	for rows.Next() {
		beneficiary := &model.Beneficiary{}
		err := rows.Scan(
			&beneficiary.ID, &beneficiary.UserID, &beneficiary.Name, &beneficiary.AccountNo, &beneficiary.BankName, &beneficiary.BankAddress, &beneficiary.Country, &beneficiary.Currency, &beneficiary.SwiftCode, &beneficiary.Iban, &beneficiary.CreatedAt, &beneficiary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		beneficiaries = append(beneficiaries, beneficiary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return beneficiaries, nil
}

// GetByNameAndUser retrieves a beneficiary by name and user
func (r *BeneficiaryRepo) GetByNameAndUser(ctx context.Context, name, userID string) (*model.Beneficiary, error) {
	query := `
		SELECT id, user_id, name, account_no, bank_name, bank_address, country, currency, swift_code, iban, created_at, updated_at
		FROM beneficiaries
		WHERE name = $1 AND user_id = $2
	`

	beneficiary := &model.Beneficiary{}
	err := r.db.QueryRowContext(ctx, query, name, userID).Scan(
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

// Update updates a beneficiary in the database
func (r *BeneficiaryRepo) Update(beneficiary *model.Beneficiary) error {
	query := `
		UPDATE beneficiaries
		SET name = $1, account_no = $2, bank_name = $3, bank_address = $4, country = $5, currency = $6, swift_code = $7, iban = $8, updated_at = $9
		WHERE id = $10
	`

	beneficiary.UpdatedAt = time.Now()
	result, err := r.db.ExecContext(context.Background(), query, beneficiary.Name, beneficiary.AccountNo, beneficiary.BankName, beneficiary.BankAddress, beneficiary.Country, beneficiary.Currency, beneficiary.SwiftCode, beneficiary.Iban, beneficiary.UpdatedAt, beneficiary.ID)
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

// Delete removes a beneficiary from the database
func (r *BeneficiaryRepo) Delete(id string) error {
	query := `DELETE FROM beneficiaries WHERE id = $1`

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
