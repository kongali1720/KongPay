ALTER TABLE transactions
ADD COLUMN IF NOT EXISTS idempotency_key VARCHAR(100);

CREATE UNIQUE INDEX IF NOT EXISTS idx_transactions_idempotency
ON transactions(idempotency_key);
