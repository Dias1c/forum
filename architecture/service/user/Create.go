package user

import (
	"errors"
	"fmt"
	"forum/architecture/models"
	ruser "forum/architecture/repository/user"
)

func (u *UserService) Create(user *models.User) (int, error) {
	if err := user.ValidateNickname(); err != nil {
		return -1, ErrInvalidNickname
	} else if err := user.ValidateEmail(); err != nil {
		return -1, ErrInvalidEmail
	}

	err := user.HashPassword()
	if err != nil {
		return -1, fmt.Errorf("user.HashPassword: %w", err)
	}

	userId, err := u.repo.Create(user)
	switch {
	case err == nil:
		return userId, nil
	case errors.Is(err, ruser.ErrExistEmail):
		return -1, ErrExistEmail
	case errors.Is(err, ruser.ErrExistNickname):
		return -1, ErrExistNickname
	}
	return -1, err
}
