package helper

import (
	"log"
	"strconv"

	"github.com/anandawira/anandapay/domain"
	"github.com/golang-jwt/jwt"
)

func VerifyToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrInvalidToken
		}

		// Hardcode, later change to env
		return []byte("secret"), nil
	})

	if err != nil {
		return 0, domain.ErrInvalidToken
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		id, err := strconv.Atoi(claims.Issuer)
		if err != nil {
			log.Fatal(domain.ErrInternalServerError)
		}
		return id, nil
	} else {
		return 0, domain.ErrInvalidToken
	}
}
