# Database Schema Documentation

This document provides detailed information about the database schema used in the Globepay application.

## Overview

Globepay uses PostgreSQL as its primary database with a well-structured schema designed for financial transactions, user management, and compliance requirements.

## Database Design Principles

1. **Normalization** - Third normal form (3NF) for data integrity
2. **Indexing** - Strategic indexes for performance optimization
3. **Constraints** - Database-level constraints for data validation
4. **Auditing** - Timestamps and user tracking for all changes
5. **Security** - Role-based access control and encryption

## Entity Relationship Diagram

```
┌─────────────┐       ┌─────────────┐
│    Users    │──────▶│  Profiles   │
└─────────────┘       └─────────────┘
       │
       ▼
┌─────────────┐       ┌─────────────┐
│  Accounts   │◀──────│ Currencies  │
└─────────────┘       └─────────────┘
       │
       ▼
┌─────────────┐       ┌─────────────┐
│ Transfers   │──────▶│ Beneficiaries│
└─────────────┘       └─────────────┘
       │
       ▼
┌─────────────┐       ┌─────────────┐
│Transactions │──────▶│   Ledgers   │
└─────────────┘       └─────────────┘
       │
       ▼
┌─────────────┐       ┌─────────────┐
│Audit Logs   │       │ Exch. Rates │
└─────────────┘       └─────────────┘
```

## Tables

### 1. Users Table

Stores user account information and authentication details.

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20),
    date_of_birth DATE,
    country_code CHAR(2),
    kyc_status VARCHAR(20) DEFAULT 'pending',
    account_status VARCHAR(20) DEFAULT 'active',
    email_verified BOOLEAN DEFAULT FALSE,
    phone_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_kyc_status ON users(kyc_status);
