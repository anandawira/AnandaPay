package main

import (
	"github.com/anandawira/anandapay/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Connect()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":1234")
}
