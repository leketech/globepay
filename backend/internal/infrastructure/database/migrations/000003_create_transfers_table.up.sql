-- Create transfers table
CREATE TABLE IF NOT EXISTS transfers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recipient_name VARCHAR(255) NOT NULL,
    recipient_email VARCHAR(255),
    recipient_country CHAR(2) NOT NULL,
    recipient_bank_name VARCHAR(255) NOT NULL,
    recipient_account_number VARCHAR(100) NOT NULL,
    recipient_swift_code VARCHAR(20),
    source_currency CHAR(3) NOT NULL,
    destination_currency CHAR(3) NOT NULL,
    source_amount DECIMAL(15,2) NOT NULL,
    destination_amount DECIMAL(15,2) NOT NULL,
    exchange_rate DECIMAL(10,6) NOT NULL,
    fee_amount DECIMAL(15,2) NOT NULL,
    purpose VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    reference_number VARCHAR(50) UNIQUE,
    estimated_arrival TIMESTAMP WITH TIME ZONE,
    processed_at TIMESTAMP WITH TIME ZONE,
    failure_reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_transfers_user_id ON transfers(user_id);
CREATE INDEX IF NOT EXISTS idx_transfers_status ON transfers(status);
CREATE INDEX IF NOT EXISTS idx_transfers_reference ON transfers(reference_number);