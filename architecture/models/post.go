package models

import "time"

type Post struct {
	Id        int64
	Title     string
	Content   string
	User      int64
	CreatedAt time.Time
}

type IPostService interface {
	Create(user *Post) error
	Update(user *Post) error
	// GetPosts() ([]*Post, error)
	GetByID(id int64) (*Post, error)
	DeleteByID(id int64) error
}

type IPostRepo interface {
	Create(user *Post) error
	Update(user *Post) error
	GetByID(id int64) (*Post, error)
	DeleteByID(id int64) error
}
