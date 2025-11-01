-- Drop indexes
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_kyc_status;
DROP INDEX IF EXISTS idx_users_account_status;

-- Drop users table
DROP TABLE IF EXISTS users;