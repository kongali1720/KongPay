package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
)

type ReconciliationEventRepository struct {
	DB *pgx.Conn
}

func NewReconciliationEventRepository(
	db *pgx.Conn,
) *ReconciliationEventRepository {

	return &ReconciliationEventRepository{
		DB: db,
	}
}

func (r *ReconciliationEventRepository) Create(
	ctx context.Context,
	event *models.ReconciliationEvent,
) error {

	query := `
	INSERT INTO reconciliation_events
	(
		id,
		settlement_id,
		reconciliation_id,
		event_type,
		old_status,
		new_status,
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
		event.ID,
		event.SettlementID,
		event.ReconciliationID,
		event.EventType,
		event.OldStatus,
		event.NewStatus,
	)

	return err
}

func (r *ReconciliationEventRepository) ListBySettlementID(
	ctx context.Context,
	settlementID interface{},
) ([]models.ReconciliationEvent, error) {

	query := `
	SELECT
		id,
		settlement_id,
		reconciliation_id,
		event_type,
		old_status,
		new_status,
		created_at
	FROM reconciliation_events
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

	var events []models.ReconciliationEvent

	for rows.Next() {

		var event models.ReconciliationEvent

		err := rows.Scan(
			&event.ID,
			&event.SettlementID,
			&event.ReconciliationID,
			&event.EventType,
			&event.OldStatus,
			&event.NewStatus,
			&event.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
