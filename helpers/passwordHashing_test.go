package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	p := "test"

	_, err := HashPassword(p)
	assert.NoError(t, err)

	p = "thisisverylongpassworderrorerrorerrofsfksehfksejfhkesrerrorerrorerrorerrorerrorerrorerrorerrorerrorerror"
	_, err = HashPassword(p)
	assert.Error(t, err)
}

func TestComparePassword(t *testing.T) {
	p := "test"

	hashedPassword, err := HashPassword(p)
	assert.NoError(t, err)

	isMatch, err := ComparePassword(p, hashedPassword)
	assert.NoError(t, err)

	assert.True(t, isMatch)

	p = "wrongpassword"

	isMatch, err = ComparePassword(p, hashedPassword)
	assert.NoError(t, err)

	assert.False(t, isMatch)
}
