package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
)

type LedgerRepository struct{}

func NewLedgerRepository() *LedgerRepository {
	return &LedgerRepository{}
}

func (r *LedgerRepository) Create(
	ctx context.Context,
	tx pgx.Tx,
	entry *models.LedgerEntry,
) error {

	query := `
	INSERT INTO ledger_entries
	(
		id,
		transaction_id,
		wallet_id,
		entry_type,
		amount,
		balance_before,
		balance_after,
		created_at
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7,NOW()
	)
	`

	_, err := tx.Exec(
		ctx,
		query,
		entry.ID,
		entry.TransactionID,
		entry.WalletID,
		entry.EntryType,
		entry.Amount,
		entry.BalanceBefore,
		entry.BalanceAfter,
	)

	return err
}

func (r *LedgerRepository) ListByWalletID(
	ctx context.Context,
	db *pgx.Conn,
	walletID uuid.UUID,
) ([]models.LedgerEntry, error) {

	query := `
	SELECT
		id,
		transaction_id,
		wallet_id,
		entry_type,
		amount,
		balance_before,
		balance_after,
		created_at
	FROM ledger_entries
	WHERE wallet_id = $1
	ORDER BY created_at DESC
	`

	rows, err := db.Query(ctx, query, walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.LedgerEntry

	for rows.Next() {

		var entry models.LedgerEntry

		err := rows.Scan(
			&entry.ID,
			&entry.TransactionID,
			&entry.WalletID,
			&entry.EntryType,
			&entry.Amount,
			&entry.BalanceBefore,
			&entry.BalanceAfter,
			&entry.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
