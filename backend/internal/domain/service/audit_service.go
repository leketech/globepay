package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"globepay/internal/domain/model"
	"globepay/internal/repository"
)

// AuditService provides audit logging functionality
type AuditService struct {
	auditRepo repository.AuditRepository
}

// NewAuditService creates a new audit service
func NewAuditService(auditRepo repository.AuditRepository) *AuditService {
	return &AuditService{
		auditRepo: auditRepo,
	}
}

// LogUserAction logs a user action
func (s *AuditService) LogUserAction(ctx context.Context, userID, action, tableName string, recordID string, oldValues, newValues interface{}) error {
	auditLog := &model.AuditLog{
		UserID:    userID,
		Action:    action,
		TableName: tableName,
		RecordID:  recordID,
		OldValues: make(map[string]interface{}),
		NewValues: make(map[string]interface{}),
		IPAddress: getClientIP(ctx),
		UserAgent: getUserAgent(ctx),
		CreatedAt: time.Now(),
	}

	// Convert old values to map
	if oldValues != nil {
		oldBytes, err := json.Marshal(oldValues)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(oldBytes, &auditLog.OldValues); err != nil {
			log.Printf("Failed to unmarshal old values: %v", err)
		}
	}

	// Convert new values to map
	if newValues != nil {
		newBytes, err := json.Marshal(newValues)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(newBytes, &auditLog.NewValues); err != nil {
			log.Printf("Failed to unmarshal new values: %v", err)
		}
	}

	return s.auditRepo.Create(ctx, auditLog)
}

// GetAuditLogsByUser retrieves audit logs for a user
func (s *AuditService) GetAuditLogsByUser(ctx context.Context, userID string, limit, offset int) ([]*model.AuditLog, error) {
	return s.auditRepo.GetByUser(ctx, userID, limit, offset)
}

// GetAuditLogsByAction retrieves audit logs for a specific action
func (s *AuditService) GetAuditLogsByAction(ctx context.Context, action string, limit, offset int) ([]*model.AuditLog, error) {
	return s.auditRepo.GetByAction(ctx, action, limit, offset)
}

// GetAuditLogsByTable retrieves audit logs for a specific table
func (s *AuditService) GetAuditLogsByTable(ctx context.Context, tableName string, limit, offset int) ([]*model.AuditLog, error) {
	return s.auditRepo.GetByTable(ctx, tableName, limit, offset)
}

// getClientIP extracts client IP from context
func getClientIP(ctx context.Context) string {
	// In a real implementation, this would extract the IP from the request
	// For now, we'll return a placeholder
	return "127.0.0.1"
}

// getUserAgent extracts user agent from context
func getUserAgent(ctx context.Context) string {
	// In a real implementation, this would extract the user agent from the request
	// For now, we'll return a placeholder
	return "Unknown"
}
