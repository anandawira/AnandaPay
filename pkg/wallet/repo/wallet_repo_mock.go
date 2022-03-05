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

func (m *MockWalletRepo) Transaction(transactionId string, transactionTime time.Time, transactionType string, creditedWallet string, debitedWallet string, notes string, amount uint32) (domain.Transaction, error) {
	args := m.Called(transactionId, transactionTime, transactionType, creditedWallet, debitedWallet, notes, amount)
	return args.Get(0).(domain.Transaction), args.Error(1)
}
