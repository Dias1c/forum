package repository

import (
	"database/sql"

	"forum/architecture/models"
	"forum/architecture/repository/post"
	"forum/architecture/repository/session"
	"forum/architecture/repository/user"
)

type Repository struct {
	User    models.IUserRepo
	Post    models.IPostRepo
	Session models.ISessionRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User:    user.NewUserRepo(db),
		Post:    post.NewPostRepo(db),
		Session: session.NewSessionRepo(db),
	}
}
