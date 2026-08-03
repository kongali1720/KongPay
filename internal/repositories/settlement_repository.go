package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
)

type SettlementRepository struct {
	DB *pgx.Conn
}

func NewSettlementRepository(
	db *pgx.Conn,
) *SettlementRepository {

	return &SettlementRepository{
		DB: db,
	}
}

func (r *SettlementRepository) Create(
	ctx context.Context,
	s *models.Settlement,
) error {

	query := `
	INSERT INTO settlements
	(
		id,
		batch_reference,
		currency,
		total_amount,
		transaction_count,
		status,
		created_at
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,NOW()
	)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		s.ID,
		s.BatchReference,
		s.Currency,
		s.TotalAmount,
		s.TransactionCount,
		s.Status,
	)

	return err
}

func (r *SettlementRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Settlement, error) {

	query := `
	SELECT
		id,
		batch_reference,
		currency,
		total_amount,
		transaction_count,
		status,
		created_at,
		completed_at
	FROM settlements
	WHERE id=$1
	`

	var s models.Settlement

	err := r.DB.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&s.ID,
		&s.BatchReference,
		&s.Currency,
		&s.TotalAmount,
		&s.TransactionCount,
		&s.Status,
		&s.CreatedAt,
		&s.CompletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func GenerateSettlementReference() string {
	return "SET-" + uuid.New().String()
}
