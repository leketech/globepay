package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"globepay/internal/domain/model"
)

// AuditRepo implements AuditRepository
type AuditRepo struct {
	db *sql.DB
}

// NewAuditRepository creates a new AuditRepo
func NewAuditRepository(db *sql.DB) AuditRepository {
	return &AuditRepo{db: db}
}

// Create inserts a new audit log into the database
func (r *AuditRepo) Create(ctx context.Context, auditLog *model.AuditLog) error {
	query := `
		INSERT INTO audit_logs (id, user_id, action, table_name, record_id, old_values, new_values, ip_address, user_agent, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`

	// Convert maps to JSON
	var oldValuesJSON, newValuesJSON []byte
	var err error
	
	if auditLog.OldValues != nil {
		oldValuesJSON, err = json.Marshal(auditLog.OldValues)
		if err != nil {
			return err
		}
	}
	
	if auditLog.NewValues != nil {
		newValuesJSON, err = json.Marshal(auditLog.NewValues)
		if err != nil {
			return err
		}
	}

	return r.db.QueryRowContext(ctx, query, auditLog.ID, auditLog.UserID, auditLog.Action, auditLog.TableName, auditLog.RecordID, oldValuesJSON, newValuesJSON, auditLog.IPAddress, auditLog.UserAgent, auditLog.CreatedAt).Scan(&auditLog.ID)
}

// GetByUser retrieves audit logs for a user
func (r *AuditRepo) GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.AuditLog, error) {
	query := `
		SELECT id, user_id, action, table_name, record_id, old_values, new_values, ip_address, user_agent, created_at
		FROM audit_logs
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	auditLogs := []*model.AuditLog{}
	for rows.Next() {
		auditLog := &model.AuditLog{
			OldValues: make(map[string]interface{}),
			NewValues: make(map[string]interface{}),
		}
		
		var oldValuesJSON, newValuesJSON []byte
		
		err := rows.Scan(
			&auditLog.ID, &auditLog.UserID, &auditLog.Action, &auditLog.TableName, &auditLog.RecordID, &oldValuesJSON, &newValuesJSON, &auditLog.IPAddress, &auditLog.UserAgent, &auditLog.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Convert JSON to maps
		if len(oldValuesJSON) > 0 {
			json.Unmarshal(oldValuesJSON, &auditLog.OldValues)
		}
		
		if len(newValuesJSON) > 0 {
			json.Unmarshal(newValuesJSON, &auditLog.NewValues)
		}
		
		auditLogs = append(auditLogs, auditLog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return auditLogs, nil
}

// GetByAction retrieves audit logs for a specific action
func (r *AuditRepo) GetByAction(ctx context.Context, action string, limit, offset int) ([]*model.AuditLog, error) {
	query := `
		SELECT id, user_id, action, table_name, record_id, old_values, new_values, ip_address, user_agent, created_at
		FROM audit_logs
		WHERE action = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, action, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	auditLogs := []*model.AuditLog{}
	for rows.Next() {
		auditLog := &model.AuditLog{
			OldValues: make(map[string]interface{}),
			NewValues: make(map[string]interface{}),
		}
		
		var oldValuesJSON, newValuesJSON []byte
		
		err := rows.Scan(
			&auditLog.ID, &auditLog.UserID, &auditLog.Action, &auditLog.TableName, &auditLog.RecordID, &oldValuesJSON, &newValuesJSON, &auditLog.IPAddress, &auditLog.UserAgent, &auditLog.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Convert JSON to maps
		if len(oldValuesJSON) > 0 {
			json.Unmarshal(oldValuesJSON, &auditLog.OldValues)
		}
		
		if len(newValuesJSON) > 0 {
			json.Unmarshal(newValuesJSON, &auditLog.NewValues)
		}
		
		auditLogs = append(auditLogs, auditLog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return auditLogs, nil
}

// GetByTable retrieves audit logs for a specific table
func (r *AuditRepo) GetByTable(ctx context.Context, tableName string, limit, offset int) ([]*model.AuditLog, error) {
	query := `
		SELECT id, user_id, action, table_name, record_id, old_values, new_values, ip_address, user_agent, created_at
		FROM audit_logs
		WHERE table_name = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, tableName, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	auditLogs := []*model.AuditLog{}
	for rows.Next() {
		auditLog := &model.AuditLog{
			OldValues: make(map[string]interface{}),
			NewValues: make(map[string]interface{}),
		}
		
		var oldValuesJSON, newValuesJSON []byte
		
		err := rows.Scan(
			&auditLog.ID, &auditLog.UserID, &auditLog.Action, &auditLog.TableName, &auditLog.RecordID, &oldValuesJSON, &newValuesJSON, &auditLog.IPAddress, &auditLog.UserAgent, &auditLog.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Convert JSON to maps
		if len(oldValuesJSON) > 0 {
			json.Unmarshal(oldValuesJSON, &auditLog.OldValues)
		}
		
		if len(newValuesJSON) > 0 {
			json.Unmarshal(newValuesJSON, &auditLog.NewValues)
		}
		
		auditLogs = append(auditLogs, auditLog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return auditLogs, nil
}