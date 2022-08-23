package models

import (
	"time"
)

type Post struct {
	Id        int64
	Title     string
	Content   string
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time

	WUserVote   int8  // -1 0 1
	WVoteUp     int64 // Like
	WVoteDown   int64 // Dislike
	WCategories []*PostCategory
	WComments   []*Comment
}

type IPostService interface {
	Create(user *Post) (int64, error)
	Update(user *Post) error
	GetAll(offset, limit int64) ([]*Post, error)
	GetByID(id int64) (*Post, error)
	DeleteByID(id int64) error
}

type IPostRepo interface {
	Create(user *Post) (int64, error)
	Update(user *Post) error
	GetAll(offset, limit int64) ([]*Post, error)
	GetByID(id int64) (*Post, error)
	DeleteByID(id int64) error
}
