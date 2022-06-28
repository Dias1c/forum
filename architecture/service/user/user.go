package user

import "forum/architecture/models"

type UserService struct {
	repo models.IUserService
}

func NewUserService(repo models.IUserService) *UserService {
	return &UserService{repo}
}
