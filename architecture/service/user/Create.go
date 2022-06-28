package user

import model "forum/architecture/models"

func (u *UserService) Create(user *model.User) error {
	// Check user
	return u.repo.Create(user)
}
