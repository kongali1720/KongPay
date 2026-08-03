CREATE TABLE IF NOT EXISTS ledger_entries (
    id UUID PRIMARY KEY,

    transaction_id UUID NOT NULL,
    wallet_id UUID NOT NULL,

    entry_type VARCHAR(10) NOT NULL,

    amount NUMERIC(20,2) NOT NULL,

    balance_before NUMERIC(20,2) NOT NULL,
    balance_after NUMERIC(20,2) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_transaction
        FOREIGN KEY(transaction_id)
        REFERENCES transactions(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_wallet
        FOREIGN KEY(wallet_id)
        REFERENCES wallets(id)
        ON DELETE CASCADE
);
