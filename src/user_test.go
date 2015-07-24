package app_user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticateWithValidToken(t *testing.T) {

	validToken := "johns_token"
	user, err := AuthenticateByToken(validToken)
	assert.Equal(t, nil, err, "there should be no error when using 'johns_token' as the token")
	assert.Equal(t, "john", user.username, "username should be john")
}

func TestAuthenticateWithTooLongToken(t *testing.T) {

	invalidToken := "abcdefghijklmnopqrstuvwxyz"
	_, err := AuthenticateByToken(invalidToken)
	assert.Error(t, err, "invalid token should return an error")
	switch err.(type) {
	case *TokenTooLongError:
		assert.Equal(t, "token too long - they should be max 16 characters and this one is 26 characters", err.Error(), "token was too long")
		// this is the one we want
	default:
		// all other error types are caught here, so fail if it gets to here
		t.Fail()
	}

}

func TestAuthenicateWithInvalidToken(t *testing.T) {

	invalidToken := "abc"
	_, err := AuthenticateByToken(invalidToken)
	assert.Error(t, err, "invalid token should return an error")
	switch err.(type) {
	case *UnauthorisedError:
		// this is the one we want
	default:
		// all other error types are caught here, so fail if it gets to here
		t.Fail()
	}
}
