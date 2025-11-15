-- Add email column back to beneficiaries table
ALTER TABLE beneficiaries ADD COLUMN IF NOT EXISTS email VARCHAR(255);