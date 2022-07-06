package repository

import (
	"database/sql"

	"forum/architecture/models"
	"forum/architecture/repository/post"
	"forum/architecture/repository/user"
)

type Repository struct {
	User models.IUserRepo
	Post models.IPostRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User: user.NewUserRepo(db),
		Post: post.NewPostRepo(db),
	}
}
