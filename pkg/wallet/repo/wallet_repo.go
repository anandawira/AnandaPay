package repo

import (
	"github.com/anandawira/anandapay/domain"
	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) domain.WalletRepository {
	return &walletRepository{db: db}
}

func (m *walletRepository) GetBalance(walletId string) (int64, error) {
	wallet := domain.Wallet{}

	result := m.db.Select("balance").Where("id = ?", walletId).First(&wallet)
	if result.Error != nil {
		return int64(0), domain.ErrWalletNotFound
	}

	return wallet.Balance, nil
}
