-- Drop indexes
DROP INDEX IF EXISTS idx_audit_logs_user_id;
DROP INDEX IF EXISTS idx_audit_logs_action;
DROP INDEX IF EXISTS idx_audit_logs_table_name;
DROP INDEX IF EXISTS idx_audit_logs_created_at;

-- Drop audit_logs table
DROP TABLE IF EXISTS audit_logs;