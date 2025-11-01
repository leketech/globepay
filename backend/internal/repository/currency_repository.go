package repository

import (
	"database/sql"

	"globepay/internal/domain/model"
)

// CurrencyRepo implements CurrencyRepository
type CurrencyRepo struct {
	db *sql.DB
}

// NewCurrencyRepository creates a new CurrencyRepo
func NewCurrencyRepository(db *sql.DB) CurrencyRepository {
	return &CurrencyRepo{db: db}
}

// GetAll retrieves all currencies from the database
func (r *CurrencyRepo) GetAll(ctx interface{}) ([]*model.Currency, error) {
	query := `
		SELECT code, name, symbol, created_at, updated_at
		FROM currencies
		ORDER BY code
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currencies := []*model.Currency{}
	for rows.Next() {
		currency := &model.Currency{}
		err := rows.Scan(
			&currency.Code, &currency.Name, &currency.Symbol, &currency.CreatedAt, &currency.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, currency)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return currencies, nil
}

// GetByCode retrieves a currency by code
func (r *CurrencyRepo) GetByCode(ctx interface{}, code string) (*model.Currency, error) {
	query := `
		SELECT code, name, symbol, created_at, updated_at
		FROM currencies
		WHERE code = $1
	`

	currency := &model.Currency{}
	err := r.db.QueryRow(query, code).Scan(
		&currency.Code, &currency.Name, &currency.Symbol, &currency.CreatedAt, &currency.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return currency, nil
}