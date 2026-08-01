package repositories

import "github.com/kongali1720/KongPay/internal/models"

type WalletRepository interface {

	Create(wallet *models.Wallet) error

	GetByID(id string) (*models.Wallet, error)

	Update(wallet *models.Wallet) error

	Delete(id string) error

	List() ([]models.Wallet, error)
}
