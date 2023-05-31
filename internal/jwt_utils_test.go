package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateErrorForEmptyJWT(t *testing.T) {
	expected := "JWT token to verify must be non-empty"
	actualStr, actualErr := VerifyAccessJWT("")
	assert.Containsf(t, actualErr.Error(), expected, "Expected error containing %v, got %v.",
		expected, actualErr)
	assert.Equalf(t, actualStr, "", "Expected empty string, got %v.", actualStr)
}

func TestValidateErrorForGarbageJWTString(t *testing.T) {
	expected := "could not parse supplied JWT"
	actualStr, actualErr := VerifyAccessJWT("GARBAGE-STRING")
	assert.Containsf(t, actualErr.Error(), expected, "Expected error containing %v, got %v.",
		expected, actualErr)
	assert.Equalf(t, actualStr, "", "Expected empty string, got %v.", actualStr)

}

func TestCanValidateGeneratedAccessJWT(t *testing.T) {
	token, err := GenerateAccessJWT("USERNAME")
	assert.NotEqualf(t, token, nil, "Expected generated JWT, received nil.")
	assert.Equalf(t, err, nil, "Recieved unexpected error generating JWT.")

	expectedStr := "USERNAME"
	actualStr, actualErr := VerifyAccessJWT(token)
	assert.Nil(t, actualErr, "Received unexpected error %v.", actualErr)
	assert.Equalf(t, expectedStr, actualStr, "Expected return value %v, got %v.", expectedStr, actualStr)
}

func TestGenerateAccessJWTWithTTLExpires(t *testing.T) {
	token, err := GenerateAccessJWTWithTTL("USERNAME", 1) // 1 second TTL
	assert.NotEqualf(t, token, nil, "Expected generated JWT, received nil.")
	assert.Equalf(t, err, nil, "Recieved unexpected error generating JWT.")

	// Needs more than 1 second to reliably be expired, even though a 1 second TTL
	time.Sleep(2 * time.Second)

	expected := "Token is expired"
	actualStr, actualErr := VerifyAccessJWT(token)
	assert.Containsf(t, actualErr.Error(), expected, "Expected error containing %v, got %v.",
		expected, actualErr)
	assert.Equalf(t, actualStr, "", "Expected empty string, got %v.", actualStr)
}
