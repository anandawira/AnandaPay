package handler

import (
	"net/http"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/user/middleware"
	"github.com/anandawira/anandapay/pkg/wallet/repo"
	"github.com/anandawira/anandapay/pkg/wallet/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WalletHandler struct {
	walletUsecase domain.WalletUsecase
}

func AttachHandler(g *gin.Engine, db *gorm.DB) {
	repo := repo.NewWalletRepository(db)
	usecase := usecase.NewWalletUsecase(repo)
	handler := &WalletHandler{
		walletUsecase: usecase,
	}

	wallet := g.Group("/wallet", middleware.Authenticate)
	wallet.GET("/balance", handler.BalanceGet)
	wallet.POST("/topup", handler.TopUpPost)
	wallet.POST("/transfer", handler.TransferPost)
}

func (h *WalletHandler) BalanceGet(c *gin.Context) {
	walletId := c.GetString("walletId")

	balance, err := h.walletUsecase.GetBalance(walletId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": domain.ErrWalletNotFound.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet balance retrieved successfully.",
		"data": BalanceResponseData{
			WalletID: walletId,
			Balance:  balance,
		},
	})
}

func (h *WalletHandler) TopUpPost(c *gin.Context) {
	reqBody := TopupRequestData{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrParameterValidation.Error(),
		})
		return
	}

	walletId := c.GetString("walletId")
	amount := reqBody.Amount

	transaction, err := h.walletUsecase.TopUp(walletId, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet balance top up success.",
		"data": TopupResponseData{
			Transaction: transaction,
		},
	})
}

func (h *WalletHandler) TransferPost(c *gin.Context) {
	reqBody := TransferRequestData{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrParameterValidation.Error(),
		})
		return
	}

	walletId := c.GetString("walletId")

	transaction, err := h.walletUsecase.Transfer(walletId, reqBody.ReceiverID, reqBody.Notes, reqBody.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer success.",
		"data": TopupResponseData{
			Transaction: transaction,
		},
	})
}
