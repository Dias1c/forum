package session

import (
	"database/sql"
)

const (
	timeFormat = "2006-01-02T15:04:05Z"
)

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{db}
}
