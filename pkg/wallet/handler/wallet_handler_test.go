package handler

import (
	"net/http"
	"testing"

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

func TestSuite(t *testing.T) {
	suite.Run(t, new(WalletHandlerTestSuite))
}
