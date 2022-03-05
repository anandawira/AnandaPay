package usecase

import (
	"github.com/anandawira/anandapay/domain"
	"github.com/stretchr/testify/mock"
)

type MockWalletUsecase struct {
	mock.Mock
}

func (m *MockWalletUsecase) GetBalance(walletId string) (uint64, error) {
	args := m.Called(walletId)
	return uint64(args.Int(0)), args.Error(1)
}

func (m *MockWalletUsecase) TopUp(walletId string, amount uint32) (domain.Transaction, error) {
	args := m.Called(walletId, amount)
	return args.Get(0).(domain.Transaction), args.Error(1)
}
