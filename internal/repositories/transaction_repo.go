package repositories

import (
    "context"
    "encoding/json"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Transaction struct {
    ID            string                 `json:"id"`
    TransactionID string                 `json:"transaction_id"`
    ProviderTxID  string                 `json:"provider_tx_id"`
    Amount        float64                `json:"amount"`
    Currency      string                 `json:"currency"`
    Method        string                 `json:"method"`
    Status        string                 `json:"status"`
    CustomerID    string                 `json:"customer_id"`
    MerchantID    string                 `json:"merchant_id"`
    RedirectURL   string                 `json:"redirect_url"`
    QRCode        string                 `json:"qr_code"`
    VirtualAccount string                `json:"virtual_account"`
    Metadata      map[string]interface{} `json:"metadata"`
    CreatedAt     time.Time              `json:"created_at"`
    UpdatedAt     time.Time              `json:"updated_at"`
    SettledAt     *time.Time             `json:"settled_at"`
}

type TransactionRepository struct {
    db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
    return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, tx *Transaction) error {
    query := `
        INSERT INTO transactions (
            transaction_id, provider_tx_id, amount, currency, method, status,
            customer_id, merchant_id, redirect_url, qr_code, virtual_account, metadata
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    `

    metadataJSON, err := json.Marshal(tx.Metadata)
    if err != nil {
        return err
    }

    _, err = r.db.Exec(ctx, query,
        tx.TransactionID, tx.ProviderTxID, tx.Amount, tx.Currency, tx.Method,
        tx.Status, tx.CustomerID, tx.MerchantID, tx.RedirectURL, tx.QRCode,
        tx.VirtualAccount, metadataJSON,
    )
    return err
}

func (r *TransactionRepository) GetByID(ctx context.Context, transactionID string) (*Transaction, error) {
    query := `
        SELECT id, transaction_id, provider_tx_id, amount, currency, method, status,
               customer_id, merchant_id, redirect_url, qr_code, virtual_account, metadata,
               created_at, updated_at, settled_at
        FROM transactions WHERE transaction_id = $1
    `

    var tx Transaction
    var metadataJSON []byte
    err := r.db.QueryRow(ctx, query, transactionID).Scan(
        &tx.ID, &tx.TransactionID, &tx.ProviderTxID, &tx.Amount, &tx.Currency, &tx.Method,
        &tx.Status, &tx.CustomerID, &tx.MerchantID, &tx.RedirectURL, &tx.QRCode,
        &tx.VirtualAccount, &metadataJSON, &tx.CreatedAt, &tx.UpdatedAt, &tx.SettledAt,
    )
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(metadataJSON, &tx.Metadata); err != nil {
        tx.Metadata = make(map[string]interface{})
    }

    return &tx, nil
}

func (r *TransactionRepository) UpdateStatus(ctx context.Context, transactionID, status string) error {
    query := `UPDATE transactions SET status = $1, updated_at = NOW() WHERE transaction_id = $2`
    _, err := r.db.Exec(ctx, query, status, transactionID)
    return err
}

func (r *TransactionRepository) UpdateSettled(ctx context.Context, transactionID string) error {
    query := `UPDATE transactions SET status = 'SETTLED', settled_at = NOW(), updated_at = NOW() WHERE transaction_id = $1`
    _, err := r.db.Exec(ctx, query, transactionID)
    return err
}
