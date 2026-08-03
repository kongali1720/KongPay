package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
)

type SettlementEventRepository struct {
	DB *pgx.Conn
}

func NewSettlementEventRepository(
	db *pgx.Conn,
) *SettlementEventRepository {

	return &SettlementEventRepository{
		DB: db,
	}
}

func (r *SettlementEventRepository) Create(
	ctx context.Context,
	event *models.SettlementEvent,
) error {

	query := `
	INSERT INTO settlement_events
	(
		id,
		settlement_id,
		event_type,
		old_status,
		new_status,
		created_at
	)
	VALUES
	(
		$1,$2,$3,$4,$5,NOW()
	)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		event.ID,
		event.SettlementID,
		event.EventType,
		event.OldStatus,
		event.NewStatus,
	)

	return err
}

func (r *SettlementEventRepository) ListBySettlementID(
	ctx context.Context,
	settlementID interface{},
) ([]models.SettlementEvent, error) {

	query := `
	SELECT
		id,
		settlement_id,
		event_type,
		old_status,
		new_status,
		created_at
	FROM settlement_events
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

	var events []models.SettlementEvent

	for rows.Next() {

		var event models.SettlementEvent

		err := rows.Scan(
			&event.ID,
			&event.SettlementID,
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
