package category

import (
	"database/sql"
)

type PostCategoryRepo struct {
	db *sql.DB
}

func NewPostCategoryRepo(db *sql.DB) *PostCategoryRepo {
	return &PostCategoryRepo{db}
}
