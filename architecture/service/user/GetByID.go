package user

import (
	"errors"
	"fmt"
	model "forum/architecture/models"
	ruser "forum/architecture/repository/user"
)

func (u *UserService) GetByID(id int64) (*model.User, error) {
	usr, err := u.repo.GetByID(id)
	switch {
	case err == nil:
		return usr, nil
	case errors.Is(err, ruser.ErrNotFound):
		return nil, ErrNotFound
	}
	return nil, fmt.Errorf("u.repo.GetByID: %w", err)
}
