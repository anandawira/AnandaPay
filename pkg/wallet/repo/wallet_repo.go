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

func (m *walletRepository) GetBalance(walletId string) (uint64, error) {
	wallet := domain.Wallet{}

	result := m.db.Select("balance").Where("id = ?", walletId).Take(&wallet)
	if result.Error != nil {
		return uint64(0), domain.ErrWalletNotFound
	}

	return wallet.Balance, nil
}

func (m *walletRepository) TopUp(walletId string, amount uint32) error {
	wallet := domain.Wallet{}
	result := m.db.Where("id = ?", walletId).Take(&wallet)
	if result.Error != nil {
		return domain.ErrWalletNotFound
	}

	result = m.db.Model(&wallet).Update("balance", gorm.Expr("balance + ?", amount))
	if result.Error != nil {
		return domain.ErrInternalServerError
	}

	return nil
}
