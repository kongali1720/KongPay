package repositories
import ("context"; "encoding/json"; "github.com/jackc/pgx/v5/pgxpool"; "github.com/kongali1720/KongPay/internal/models")
type TransactionRepository struct { db *pgxpool.Pool }
func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository { return &TransactionRepository{db: db} }
func (r *TransactionRepository) Create(ctx context.Context, tx *models.Transaction) error {
    metadataJSON, _ := json.Marshal(tx.Metadata)
    _, err := r.db.Exec(ctx, `INSERT INTO transactions (transaction_id, provider_tx_id, amount, currency, method, status, customer_id, merchant_id, redirect_url, qr_code, virtual_account, metadata) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
        tx.TransactionID, tx.ProviderTxID, tx.Amount, tx.Currency, tx.Method, tx.Status, tx.CustomerID, tx.MerchantID, tx.RedirectURL, tx.QRCode, tx.VirtualAccount, metadataJSON)
    return err
}
func (r *TransactionRepository) GetByID(ctx context.Context, transactionID string) (*models.Transaction, error) {
    var tx models.Transaction; var metadataJSON []byte
    err := r.db.QueryRow(ctx, `SELECT id, transaction_id, provider_tx_id, amount, currency, method, status, customer_id, merchant_id, redirect_url, qr_code, virtual_account, metadata, created_at, updated_at, settled_at FROM transactions WHERE transaction_id = $1`, transactionID).Scan(
        &tx.ID, &tx.TransactionID, &tx.ProviderTxID, &tx.Amount, &tx.Currency, &tx.Method, &tx.Status, &tx.CustomerID, &tx.MerchantID, &tx.RedirectURL, &tx.QRCode, &tx.VirtualAccount, &metadataJSON, &tx.CreatedAt, &tx.UpdatedAt, &tx.SettledAt)
    if err != nil { return nil, err }
    json.Unmarshal(metadataJSON, &tx.Metadata)
    return &tx, nil
}
func (r *TransactionRepository) UpdateStatus(ctx context.Context, transactionID, status string) error {
    _, err := r.db.Exec(ctx, `UPDATE transactions SET status = $1, updated_at = NOW() WHERE transaction_id = $2`, status, transactionID)
    return err
}
func (r *TransactionRepository) UpdateSettled(ctx context.Context, transactionID string) error {
    _, err := r.db.Exec(ctx, `UPDATE transactions SET status = 'SETTLED', settled_at = NOW(), updated_at = NOW() WHERE transaction_id = $1`, transactionID)
    return err
}
