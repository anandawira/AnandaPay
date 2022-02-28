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
