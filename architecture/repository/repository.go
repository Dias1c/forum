package repository

import (
	"database/sql"

	"github.com/Dias1c/forum/architecture/models"
	"github.com/Dias1c/forum/architecture/repository/post"
	"github.com/Dias1c/forum/architecture/repository/post_category"
	"github.com/Dias1c/forum/architecture/repository/post_vote"
	"github.com/Dias1c/forum/architecture/repository/session"
	"github.com/Dias1c/forum/architecture/repository/user"
)

type Repository struct {
	User         models.IUserRepo
	Post         models.IPostRepo
	PostVote     models.IPostVoteRepo
	PostCategory models.IPostCategoryRepo
	Session      models.ISessionRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User:         user.NewUserRepo(db),
		Post:         post.NewPostRepo(db),
		PostVote:     post_vote.NewPostVoteRepo(db),
		PostCategory: post_category.NewPostCategoryRepo(db),
		Session:      session.NewSessionRepo(db),
	}
}
