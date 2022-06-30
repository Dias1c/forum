package service

import (
	"forum/architecture/models"
	"forum/architecture/repository"
	"forum/architecture/service/question"
	"forum/architecture/service/user"
)

type Service struct {
	User     models.IUserService
	Question models.IQuestionService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:     user.NewUserService(repo.User),
		Question: question.NewQuestionService(repo.Question),
	}
}
