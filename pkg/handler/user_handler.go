package handler

import (
	"net/http"

	"github.com/anandawira/anandapay/pkg/model"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase model.UserUsecase
}

func AttachHandler(g *gin.Engine, usecase model.UserUsecase) {
	handler := &UserHandler{
		userUsecase: usecase,
	}
	g.POST("/users", handler.RegisterPost)
}

func (h *UserHandler) RegisterPost(c *gin.Context) {
	reqBody := RegisterRequest{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
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
