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
			"message": domain.ErrInvalidToken,
		})
		return
	}

	tokenString := authHeader[len(BEARER_PREFIX):]

	id, err := helper.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err,
		})
		return
	}

	c.Set("userId", id)
	c.Next()
}
