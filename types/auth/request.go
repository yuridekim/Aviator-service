package auth

import "github.com/dgrijalva/jwt-go/v4"

type AuthTokenClaims struct {
	Username string   `json:"Username"`
	Name     string   `json:"Name"`
	Role     []string `json:"role"`
	jwt.StandardClaims
}
