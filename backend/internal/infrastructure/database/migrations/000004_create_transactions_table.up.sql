-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    transfer_id UUID REFERENCES transfers(id) ON DELETE SET NULL,
    type VARCHAR(20) NOT NULL, -- DEPOSIT, WITHDRAWAL, TRANSFER, FEE
    status VARCHAR(20) DEFAULT 'completed',
    amount DECIMAL(15,2) NOT NULL,
    currency_code CHAR(3) NOT NULL,
    fee_amount DECIMAL(15,2) DEFAULT 0.00,
    description TEXT,
    reference_number VARCHAR(50) UNIQUE,
    processed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_transactions_account_id ON transactions(account_id);
CREATE INDEX IF NOT EXISTS idx_transactions_transfer_id ON transactions(transfer_id);
CREATE INDEX IF NOT EXISTS idx_transactions_type ON transactions(type);
CREATE INDEX IF NOT EXISTS idx_transactions_status ON transactions(status);
CREATE INDEX IF NOT EXISTS idx_transactions_reference ON transactions(reference_number);