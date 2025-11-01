-- Drop indexes
DROP INDEX IF EXISTS idx_transfers_user_id;
DROP INDEX IF EXISTS idx_transfers_status;
DROP INDEX IF EXISTS idx_transfers_reference;

-- Drop transfers table
DROP TABLE IF EXISTS transfers;