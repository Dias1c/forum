package models

import "time"

type Session struct {
	Id        int64
	Uuid      string
	ExpiredAt time.Time
	UserId    int64
}

type ISessionRepo interface {
	Create(session *Session) error
	Delete(id int64) error
	GetByUserId(userId int64) (*Session, error)
	GetByUuid(uuid string) (*Session, error)
}

type ISessionService interface {
	Create(userId int64) (*Session, error)
	Delete(id int64) error
	GetByUserId(userId int64) (*Session, error)
	GetByUuid(uuid string) (*Session, error)
}
