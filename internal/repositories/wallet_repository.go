package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
)

type WalletRepository struct {
	DB *pgx.Conn
}

func NewWalletRepository(db *pgx.Conn) *WalletRepository {
	return &WalletRepository{
		DB: db,
	}
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {

	query := `
	INSERT INTO wallets
	(id, user_id, currency, balance, status, created_at, updated_at)
	VALUES
	($1,$2,$3,$4,$5,NOW(),NOW())
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		wallet.ID,
		wallet.UserID,
		wallet.Currency,
		wallet.Balance,
		wallet.Status,
	)

	return err
}

func (r *WalletRepository) GetWalletByID(ctx context.Context, id uuid.UUID) (*models.Wallet, error) {

	query := `
	SELECT
		id,
		user_id,
		balance,
		currency,
		status,
		created_at,
		updated_at
	FROM wallets
	WHERE id = $1
	`

	var wallet models.Wallet

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Balance,
		&wallet.Currency,
		&wallet.Status,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *WalletRepository) ListWallets(ctx context.Context) ([]models.Wallet, error) {

	query := `
	SELECT
		id,
		user_id,
		balance,
		currency,
		status,
		created_at,
		updated_at
	FROM wallets
	ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []models.Wallet

	for rows.Next() {

		var wallet models.Wallet

		err := rows.Scan(
			&wallet.ID,
			&wallet.UserID,
			&wallet.Balance,
			&wallet.Currency,
			&wallet.Status,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

func (r *WalletRepository) UpdateWallet(ctx context.Context, wallet *models.Wallet) error {

	query := `
	UPDATE wallets
	SET
		currency = $2,
		balance = $3,
		status = $4,
		updated_at = NOW()
	WHERE id = $1
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		wallet.ID,
		wallet.Currency,
		wallet.Balance,
		wallet.Status,
	)

	return err
}

func (r *WalletRepository) DeleteWallet(ctx context.Context, id uuid.UUID) error {

	query := `
	DELETE FROM wallets
	WHERE id = $1
	`

	_, err := r.DB.Exec(ctx, query, id)

	return err
}
