package services

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	assert := assert.New(t)

	jwtKey := []byte("MY_PASS")
	logger := zerolog.New(os.Stderr)

	service := NewJwtService(jwtKey, &logger)

	// CASE 1: Generate token
	token, err := service.GenerateJWT("ID", "EMAIL", "NAME")

	assert.Nil(err, "Error must be nil")
	assert.NotEqual(token, "", "Token must not be empty")

	// CASE 2: Decode token
	claims, err := service.VerifyJWT(token)

	assert.Nil(err, "Error must be nil")
	assert.Equal(claims.Email, "EMAIL", "Email must be decoded correctly")
	assert.Equal(claims.Name, "NAME", "Name must be decoded correctly")

	// CASE 3: Fail to decode invalid token
	claims, err = service.VerifyJWT("")

	assert.Nil(claims, "Token claims must be nil")
	assert.NotNil(err, "Error must not be nil")
	fmt.Println(err)
}
