package models

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	Account string
	jwt.StandardClaims
}
