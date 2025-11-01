-- Drop indexes
DROP INDEX IF EXISTS idx_money_requests_requester_id;
DROP INDEX IF EXISTS idx_money_requests_recipient_id;
DROP INDEX IF EXISTS idx_money_requests_status;
DROP INDEX IF EXISTS idx_money_requests_payment_link;

-- Drop money_requests table
DROP TABLE IF EXISTS money_requests;