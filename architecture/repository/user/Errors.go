package user

import "errors"

var (
	ErrExistNickname = errors.New("user with this nickname exists")
	ErrExistEmail    = errors.New("user with this email exists")

	ErrNotFound = errors.New("user not found")
)
