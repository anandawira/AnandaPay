package helper

import (
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifyToken(t *testing.T) {
	// Hardcode, later change to env
	var secretKey string = "secret"
	const id int = 1

	t.Run("It should return userId on valid token", func(t *testing.T) {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(id),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		})

		validToken, err := claims.SignedString([]byte(secretKey))
		if err != nil {
			t.Fatal("JWT token generation failed.", err.Error())
		}

		userId, err := VerifyToken(validToken)
		require.NoError(t, err)
		assert.Equal(t, id, userId)
	})
	t.Run("It should return error on invalid token", func(t *testing.T) {
		_, err := VerifyToken("invalid token")
		require.Error(t, err)
	})
	t.Run("It should return error on token signed with different secret key", func(t *testing.T) {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(id),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		})

		token, err := claims.SignedString([]byte("other secret key"))
		if err != nil {
			t.Fatal("JWT token generation failed.", err.Error())
		}

		_, err = VerifyToken(token)
		require.Error(t, err)
	})
}
