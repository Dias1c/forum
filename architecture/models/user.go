package models

import "time"

// User -
type User struct {
	Id          int
	Nickname    string
	Email       string
	Password    string
	CreatedTime time.Time
}

type IUserService interface {
	Create(user *User) error
	Update(user *User) error
	DeleteByID(id int) error

	GetByID(id int) (*User, error)
	// GetByNickname(nickname string) (*model.User, error)
	// GetAll(from, offset int) error
	// CanLogin(user *model.User) error
}

type IUserRepo interface {
	Create(user *User) error
	Update(user *User) error
	DeleteByID(id int) error

	GetByID(id int) (*User, error)
	// GetByNickname(nickname string) (*model.User, error)
	// GetAll(from, offset int) error
	// CanLogin(user *model.User) error
}
