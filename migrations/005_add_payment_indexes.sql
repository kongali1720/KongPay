CREATE INDEX IF NOT EXISTS idx_transactions_sender_wallet
ON transactions(sender_wallet_id);

CREATE INDEX IF NOT EXISTS idx_transactions_receiver_wallet
ON transactions(receiver_wallet_id);

CREATE INDEX IF NOT EXISTS idx_ledger_wallet
ON ledger_entries(wallet_id);

CREATE INDEX IF NOT EXISTS idx_ledger_transaction
ON ledger_entries(transaction_id);
