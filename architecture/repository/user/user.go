package user

import (
	"database/sql"
	"time"
)

const (
	timeFormat = time.RFC3339
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}