CREATE INDEX idx_users_account_status ON users(account_status);
```

**Fields:**
- `id`: Unique identifier (UUID)
- `email`: User's email address (unique)
- `password_hash`: Bcrypt hashed password
- `first_name`: User's first name
- `last_name`: User's last name
- `phone_number`: User's phone number
- `date_of_birth`: User's date of birth
- `country_code`: ISO 3166-1 alpha-2 country code
- `kyc_status`: KYC verification status (pending, verified, rejected)
- `account_status`: Account status (active, suspended, closed)
- `email_verified`: Email verification status
- `phone_verified`: Phone verification status
- `created_at`: Record creation timestamp
- `updated_at`: Record last update timestamp

### 2. Profiles Table

Stores extended user profile information.

```sql
CREATE TABLE profiles (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    address_line1 VARCHAR(255),
    address_line2 VARCHAR(255),
    city VARCHAR(100),
    state_province VARCHAR(100),
    postal_code VARCHAR(20),
    id_document_type VARCHAR(50),
    id_document_number VARCHAR(100),
    id_document_expiry DATE,
    id_document_front_url TEXT,
    id_document_back_url TEXT,
    selfie_url TEXT,
    occupation VARCHAR(100),
    annual_income DECIMAL(15,2),
    source_of_funds VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_profiles_user_id ON profiles(user_id);
```

### 3. Accounts Table

Stores user financial accounts.

```sql
CREATE TABLE accounts (
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

-- Indexes
CREATE INDEX idx_accounts_user_id ON accounts(user_id);
CREATE INDEX idx_accounts_currency ON accounts(currency_code);
CREATE INDEX idx_accounts_status ON accounts(status);
```

### 4. Currencies Table

Stores supported currencies and their properties.

```sql
CREATE TABLE currencies (
    code CHAR(3) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    decimal_places INTEGER DEFAULT 2,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Sample data
INSERT INTO currencies (code, name, symbol) VALUES 
('USD', 'US Dollar', '$'),
('EUR', 'Euro', '€'),
('GBP', 'British Pound', '£'),
('JPY', 'Japanese Yen', '¥');
```

### 5. Transfers Table

Stores international money transfer records.

```sql
CREATE TABLE transfers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recipient_name VARCHAR(255) NOT NULL,
    recipient_email VARCHAR(255),
    recipient_country CHAR(2) NOT NULL,
    recipient_bank_name VARCHAR(255) NOT NULL,
    recipient_account_number VARCHAR(100) NOT NULL,
    recipient_swift_code VARCHAR(20),
    source_currency CHAR(3) NOT NULL REFERENCES currencies(code),
    destination_currency CHAR(3) NOT NULL REFERENCES currencies(code),
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

-- Indexes
CREATE INDEX idx_transfers_user_id ON transfers(user_id);
CREATE INDEX idx_transfers_status ON transfers(status);
CREATE INDEX idx_transfers_created_at ON transfers(created_at);
CREATE INDEX idx_transfers_reference ON transfers(reference_number);
```

### 6. Beneficiaries Table

Stores user's saved beneficiaries for quick transfers.

```sql
CREATE TABLE beneficiaries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    country CHAR(2) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(100) NOT NULL,
    swift_code VARCHAR(20),
    iban VARCHAR(50),
    bank_address TEXT,
    currency CHAR(3),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_beneficiaries_user_id ON beneficiaries(user_id);
CREATE INDEX idx_beneficiaries_name ON beneficiaries(name);
```

### 7. Transactions Table

Stores detailed transaction records.

```sql
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    transfer_id UUID REFERENCES transfers(id) ON DELETE SET NULL,
    type VARCHAR(20) NOT NULL, -- DEPOSIT, WITHDRAWAL, TRANSFER, FEE
    status VARCHAR(20) DEFAULT 'completed',
    amount DECIMAL(15,2) NOT NULL,
    currency_code CHAR(3) NOT NULL REFERENCES currencies(code),
    fee_amount DECIMAL(15,2) DEFAULT 0.00,
    description TEXT,
    reference_number VARCHAR(50) UNIQUE,
    processed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_account_id ON transactions(account_id);
CREATE INDEX idx_transactions_type ON transactions(type);
CREATE INDEX idx_transactions_status ON transactions(status);
CREATE INDEX idx_transactions_reference ON transactions(reference_number);
```

### 8. Ledgers Table

Stores double-entry bookkeeping records.

```sql
CREATE TABLE ledgers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    transaction_id UUID NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    entry_type VARCHAR(10) NOT NULL, -- DEBIT, CREDIT
    amount DECIMAL(15,2) NOT NULL,
    currency_code CHAR(3) NOT NULL REFERENCES currencies(code),
    balance_after DECIMAL(15,2) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_ledgers_account_id ON ledgers(account_id);
CREATE INDEX idx_ledgers_transaction_id ON ledgers(transaction_id);
CREATE INDEX idx_ledgers_created_at ON ledgers(created_at);
```

### 9. Exchange Rates Table

Stores historical exchange rates.

```sql
CREATE TABLE exchange_rates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_currency CHAR(3) NOT NULL REFERENCES currencies(code),
    to_currency CHAR(3) NOT NULL REFERENCES currencies(code),
    rate DECIMAL(10,6) NOT NULL,
    fee_percentage DECIMAL(5,4) DEFAULT 0.0000,
    valid_from TIMESTAMP WITH TIME ZONE NOT NULL,
    valid_to TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_exchange_rates_currencies ON exchange_rates(from_currency, to_currency);
CREATE INDEX idx_exchange_rates_valid_from ON exchange_rates(valid_from);
```

### 10. Audit Logs Table

Stores audit trail of all important operations.

```sql
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    action VARCHAR(100) NOT NULL,
    table_name VARCHAR(100) NOT NULL,
    record_id UUID,
    old_values JSONB,
    new_values JSONB,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_table_name ON audit_logs(table_name);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);
```

## Constraints and Validation

### Check Constraints

```sql
-- Ensure positive amounts
ALTER TABLE accounts ADD CONSTRAINT chk_account_balance_non_negative 
CHECK (balance >= 0);

ALTER TABLE transfers ADD CONSTRAINT chk_transfer_amounts_positive 
CHECK (source_amount > 0 AND destination_amount > 0);

ALTER TABLE transactions ADD CONSTRAINT chk_transaction_amount_positive 
CHECK (amount > 0);

-- Ensure valid statuses
ALTER TABLE users ADD CONSTRAINT chk_user_kyc_status 
CHECK (kyc_status IN ('pending', 'verified', 'rejected'));

ALTER TABLE users ADD CONSTRAINT chk_user_account_status 
CHECK (account_status IN ('active', 'suspended', 'closed'));

ALTER TABLE transfers ADD CONSTRAINT chk_transfer_status 
CHECK (status IN ('pending', 'processing', 'completed', 'failed', 'cancelled'));

ALTER TABLE transactions ADD CONSTRAINT chk_transaction_type 
CHECK (type IN ('DEPOSIT', 'WITHDRAWAL', 'TRANSFER', 'FEE'));

ALTER TABLE transactions ADD CONSTRAINT chk_transaction_status 
CHECK (status IN ('pending', 'processing', 'completed', 'failed', 'cancelled'));
```

### Foreign Key Constraints

All relationships are enforced through foreign key constraints with appropriate cascading rules:

- **CASCADE DELETE**: For parent-child relationships where child records should be deleted with parent
- **SET NULL**: For optional relationships where child records should be preserved
- **RESTRICT**: For critical relationships where deletion should be prevented

## Indexing Strategy

### Primary Indexes

- All tables have a primary key (UUID) with automatic indexing

### Secondary Indexes

Strategic indexes are created for:

1. **Foreign Key Lookups**: All foreign key columns
2. **Common Query Patterns**: Frequently filtered columns
3. **Sorting Requirements**: Columns used in ORDER BY clauses
4. **Unique Constraints**: Columns with unique requirements

### Performance Considerations

- Composite indexes for multi-column queries
- Partial indexes for filtered data subsets
- Expression indexes for computed values
- Regular index maintenance and statistics updates

## Triggers

### Updated At Trigger

```sql
-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger for all tables with updated_at column
CREATE TRIGGER update_users_updated_at 
BEFORE UPDATE ON users 
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_accounts_updated_at 
BEFORE UPDATE ON accounts 
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_transfers_updated_at 
BEFORE UPDATE ON transfers 
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

### Audit Trail Trigger

```sql
-- Function to log changes to audit table
CREATE OR REPLACE FUNCTION audit_trigger()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO audit_logs (user_id, action, table_name, record_id, old_values)
        VALUES (NULL, TG_OP, TG_TABLE_NAME, OLD.id, row_to_json(OLD));
        RETURN OLD;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO audit_logs (user_id, action, table_name, record_id, old_values, new_values)
        VALUES (NULL, TG_OP, TG_TABLE_NAME, NEW.id, row_to_json(OLD), row_to_json(NEW));
        RETURN NEW;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO audit_logs (user_id, action, table_name, record_id, new_values)
        VALUES (NULL, TG_OP, TG_TABLE_NAME, NEW.id, row_to_json(NEW));
        RETURN NEW;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Audit triggers for sensitive tables
CREATE TRIGGER audit_users
AFTER INSERT OR UPDATE OR DELETE ON users
FOR EACH ROW EXECUTE FUNCTION audit_trigger();

CREATE TRIGGER audit_accounts
AFTER INSERT OR UPDATE OR DELETE ON accounts
FOR EACH ROW EXECUTE FUNCTION audit_trigger();

CREATE TRIGGER audit_transfers
AFTER INSERT OR UPDATE OR DELETE ON transfers
FOR EACH ROW EXECUTE FUNCTION audit_trigger();
```

## Database Migrations

### Migration Strategy

Database schema changes are managed through:

1. **Sequential Numbering**: Migration files numbered sequentially
2. **Idempotent Operations**: Migrations can be run multiple times safely
3. **Rollback Support**: Each migration has a corresponding rollback
4. **Version Tracking**: Migration version stored in database

### Migration File Structure

```
backend/internal/infrastructure/database/migrations/
├── 000001_create_users_table.up.sql
├── 000001_create_users_table.down.sql
├── 000002_create_accounts_table.up.sql
├── 000002_create_accounts_table.down.sql
└── ...
```

### Example Migration

```sql
-- 000001_create_users_table.up.sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20),
    date_of_birth DATE,
    country_code CHAR(2),
    kyc_status VARCHAR(20) DEFAULT 'pending',
    account_status VARCHAR(20) DEFAULT 'active',
    email_verified BOOLEAN DEFAULT FALSE,
    phone_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_kyc_status ON users(kyc_status);
CREATE INDEX idx_users_account_status ON users(account_status);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';
-- +goose StatementEnd

CREATE TRIGGER update_users_updated_at 
BEFORE UPDATE ON users 
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

## Backup and Recovery

### Backup Strategy

- **Daily Full Backups**: Complete database snapshots
- **Hourly Incremental Backups**: Transaction log backups
- **Encrypted Storage**: All backups encrypted at rest
- **Cross-Region Replication**: Backups stored in multiple regions

### Recovery Procedures

1. **Point-in-Time Recovery**: Restore to specific timestamp
2. **Full Database Restore**: Complete database restoration
3. **Table-Level Recovery**: Restore individual tables
4. **Automated Failover**: Automatic recovery in case of failure

## Performance Optimization

### Query Optimization

- **EXPLAIN ANALYZE**: Regular query performance analysis
- **Query Plans**: Optimized execution plans
- **Connection Pooling**: PgBouncer for connection management
- **Read Replicas**: Separate read-only database instances

### Monitoring

- **Slow Query Logs**: Identification of performance bottlenecks
- **Database Metrics**: CPU, memory, I/O monitoring
- **Connection Statistics**: Active connection tracking
- **Cache Hit Ratios**: Buffer pool efficiency monitoring

This database schema documentation provides a comprehensive overview of the Globepay database structure, ensuring data integrity, performance, and compliance with financial industry standards.