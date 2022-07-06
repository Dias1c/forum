package post

import "database/sql"

type QuestionRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *QuestionRepo {
	return &QuestionRepo{db}
}
