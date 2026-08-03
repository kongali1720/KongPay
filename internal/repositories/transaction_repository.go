package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
)

type TransactionRepository struct {
	DB *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (r *TransactionRepository) Create(
	ctx context.Context,
	tx pgx.Tx,
	t *models.Transaction,
) error {

	query := `
	INSERT INTO transactions
	(
		id,
		reference_no,
		sender_wallet_id,
		receiver_wallet_id,
		amount,
		currency,
		status,
		created_at,
		updated_at
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7,NOW(),NOW()
	)
	`

	_, err := tx.Exec(
		ctx,
		query,
		t.ID,
		t.ReferenceNo,
		t.SenderWalletID,
		t.ReceiverWalletID,
		t.Amount,
		t.Currency,
		t.Status,
	)

	return err
}

func (r *TransactionRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Transaction, error) {

	query := `
	SELECT
		id,
		reference_no,
		sender_wallet_id,
		receiver_wallet_id,
		amount,
		currency,
		status,
		created_at,
		updated_at
	FROM transactions
	WHERE id=$1
	`

	var t models.Transaction

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&t.ID,
		&t.ReferenceNo,
		&t.SenderWalletID,
		&t.ReceiverWalletID,
		&t.Amount,
		&t.Currency,
		&t.Status,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
