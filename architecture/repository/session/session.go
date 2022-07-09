package session

import (
	"database/sql"
	"time"
)

const (
	timeFormat = time.RFC3339
)

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{db}
}
