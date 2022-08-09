package service

import (
	"forum/architecture/models"
	"forum/architecture/repository"
	"forum/architecture/service/post"
	"forum/architecture/service/post_category"
	"forum/architecture/service/post_vote"
	"forum/architecture/service/session"
	"forum/architecture/service/user"
)

type Service struct {
	User         models.IUserService
	Post         models.IPostService
	PostVote     models.IPostVoteService
	PostCategory models.IPostCategoryService
	Session      models.ISessionService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:         user.NewUserService(repo.User),
		Post:         post.NewPostService(repo.Post),
		PostVote:     post_vote.NewPostVoteService(repo.PostVote),
		PostCategory: post_category.NewPostCategoryService(repo.PostCategory),
		Session:      session.NewSessionService(repo.Session),
	}
}
