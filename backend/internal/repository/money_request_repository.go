package repository

import (
	"context"
	"database/sql"
	"globepay/internal/domain/model"
	"globepay/internal/utils"
	"time"
)

// MoneyRequestRepo implements MoneyRequestRepository
type MoneyRequestRepo struct {
	db *sql.DB
}

// NewMoneyRequestRepository creates a new money request repository
func NewMoneyRequestRepository(db *sql.DB) MoneyRequestRepository {
	return &MoneyRequestRepo{db: db}
}

// Create inserts a new money request into the database
func (r *MoneyRequestRepo) Create(ctx context.Context, request *model.MoneyRequest) error {
	query := `
		INSERT INTO money_requests (
			id, requester_id, recipient_id, amount, currency, description, 
			status, payment_link, expires_at, created_at, updated_at, paid_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	request.ID = utils.GenerateUUID()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	_, err := r.db.ExecContext(
		ctx, query,
		request.ID, request.RequesterID, request.RecipientID, request.Amount,
		request.Currency, request.Description, request.Status, request.PaymentLink,
		request.ExpiresAt, request.CreatedAt, request.UpdatedAt, request.PaidAt,
	)

	return err
}

// GetByID retrieves a money request by its ID
func (r *MoneyRequestRepo) GetByID(ctx context.Context, id string) (*model.MoneyRequest, error) {
	query := `
		SELECT id, requester_id, recipient_id, amount, currency, description,
		       status, payment_link, expires_at, created_at, updated_at, paid_at
		FROM money_requests
		WHERE id = $1
	`

	request := &model.MoneyRequest{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&request.ID, &request.RequesterID, &request.RecipientID, &request.Amount,
		&request.Currency, &request.Description, &request.Status, &request.PaymentLink,
		&request.ExpiresAt, &request.CreatedAt, &request.UpdatedAt, &request.PaidAt,
	)

	if err != nil {
		return nil, err
	}

	return request, nil
}

// GetByRequester retrieves all money requests made by a user
func (r *MoneyRequestRepo) GetByRequester(ctx context.Context, requesterID string) ([]*model.MoneyRequest, error) {
	query := `
		SELECT id, requester_id, recipient_id, amount, currency, description,
		       status, payment_link, expires_at, created_at, updated_at, paid_at
		FROM money_requests
		WHERE requester_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, requesterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []*model.MoneyRequest
	for rows.Next() {
		request := &model.MoneyRequest{}
		err := rows.Scan(
			&request.ID, &request.RequesterID, &request.RecipientID, &request.Amount,
			&request.Currency, &request.Description, &request.Status, &request.PaymentLink,
			&request.ExpiresAt, &request.CreatedAt, &request.UpdatedAt, &request.PaidAt,
		)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}

// GetByRecipient retrieves all money requests for a recipient
func (r *MoneyRequestRepo) GetByRecipient(ctx context.Context, recipientID string) ([]*model.MoneyRequest, error) {
	query := `
		SELECT id, requester_id, recipient_id, amount, currency, description,
		       status, payment_link, expires_at, created_at, updated_at, paid_at
		FROM money_requests
		WHERE recipient_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, recipientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []*model.MoneyRequest
	for rows.Next() {
		request := &model.MoneyRequest{}
		err := rows.Scan(
			&request.ID, &request.RequesterID, &request.RecipientID, &request.Amount,
			&request.Currency, &request.Description, &request.Status, &request.PaymentLink,
			&request.ExpiresAt, &request.CreatedAt, &request.UpdatedAt, &request.PaidAt,
		)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}

// UpdateStatus updates the status of a money request
func (r *MoneyRequestRepo) UpdateStatus(ctx context.Context, id, status string, paidAt *time.Time) error {
	query := `
		UPDATE money_requests
		SET status = $1, paid_at = $2, updated_at = $3
		WHERE id = $4
	`

	_, err := r.db.ExecContext(ctx, query, status, paidAt, time.Now(), id)
	return err
}

// UpdatePaymentLink updates the payment link for a money request
func (r *MoneyRequestRepo) UpdatePaymentLink(ctx context.Context, id, paymentLink string) error {
	query := `
		UPDATE money_requests
		SET payment_link = $1, updated_at = $2
		WHERE id = $3
	`

	_, err := r.db.ExecContext(ctx, query, paymentLink, time.Now(), id)
	return err
}