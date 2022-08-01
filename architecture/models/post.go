package models

import "time"

type Post struct {
	Id        int64
	Title     string
	Content   string
	UserId    int64
	CreatedAt time.Time

	WCategories []*Category
}

type IPostService interface {
	Create(user *Post) (int64, error)
	Update(user *Post) error
	// GetPosts() ([]*Post, error)
	GetByID(id int64) (*Post, error)
	DeleteByID(id int64) error
}

type IPostRepo interface {
	Create(user *Post) (int64, error)
	Update(user *Post) error
	GetByID(id int64) (*Post, error)
	DeleteByID(id int64) error
}
