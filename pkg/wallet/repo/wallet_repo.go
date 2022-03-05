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

		result = m.db.Model(&wallet).Where("id = ?", creditedWallet).Update("balance", gorm.Expr("balance + ?", amount))
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

func (m *walletRepository) Transaction(transactionId string, transactionTime time.Time, transactionType string, creditedWallet string, debitedWallet string, notes string, amount uint32) (domain.Transaction, error) {
	if transactionType != domain.TYPE_TRANSFER && transactionType != domain.TYPE_TOPUP {
		return domain.Transaction{}, domain.ErrInternalServerError
	}

	transaction := domain.Transaction{
		ID:              transactionId,
		TransactionTime: transactionTime,
		TransactionType: transactionType,
		CreditedWallet:  creditedWallet,
		DebitedWallet:   debitedWallet,
		Notes:           notes,
		Amount:          uint64(amount),
	}

	err := m.db.Transaction(func(tx *gorm.DB) error {
		// If transfer, we want to check debitedWallet's balance first.
		if transaction.TransactionType == domain.TYPE_TRANSFER {
			wallet := domain.Wallet{}
			result := m.db.Where("id = ?", transaction.DebitedWallet).Take(&wallet)
			if result.Error != nil {
				return domain.ErrWalletNotFound
			}
			if wallet.Balance < uint64(amount) {
				return domain.ErrInsufficientBalance
			}
		}

		// Add transaction
		result := m.db.Create(&transaction)
		if result.Error != nil {
			return domain.ErrInternalServerError
		}

		// Add balance to credited wallet
		result = m.db.Model(&domain.Wallet{}).Where("id = ?", creditedWallet).Update("balance", gorm.Expr("balance + ?", transaction.Amount))
		if result.Error != nil {
			return domain.ErrInternalServerError
		}

		if result.RowsAffected == 0 {
			return domain.ErrWalletNotFound
		}

		// If type == transfer, remove balance from debited wallet
		if transactionType == domain.TYPE_TRANSFER {
			result = m.db.Model(&domain.Wallet{}).Where("id = ?", debitedWallet).Update("balance", gorm.Expr("balance - ?", transaction.Amount))
			if result.Error != nil {
				return domain.ErrInternalServerError
			}

			if result.RowsAffected == 0 {
				return domain.ErrWalletNotFound
			}
		}

		return nil
	})

	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}
