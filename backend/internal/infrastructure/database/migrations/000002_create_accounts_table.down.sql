-- Drop indexes
DROP INDEX IF EXISTS idx_accounts_user_id;
DROP INDEX IF EXISTS idx_accounts_currency;
DROP INDEX IF EXISTS idx_accounts_status;

-- Drop accounts table
DROP TABLE IF EXISTS accounts;