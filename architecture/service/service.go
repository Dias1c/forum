package service

import (
	"github.com/Dias1c/forum/architecture/models"
	"github.com/Dias1c/forum/architecture/repository"
	"github.com/Dias1c/forum/architecture/service/category"
	"github.com/Dias1c/forum/architecture/service/post"
	"github.com/Dias1c/forum/architecture/service/post_comment"
	"github.com/Dias1c/forum/architecture/service/post_comment_vote"
	"github.com/Dias1c/forum/architecture/service/post_vote"
	"github.com/Dias1c/forum/architecture/service/session"
	"github.com/Dias1c/forum/architecture/service/user"
)

type Service struct {
	User            models.IUserService
	Post            models.IPostService
	PostVote        models.IPostVoteService
	Category        models.ICategoryService
	PostComment     models.IPostCommentService
	PostCommentVote models.IPostCommentVoteService
	Session         models.ISessionService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:            user.NewUserService(repo.User),
		Post:            post.NewPostService(repo.Post),
		PostVote:        post_vote.NewPostVoteService(repo.PostVote),
		Category:        category.NewPostCategoryService(repo.Category),
		PostComment:     post_comment.NewPostCommentService(repo.PostComment),
		PostCommentVote: post_comment_vote.NewPostCommentVoteService(repo.PostCommentVote),
		Session:         session.NewSessionService(repo.Session),
	}
}
