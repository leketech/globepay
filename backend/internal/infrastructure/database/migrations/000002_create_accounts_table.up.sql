-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    currency_code CHAR(3) NOT NULL,
    balance DECIMAL(15,2) DEFAULT 0.00,
    account_number VARCHAR(50) UNIQUE NOT NULL,
    account_type VARCHAR(20) DEFAULT 'checking',
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_accounts_user_id ON accounts(user_id);
CREATE INDEX IF NOT EXISTS idx_accounts_currency ON accounts(currency_code);
CREATE INDEX IF NOT EXISTS idx_accounts_status ON accounts(status);