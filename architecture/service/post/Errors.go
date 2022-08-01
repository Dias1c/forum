package post

import "errors"

var (
	ErrInvalidTitleLength = errors.New("invalid title length")
	ErrNotFound           = errors.New("post not found")
)
