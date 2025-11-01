-- Drop indexes
DROP INDEX IF EXISTS idx_beneficiaries_user_id;
DROP INDEX IF EXISTS idx_beneficiaries_name;

-- Drop beneficiaries table
DROP TABLE IF EXISTS beneficiaries;