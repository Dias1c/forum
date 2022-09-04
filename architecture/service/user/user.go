package user

import "github.com/Dias1c/forum/architecture/models"

type UserService struct {
	repo models.IUserRepo
}

func NewUserService(repo models.IUserRepo) *UserService {
	return &UserService{repo}
}
