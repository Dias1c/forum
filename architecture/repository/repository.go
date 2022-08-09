package repository

import (
	"database/sql"

	"forum/architecture/models"
	"forum/architecture/repository/post"
	"forum/architecture/repository/post_category"
	"forum/architecture/repository/post_vote"
	"forum/architecture/repository/session"
	"forum/architecture/repository/user"
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
