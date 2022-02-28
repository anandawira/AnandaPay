package middleware

import (
	"net/http"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/user/middleware/helper"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	const BEARER_PREFIX = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) <= len(BEARER_PREFIX) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": domain.ErrInvalidToken.Error(),
		})
		return
	}

	tokenString := authHeader[len(BEARER_PREFIX):]

	userId, walletId, err := helper.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": domain.ErrInvalidToken.Error(),
		})
		return
	}

	c.Set("userId", userId)
	c.Set("walletId", walletId)
	c.Next()
}
