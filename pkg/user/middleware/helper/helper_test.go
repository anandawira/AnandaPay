package helper

import (
	"strconv"
	"testing"
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifyToken(t *testing.T) {
	// Hardcode, later change to env
	var secretKey string = "secret"
	const userId int = 1
	const walletId string = "wallet id"

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.CustomJwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    strconv.Itoa(userId),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		},
		WalletID: walletId,
	})
	t.Run("It should return userId and walletId on valid token", func(t *testing.T) {
		validToken, err := claims.SignedString([]byte(secretKey))
		if err != nil {
			t.Fatal("JWT token generation failed.", err.Error())
		}

		resUserId, resWalletId, err := VerifyToken(validToken)
		require.NoError(t, err)
		assert.Equal(t, userId, resUserId)
		assert.Equal(t, walletId, resWalletId)
	})
	t.Run("It should return error on invalid token", func(t *testing.T) {
		_, _, err := VerifyToken("invalid token")
		require.Error(t, err)
	})
	t.Run("It should return error on token signed with different secret key", func(t *testing.T) {
		token, err := claims.SignedString([]byte("other secret key"))
		if err != nil {
			t.Fatal("JWT token generation failed.", err.Error())
		}

		_, _, err = VerifyToken(token)
		require.Error(t, err)
	})
}
