package handler

import (
	"net/http"
	"testing"
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/helper"
	"github.com/anandawira/anandapay/pkg/wallet/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WalletHandlerTestSuite struct {
	suite.Suite
	MockWalletUsecase *usecase.MockWalletUsecase
	handler           *WalletHandler
}

func (ts *WalletHandlerTestSuite) SetupSuite() {
	ts.MockWalletUsecase = new(usecase.MockWalletUsecase)
	ts.handler = &WalletHandler{
		walletUsecase: ts.MockWalletUsecase,
	}
	gin.SetMode(gin.TestMode)
}

func (ts *WalletHandlerTestSuite) TestGetBalance() {
	ts.T().Run("It should return with StatusOK and balance on wallet found", func(t *testing.T) {
		ts.MockWalletUsecase.On(
			"GetBalance",
			mock.AnythingOfType("string"),
		).Return(12, nil).Once()

		c, rec := helper.CreateGetContext()

		ts.handler.BalanceGet(c)

		helper.AssertResponse(t, http.StatusOK, gin.H{
			"message": "Wallet balance retrieved successfully.",
			"data": BalanceResponseData{
				Balance: 12,
			},
		}, rec)
	})

	ts.T().Run("It should return with StatusNotFound on wallet not found", func(t *testing.T) {
		ts.MockWalletUsecase.On(
			"GetBalance",
			mock.AnythingOfType("string"),
		).Return(0, domain.ErrWalletNotFound).Once()

		c, rec := helper.CreateGetContext()

		ts.handler.BalanceGet(c)

		helper.AssertResponse(t, http.StatusNotFound, gin.H{"message": domain.ErrWalletNotFound.Error()}, rec)
	})
}

func (ts *WalletHandlerTestSuite) TestTopUp() {
	ts.T().Run("It should return with StatusOK on success top up", func(t *testing.T) {
		transaction := domain.Transaction{
			ID:              "id",
			TransactionTime: time.Now(),
			TransactionType: domain.TYPE_TOPUP,
			CreditedWallet:  "credited wallet",
			DebitedWallet:   "",
			Notes:           "topup",
			Amount:          100000,
		}
		ts.MockWalletUsecase.On(
			"TopUp",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("uint32"),
		).Return(transaction, nil)

		body := map[string]string{
			"amount": "5000000",
		}

		c, rec := helper.CreatePostContext(body)

		ts.handler.TopUpPost(c)
		helper.AssertResponse(t, http.StatusOK, gin.H{
			"message": "Wallet balance top up success.",
			"data": TopupResponseData{
				Transaction: transaction,
			},
		}, rec)
	})

	ts.T().Run("It should return error on bad request body", func(t *testing.T) {
		body := map[string]string{
			"amount": "dddddd",
		}
		c, rec := helper.CreatePostContext(body)

		ts.handler.TopUpPost(c)
		helper.AssertResponse(t, http.StatusBadRequest, gin.H{
			"message": domain.ErrParameterValidation.Error(),
		}, rec)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(WalletHandlerTestSuite))
}
