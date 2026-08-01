CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS wallets (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    balance NUMERIC(20,2) NOT NULL DEFAULT 0,

    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',

    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',

    created_at TIMESTAMP DEFAULT NOW(),

    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_wallet_user
ON wallets(user_id);
