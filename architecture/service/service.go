package service

import (
	"forum/architecture/models"
	"forum/architecture/repository"
	"forum/architecture/service/category"
	"forum/architecture/service/post"
	"forum/architecture/service/session"
	"forum/architecture/service/user"
)

type Service struct {
	User     models.IUserService
	Post     models.IPostService
	Category models.ICategoryService
	Session  models.ISessionService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:     user.NewUserService(repo.User),
		Post:     post.NewPostService(repo.Post),
		Category: category.NewCategoryService(repo.Category),
		Session:  session.NewSessionService(repo.Session),
	}
}
