package helpers

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var JWT *JWTService

func init() {
	jwtServiceInstance, err := NewJWTService(Config{
		SecretKey:                "secret",
		TokenExpirationInSeconds: 5,
	})
	if err != nil {
		log.Fatalln(err)
	}

	JWT = jwtServiceInstance
}

func generateJWT(t *testing.T, claims Claims) string {
	token, err := JWT.GenerateJWT(claims)
	assert.NoError(t, err)

	return token
}

func TestGenerateJWT(t *testing.T) {
	token := generateJWT(t, Claims{1})

	t.Logf("created token : %s", token)
}

func TestValidateJWT(t *testing.T) {
	c := Claims{
		UserID: 1,
	}

	token := generateJWT(t, c)

	claims, err := JWT.ValidateJWT(token)
	assert.NoError(t, err)

	assert.Equal(t, claims.UserID, c.UserID)

	claims, err = JWT.ValidateJWT("its.invalid.token")
	assert.Error(t, err)

	time.Sleep(time.Second * 6)

	claims, err = JWT.ValidateJWT(token)
	assert.Error(t, err)
}
