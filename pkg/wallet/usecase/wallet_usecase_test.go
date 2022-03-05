package usecase

import (
	"testing"
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/wallet/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type WalletUsecaseTestSuite struct {
	suite.Suite

	mockRepo *repo.MockWalletRepo
	usecase  domain.WalletUsecase
}

func (ts *WalletUsecaseTestSuite) SetupSuite() {
	ts.mockRepo = new(repo.MockWalletRepo)
	ts.usecase = NewWalletUsecase(ts.mockRepo)
}

func (ts *WalletUsecaseTestSuite) TestGetBalance() {
	ts.T().Run("It should return balance if wallet found", func(t *testing.T) {
		ts.mockRepo.On(
			"GetBalance",
			mock.AnythingOfType("string"),
		).Return(12, nil).Once()

		balance, err := ts.usecase.GetBalance("walletId1")
		assert.NoError(t, err)
		assert.Equal(t, uint64(12), balance)
	})

	ts.T().Run("It should return error if wallet not found", func(t *testing.T) {
		ts.mockRepo.On(
			"GetBalance",
			mock.AnythingOfType("string"),
		).Return(0, domain.ErrWalletNotFound).Once()

		_, err := ts.usecase.GetBalance("walletId1")
		assert.Error(t, err)
	})
}

func (ts *WalletUsecaseTestSuite) TestTopUp() {
	const WALLET_ID = "wallet id"
	const AMOUNT = 500000
	ts.T().Run("It should return no error on wallet found", func(t *testing.T) {
		want := domain.Transaction{
			ID:              WALLET_ID,
			TransactionTime: time.Now(),
			TransactionType: domain.TYPE_TOPUP,
			CreditedWallet:  WALLET_ID,
			DebitedWallet:   "",
			Notes:           "Free topup",
			Amount:          AMOUNT,
		}
		ts.mockRepo.On(
			"Transaction",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("Time"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("uint32"),
		).Return(want, nil).Once()

		got, err := ts.usecase.TopUp("wallet id", 500000)
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	ts.T().Run("It should return error on wallet not found", func(t *testing.T) {
		ts.mockRepo.On(
			"Transaction",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("Time"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("uint32"),
		).Return(domain.Transaction{}, domain.ErrWalletNotFound).Once()

		_, err := ts.usecase.TopUp("wallet id", 500000)
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(WalletUsecaseTestSuite))
}
