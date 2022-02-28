package usecase

import "github.com/stretchr/testify/mock"

type MockWalletUsecase struct {
	mock.Mock
}

func (m *MockWalletUsecase) GetBalance(walletId string) (int64, error) {
	args := m.Called(walletId)
	return int64(args.Int(0)), args.Error(1)
}
