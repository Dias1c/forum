package post_vote

import "errors"

var (
	ErrNotFound = errors.New("no one user vote in post")

// ErrExistNickname       = errors.New("user with this nickname exists")
// ErrExistEmail          = errors.New("user with this email exists")
// ErrWrongLengthNickname = errors.New("user nickname length is wrong")
// ErrWrongLengthEmail    = errors.New("user email length is wrong")
)
