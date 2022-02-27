package middleware

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	// Hardcode, later change to env
	var secretKey string = "secret"
	const id int = 1
	t.Run("It should add userId to context on valid token", func(t *testing.T) {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(id),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		})

		validToken, err := claims.SignedString([]byte(secretKey))
		if err != nil {
			t.Fatal("JWT token generation failed.", err.Error())
		}

		req := httptest.NewRequest("POST", "/", nil)
		req.Header.Add("Authorization", "Bearer "+validToken)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		Authenticate(c)

		assert.Equal(t, id, c.GetInt("userId"))
	})
	t.Run("It should return StatusStatusUnauthorized on invalid token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", nil)
		req.Header.Add("Authorization", "Bearer invalidtoken")
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		Authenticate(c)

		helper.AssertResponse(t, http.StatusUnauthorized, gin.H{"message": domain.ErrInvalidToken.Error()}, rec)
	})
}
