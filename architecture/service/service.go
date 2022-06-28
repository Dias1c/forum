package service

import (
	"forum/architecture/models"
	"forum/architecture/repository"
	"forum/architecture/service/post"
	"forum/architecture/service/user"
)

type Service struct {
	User models.IUserService
	Post models.IPostService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: user.NewUserService(repo.User),
		Post: post.NewPostService(repo.Question),
	}
}
