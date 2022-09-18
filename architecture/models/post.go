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

	WFUser      *User
	WFUserVote  int8  // -1 0 1
	WFVoteUp    int64 // Like
	WFVoteDown  int64 // Dislike
	WCategories []*PostCategory
	WComments   []*PostComment
}

type IPostService interface {
	Create(post *Post) (int64, error)
	Update(post *Post) error
	GetAll(offset, limit int64) ([]*Post, error)
	GetByID(id int64) (*Post, error)
	GetWFullPostByID(id, userId int64) (*Post, error)
	DeleteByID(id int64) error
}

type IPostRepo interface {
	Create(post *Post) (int64, error)
	Update(post *Post) error
	GetAll(offset, limit int64) ([]*Post, error)
	GetByID(id int64) (*Post, error)
	GetWFullPostByID(id, userId int64) (*Post, error)
	DeleteByID(id int64) error
}
