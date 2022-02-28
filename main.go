package main

import (
	"net/http"

	"github.com/anandawira/anandapay/config"
	"github.com/anandawira/anandapay/pkg/user/handler"
	"github.com/anandawira/anandapay/pkg/user/middleware"
	"github.com/anandawira/anandapay/pkg/user/repo"
	"github.com/anandawira/anandapay/pkg/user/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.Connect()

	userRepo := repo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.AttachHandler(r, userUsecase)
	r.GET("/test-jwt", middleware.Authenticate, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Authenticated",
		})
	})

	r.Run(":1234")
}
