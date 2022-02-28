package main

import (
	"net/http"

	"github.com/anandawira/anandapay/config"
	userHandler "github.com/anandawira/anandapay/pkg/user/handler"
	"github.com/anandawira/anandapay/pkg/user/middleware"
	walletHandler "github.com/anandawira/anandapay/pkg/wallet/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.Connect()

	userHandler.AttachHandler(r, db)
	walletHandler.AttachHandler(r, db)

	r.GET("/test-jwt", middleware.Authenticate, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Authenticated",
		})
	})

	r.Run(":1234")
}
