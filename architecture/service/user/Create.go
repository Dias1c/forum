package user

import (
	"errors"
	"fmt"
	"forum/architecture/models"
	ruser "forum/architecture/repository/user"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserService) Create(user *models.User) (int, error) {
	if err := user.ValidateNickname(); err != nil {
		return -1, ErrInvalidNickname
	} else if err := user.ValidateEmail(); err != nil {
		return -1, ErrInvalidEmail
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}
	user.Password = string(pass)

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
