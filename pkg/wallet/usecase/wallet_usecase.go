package usecase

import (
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/google/uuid"
)

type walletUsecase struct {
	walletRepo domain.WalletRepository
}

func NewWalletUsecase(repo domain.WalletRepository) domain.WalletUsecase {
	return &walletUsecase{walletRepo: repo}
}

func (m *walletUsecase) GetBalance(walletId string) (uint64, error) {
	return m.walletRepo.GetBalance(walletId)
}

func (m *walletUsecase) TopUp(walletId string, amount uint32) error {
	return m.walletRepo.TopUp(
		uuid.NewString(),
		time.Now(),
		walletId,
		"Free top up",
		amount,
	)
}
