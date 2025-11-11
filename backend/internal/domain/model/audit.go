package model

import (
	"time"
)

// AuditLog represents an audit log entry
type AuditLog struct {
	ID        string                 `json:"id" db:"id"`
	UserID    string                 `json:"user_id" db:"user_id"`
	Action    string                 `json:"action" db:"action"`
	TableName string                 `json:"table_name" db:"table_name"`
	RecordID  string                 `json:"record_id" db:"record_id"`
	OldValues map[string]interface{} `json:"old_values" db:"old_values"`
	NewValues map[string]interface{} `json:"new_values" db:"new_values"`
	IPAddress string                 `json:"ip_address" db:"ip_address"`
	UserAgent string                 `json:"user_agent" db:"user_agent"`
	CreatedAt time.Time              `json:"created_at" db:"created_at"`
}

// NewAuditLog creates a new audit log entry
func NewAuditLog(
	userID, action, tableName, recordID, ipAddress, userAgent string,
	oldValues, newValues map[string]interface{},
) *AuditLog {
	return &AuditLog{
		UserID:    userID,
		Action:    action,
		TableName: tableName,
		RecordID:  recordID,
		OldValues: oldValues,
		NewValues: newValues,
		IPAddress: ipAddress,
		UserAgent: userAgent,
		CreatedAt: time.Now(),
	}
}
