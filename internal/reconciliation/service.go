package reconciliation

import (
	"context"

	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type Service struct {
	Repo *repositories.SettlementReconciliationRepository

	EventRepo *repositories.ReconciliationEventRepository
}

func NewService(
	repo *repositories.SettlementReconciliationRepository,
	eventRepo *repositories.ReconciliationEventRepository,
) *Service {

	return &Service{

		Repo: repo,

		EventRepo: eventRepo,
	}
}

func (s *Service) Reconcile(
	ctx context.Context,
	settlementID uuid.UUID,
	expected float64,
	actual float64,
) (*models.SettlementReconciliation, error) {

	resultID := uuid.New()

	// Initial reconciliation status

	status := StatusMatched

	if expected != actual {

		status = StatusMismatch

	}

	result := &models.SettlementReconciliation{

		ID: resultID,

		SettlementID: settlementID,

		ExpectedAmount: expected,

		ActualAmount: actual,

		Difference: expected - actual,

		Status: status,
	}

	// Save reconciliation result

	err := s.Repo.Create(
		ctx,
		result,
	)

	if err != nil {
		return nil, err
	}

	// Create audit event

	eventType := EventMatched

	if status == StatusMismatch {

		eventType = EventMismatch

	}

	err = s.EventRepo.Create(
		ctx,
		&models.ReconciliationEvent{

			ID: uuid.New(),

			SettlementID: settlementID,

			ReconciliationID: result.ID,

			EventType: eventType,

			OldStatus: "CREATED",

			NewStatus: status,
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
