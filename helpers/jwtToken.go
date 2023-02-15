package helpers

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID uint
}

type TokenClaims struct {
	jwt.StandardClaims
	Claims
}

func GenerateJWT(tokenClaims TokenClaims, secretKey string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString([]byte(secretKey))
}

func ValidateJWT(token string, secretKey string) (TokenClaims, error) {
	tokenClaims := new(TokenClaims)

	_, err := jwt.ParseWithClaims(token, tokenClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return TokenClaims{}, err
	}

	return *tokenClaims, nil
}
