package domain

import "github.com/golang-jwt/jwt"

type CustomJwtClaim struct {
	jwt.StandardClaims
	WalletID string
}
