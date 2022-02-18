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
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	userRepo := repo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, time.Second*5)
	handler.NewUserHandler(r, userUsecase)

	r.Run(":1234")
}
