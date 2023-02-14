package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	ErrInvalidJWTConfig = errors.New("invalid jwt config")
)

type Claims struct {
	UserID uint
}

type tokenClaims struct {
	jwt.StandardClaims
	Claims
}

var jwtInstance *JWTService

type Config struct {
	SecretKey                string
	TokenExpirationInSeconds time.Duration
}

type JWTService struct {
	Config
}

func NewJWTService(jwtConfig Config) (*JWTService, error) {
	if jwtInstance == nil {
		if jwtConfig.SecretKey == "" || jwtConfig.TokenExpirationInSeconds < 1 {
			return nil, ErrInvalidJWTConfig
		}

		jwtInstance = &JWTService{
			Config: jwtConfig,
		}
	}

	return jwtInstance, nil
}

func (js *JWTService) GenerateJWT(claims Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * js.TokenExpirationInSeconds).Unix(),
		},
		Claims: claims,
	}).SignedString([]byte(js.SecretKey))
}

func (js *JWTService) ValidateJWT(token string) (Claims, error) {
	claims := new(tokenClaims)

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(js.SecretKey), nil
	})
	if err != nil {
		return Claims{}, err
	}

	return (*claims).Claims, nil
}
