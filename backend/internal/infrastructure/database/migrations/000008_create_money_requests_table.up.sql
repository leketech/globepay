-- Create money_requests table
CREATE TABLE IF NOT EXISTS money_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    requester_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recipient_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount DECIMAL(15,2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    payment_link VARCHAR(255),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    paid_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_money_requests_requester_id ON money_requests(requester_id);
CREATE INDEX IF NOT EXISTS idx_money_requests_recipient_id ON money_requests(recipient_id);
CREATE INDEX IF NOT EXISTS idx_money_requests_status ON money_requests(status);
CREATE INDEX IF NOT EXISTS idx_money_requests_payment_link ON money_requests(payment_link);