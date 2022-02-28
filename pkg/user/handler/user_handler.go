package handler

import (
	"net/http"

	"github.com/anandawira/anandapay/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func AttachHandler(g *gin.Engine, usecase domain.UserUsecase) {
	handler := &UserHandler{
		userUsecase: usecase,
	}
	g.POST("/register", handler.RegisterPost)
	g.POST("/login", handler.LoginPost)
}

func (h *UserHandler) RegisterPost(c *gin.Context) {
	reqBody := RegisterRequest{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrParameterValidation.Error(),
		})
		return
	}

	err = h.userUsecase.Register(reqBody.Fullname, reqBody.Email, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "User registered to the database successfully.",
		})
	}
}

func (h *UserHandler) LoginPost(c *gin.Context) {
	reqBody := LoginRequest{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrParameterValidation.Error(),
		})
		return
	}

	user, wallet, token, err := h.userUsecase.Login(reqBody.Email, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "User logged in successfully.",
			"data": LoginResponseData{
				UserID:      user.ID,
				WalletID:    wallet.ID,
				Fullname:    user.FullName,
				Email:       user.Email,
				AccessToken: token,
			},
		})
	}
}
