package repositories

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"

    "github.com/kongali1720/KongPay/internal/models"
)

type TransactionRepository struct {
    db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
    return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, tx *models.Transaction) error {
    query := `
        INSERT INTO transactions (
            id, transaction_id, provider_tx_id, amount, currency, method, status,
            customer_id, merchant_id, redirect_url, qr_code, virtual_account, metadata,
            created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
    `

    metadataJSON, err := json.Marshal(tx.Metadata)
    if err != nil {
        return err
    }

    _, err = r.db.ExecContext(ctx, query,
        tx.ID, tx.TransactionID, tx.ProviderTxID, tx.Amount, tx.Currency, tx.Method,
        tx.Status, tx.CustomerID, tx.MerchantID, tx.RedirectURL, tx.QRCode,
        tx.VirtualAccount, metadataJSON, tx.CreatedAt, tx.UpdatedAt,
    )
    return err
}

func (r *TransactionRepository) GetByID(ctx context.Context, transactionID string) (*models.Transaction, error) {
    query := `
        SELECT id, transaction_id, provider_tx_id, amount, currency, method, status,
               customer_id, merchant_id, redirect_url, qr_code, virtual_account, metadata,
               created_at, updated_at, settled_at
        FROM transactions WHERE transaction_id = $1
    `

    var tx models.Transaction
    var metadataJSON []byte
    var settledAt sql.NullTime

    err := r.db.QueryRowContext(ctx, query, transactionID).Scan(
        &tx.ID, &tx.TransactionID, &tx.ProviderTxID, &tx.Amount, &tx.Currency, &tx.Method,
        &tx.Status, &tx.CustomerID, &tx.MerchantID, &tx.RedirectURL, &tx.QRCode,
        &tx.VirtualAccount, &metadataJSON, &tx.CreatedAt, &tx.UpdatedAt, &settledAt,
    )
    if err != nil {
        return nil, err
    }

    if settledAt.Valid {
        tx.SettledAt = &settledAt.Time
    }

    if err := json.Unmarshal(metadataJSON, &tx.Metadata); err != nil {
        tx.Metadata = make(map[string]interface{})
    }

    return &tx, nil
}

func (r *TransactionRepository) UpdateStatus(ctx context.Context, transactionID, status string) error {
    query := `UPDATE transactions SET status = $1, updated_at = NOW() WHERE transaction_id = $2`
    result, err := r.db.ExecContext(ctx, query, status, transactionID)
    if err != nil {
        return err
    }
    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("transaction not found: %s", transactionID)
    }
    return nil
}

func (r *TransactionRepository) UpdateSettled(ctx context.Context, transactionID string) error {
    query := `UPDATE transactions SET status = 'SETTLED', settled_at = NOW(), updated_at = NOW() WHERE transaction_id = $1`
    _, err := r.db.ExecContext(ctx, query, transactionID)
    return err
}
