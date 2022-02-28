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
	panic("not implemented") // TODO: Implement
}


