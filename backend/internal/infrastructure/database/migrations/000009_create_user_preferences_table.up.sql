-- Create user_preferences table
CREATE TABLE IF NOT EXISTS user_preferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    email_notifications BOOLEAN NOT NULL DEFAULT true,
    push_notifications BOOLEAN NOT NULL DEFAULT false,
    sms_notifications BOOLEAN NOT NULL DEFAULT false,
    transaction_alerts BOOLEAN NOT NULL DEFAULT true,
    security_alerts BOOLEAN NOT NULL DEFAULT true,
    marketing_emails BOOLEAN NOT NULL DEFAULT false,
    two_factor_enabled BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index
CREATE INDEX IF NOT EXISTS idx_user_preferences_user_id ON user_preferences(user_id);

-- Insert default preferences for existing users
INSERT INTO user_preferences (user_id, email_notifications, push_notifications, sms_notifications, transaction_alerts, security_alerts, marketing_emails, two_factor_enabled)
SELECT id, true, false, false, true, true, false, false
FROM users
ON CONFLICT DO NOTHING;