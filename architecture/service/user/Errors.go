package user

import "errors"

var (
	ErrInvalidNickname = errors.New("invalid nickname")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrExistNickname   = errors.New("user with this nickname exists")
	ErrExistEmail      = errors.New("user with this email exists")
	ErrNotFound        = errors.New("user not found")
)
