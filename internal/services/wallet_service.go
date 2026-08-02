package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type WalletService struct {
	Repo *repositories.WalletRepository
}

func NewWalletService(repo *repositories.WalletRepository) *WalletService {
	return &WalletService{
		Repo: repo,
	}
}

func (s *WalletService) CreateWallet(
	ctx context.Context,
	userID uuid.UUID,
	currency string,
) (*models.Wallet, error) {

	wallet := &models.Wallet{
		ID:       uuid.New(),
		UserID:   userID,
		Currency: currency,
		Balance:  0,
		Status:   "ACTIVE",
	}

	err := s.Repo.CreateWallet(ctx, wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (s *WalletService) GetWallet(
	ctx context.Context,
	id uuid.UUID,
) (*models.Wallet, error) {

	return s.Repo.GetWalletByID(ctx, id)
}

func (s *WalletService) ListWallets(
	ctx context.Context,
) ([]models.Wallet, error) {

	return s.Repo.ListWallets(ctx)
}

func (s *WalletService) UpdateWallet(
	ctx context.Context,
	wallet *models.Wallet,
) error {

	return s.Repo.UpdateWallet(ctx, wallet)
}

func (s *WalletService) DeleteWallet(
	ctx context.Context,
	id uuid.UUID,
) error {

	return s.Repo.DeleteWallet(ctx, id)
}
