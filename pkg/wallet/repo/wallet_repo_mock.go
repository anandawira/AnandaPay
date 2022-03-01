package repo

import "github.com/stretchr/testify/mock"

type MockWalletRepo struct {
	mock.Mock
}

func (m *MockWalletRepo) GetBalance(walletId string) (uint64, error) {
	args := m.Called(walletId)
	return uint64(args.Int(0)), args.Error(1)
}
