package helpers

import (
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func generateJWT(t *testing.T, tokenClaims TokenClaims) string {
	token, err := GenerateJWT(tokenClaims, "secret")
	assert.NoError(t, err)

	return token
}

func TestGenerateJWT(t *testing.T) {
	token := generateJWT(t, TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 3).Unix(),
		},
		Claims: Claims{
			UserID: 1,
		},
	})

	t.Logf("created token : %s", token)
}

func TestValidateJWT(t *testing.T) {
	tc := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 3).Unix(),
		},
		Claims: Claims{
			UserID: 1,
		},
	}

	token := generateJWT(t, tc)

	claims, err := ValidateJWT(token, "secret")
	assert.NoError(t, err)

	assert.Equal(t, claims.UserID, tc.UserID)

	claims, err = ValidateJWT("its.invalid.token", "secret")
	assert.Error(t, err)

	time.Sleep(time.Second * 5)

	claims, err = ValidateJWT(token, "secret")
	assert.Error(t, err)
}
