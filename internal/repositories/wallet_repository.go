package repositories

import (
	"context"

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
