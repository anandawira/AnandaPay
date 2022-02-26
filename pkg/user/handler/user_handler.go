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
			"message": "Parameter validation failed.",
		})
		return
	}

	ctx := c.Request.Context()
	err = h.userUsecase.Register(ctx, reqBody.Fullname, reqBody.Email, reqBody.Password)
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
			"message": "Parameter validation failed.",
		})
		return
	}

	ctx := c.Request.Context()
	token, err := h.userUsecase.Login(ctx, reqBody.Email, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.SetCookie("access_token", token, 3600*12, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "User logged in successfully.",
		})
	}
}
