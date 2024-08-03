package middleware

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
