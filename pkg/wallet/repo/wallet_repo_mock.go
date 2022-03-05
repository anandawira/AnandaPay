package repo

import (
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/stretchr/testify/mock"
)

type MockWalletRepo struct {
	mock.Mock
}

func (m *MockWalletRepo) GetBalance(walletId string) (uint64, error) {
	args := m.Called(walletId)
	return uint64(args.Int(0)), args.Error(1)
}

func (m *MockWalletRepo) TopUp(transactionId string, transactionTime time.Time, creditedWallet string, notes string, amount uint32) error {
	args := m.Called(transactionId, transactionTime, creditedWallet, notes, amount)
	return args.Error(0)
}

func (m *MockWalletRepo) Transaction(transactionId string, transactionTime time.Time, transactionType string, creditedWallet string, debitedWallet string, notes string, amount uint32) (domain.Transaction, error) {
	panic("not implemented") // TODO: Implement
}
