package repo

import (
	"time"

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

func (m *walletRepository) TopUp(transactionId string, transactionTime time.Time, creditedWallet string, notes string, amount uint32) error {
	wallet := domain.Wallet{}
	result := m.db.Where("id = ?", creditedWallet).Take(&wallet)
	if result.Error != nil {
		return domain.ErrWalletNotFound
	}

	transaction := domain.Transaction{
		ID:              transactionId,
		TransactionTime: transactionTime,
		CreditedWallet:  creditedWallet,
		Notes:           notes,
		Amount:          uint64(amount),
	}

	err := m.db.Transaction(func(tx *gorm.DB) error {
		result := m.db.Create(&transaction)
		if result.Error != nil {
			return result.Error
		}

		result = m.db.Model(&wallet).Update("balance", gorm.Expr("balance + ?", amount))
		if result.Error != nil {
			return domain.ErrInternalServerError
		}

		if result.RowsAffected == 0 {
			return domain.ErrInternalServerError
		}

		return nil
	})

	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}
