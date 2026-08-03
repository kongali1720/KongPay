CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY,
    reference_no VARCHAR(50) UNIQUE NOT NULL,

    sender_wallet_id UUID NOT NULL,
    receiver_wallet_id UUID NOT NULL,

    amount NUMERIC(20,2) NOT NULL,
    currency VARCHAR(10) NOT NULL,

    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_sender_wallet
        FOREIGN KEY(sender_wallet_id)
        REFERENCES wallets(id)
        ON DELETE RESTRICT,

    CONSTRAINT fk_receiver_wallet
        FOREIGN KEY(receiver_wallet_id)
        REFERENCES wallets(id)
        ON DELETE RESTRICT
);
