package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
)

type SettlementReconciliationRepository struct {
	DB *pgx.Conn
}

func NewSettlementReconciliationRepository(
	db *pgx.Conn,
) *SettlementReconciliationRepository {

	return &SettlementReconciliationRepository{
		DB: db,
	}
}

func (r *SettlementReconciliationRepository) Create(
	ctx context.Context,
	item *models.SettlementReconciliation,
) error {

	query := `
	INSERT INTO settlement_reconciliations
	(
		id,
		settlement_id,
		transaction_id,
		expected_amount,
		actual_amount,
		difference,
		status,
		notes,
		created_at,
		updated_at
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7,$8,NOW(),NOW()
	)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		item.ID,
		item.SettlementID,
		item.TransactionID,
		item.ExpectedAmount,
		item.ActualAmount,
		item.Difference,
		item.Status,
		item.Notes,
	)

	return err
}

func (r *SettlementReconciliationRepository) ListBySettlementID(
	ctx context.Context,
	settlementID interface{},
) ([]models.SettlementReconciliation, error) {

	query := `
	SELECT
		id,
		settlement_id,
		transaction_id,
		expected_amount,
		actual_amount,
		difference,
		status,
		notes,
		created_at,
		updated_at
	FROM settlement_reconciliations
	WHERE settlement_id=$1
	ORDER BY created_at ASC
	`

	rows, err := r.DB.Query(
		ctx,
		query,
		settlementID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []models.SettlementReconciliation

	for rows.Next() {

		var item models.SettlementReconciliation

		err := rows.Scan(
			&item.ID,
			&item.SettlementID,
			&item.TransactionID,
			&item.ExpectedAmount,
			&item.ActualAmount,
			&item.Difference,
			&item.Status,
			&item.Notes,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	return results, nil
}
