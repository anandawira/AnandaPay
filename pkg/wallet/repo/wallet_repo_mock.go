package repo

import "github.com/stretchr/testify/mock"

type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) GetBalance(walletId string) (int64, error) {
	args := m.Called(walletId)
	return int64(args.Int(0)), args.Error(1)
}