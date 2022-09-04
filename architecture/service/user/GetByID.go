package user

import (
	"errors"
	"fmt"

	model "github.com/Dias1c/forum/architecture/models"
	ruser "github.com/Dias1c/forum/architecture/repository/user"
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
