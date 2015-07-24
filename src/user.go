package app_user

import "fmt"

type UnauthorisedError struct{}

func (e *UnauthorisedError) Error() string {
	return "you are not authorised"
}

type TokenTooLongError struct {
	length int
}

func (e *TokenTooLongError) Error() string {
	return fmt.Sprintf("token too long - they should be max 16 characters and this one is %d characters", e.length)
}

type User struct {
	username string
}

func AuthenticateByToken(authtoken string) (*User, error) {
	validToken := "johns_token"
	if 16 < len(authtoken) {
		return nil, &TokenTooLongError{length: len(authtoken)}
	}
	if validToken != authtoken {
		return nil, &UnauthorisedError{}
	}
	return &User{username: "john"}, nil
}
