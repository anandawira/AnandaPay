package handler

import (
	"net/http"

	"github.com/anandawira/anandapay/pkg/model"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase model.UserUsecase
}

func NewUserHandler(g *gin.Engine, usecase model.UserUsecase) {
	handler := &userHandler{
		userUsecase: usecase,
	}
	g.POST("/users", handler.RegisterPost)
}

func (h *userHandler) RegisterPost(c *gin.Context) {
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
			"message": "User has been registered to the database.",
		})
	}
}
