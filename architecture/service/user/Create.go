package user

import (
	"fmt"
	model "forum/architecture/models"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserService) Create(user *model.User) (int, error) {
	if err := user.ValidateNickname(); err != nil {
		return -1, fmt.Errorf("client: %w", err)
	} else if err := user.ValidateEmail(); err != nil {
		return -1, fmt.Errorf("client: %w", err)
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, fmt.Errorf("server: %w", err)
	}
	user.Password = string(pass)
	return u.repo.Create(user)
}
