CREATE TABLE IF NOT EXISTS settlements (

    id UUID PRIMARY KEY,

    batch_reference VARCHAR(100) NOT NULL UNIQUE,

    currency VARCHAR(10) NOT NULL,

    total_amount NUMERIC(20,2) NOT NULL,

    transaction_count INTEGER NOT NULL,

    status VARCHAR(30) NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),

    completed_at TIMESTAMP
);
