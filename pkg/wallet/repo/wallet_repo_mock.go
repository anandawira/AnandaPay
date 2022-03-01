package repo

import "github.com/stretchr/testify/mock"

type MockWalletRepo struct {
	mock.Mock
}

func (m *MockWalletRepo) GetBalance(walletId string) (uint64, error) {
	args := m.Called(walletId)
	return uint64(args.Int(0)), args.Error(1)
}

func (m *MockWalletRepo) TopUp(walletId string, amount uint32) error {
	args := m.Called(walletId, amount)
	return args.Error(0)
}
