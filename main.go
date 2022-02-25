package main

import (
	"time"

	"github.com/anandawira/anandapay/config"
	"github.com/anandawira/anandapay/pkg/handler"
	"github.com/anandawira/anandapay/pkg/repo"
	"github.com/anandawira/anandapay/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.Connect()

	userRepo := repo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, time.Second*5)
	handler.AttachHandler(r, userUsecase)

	r.Run(":1234")
}
