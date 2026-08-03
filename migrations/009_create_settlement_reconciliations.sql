CREATE TABLE IF NOT EXISTS settlement_reconciliations (

    id UUID PRIMARY KEY,

    settlement_id UUID NOT NULL,

    transaction_id UUID,

    expected_amount NUMERIC(20,2) NOT NULL,

    actual_amount NUMERIC(20,2) NOT NULL,

    difference NUMERIC(20,2) NOT NULL,

    status VARCHAR(30) NOT NULL,

    notes TEXT,

    created_at TIMESTAMP DEFAULT NOW(),

    updated_at TIMESTAMP DEFAULT NOW()
);
