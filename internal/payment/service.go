package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/repositories"
)

type Service struct {
	DB         *pgx.Conn
	WalletRepo *repositories.WalletRepository
	TxRepo     *repositories.TransactionRepository
	LedgerRepo *repositories.LedgerRepository
}

func NewService(
	db *pgx.Conn,
	wallet *repositories.WalletRepository,
	txRepo *repositories.TransactionRepository,
	ledger *repositories.LedgerRepository,
) *Service {

	return &Service{
		DB:         db,
		WalletRepo: wallet,
		TxRepo:     txRepo,
		LedgerRepo: ledger,
	}
}

func (s *Service) Transfer(ctx context.Context, req TransferRequest) error {

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Logic akan kita isi pada step berikutnya

	return tx.Commit(ctx)
}

func GenerateReference() string {
	return "KP-" + uuid.New().String()
}
