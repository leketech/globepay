-- Remove email column from beneficiaries table
ALTER TABLE beneficiaries DROP COLUMN IF EXISTS email;