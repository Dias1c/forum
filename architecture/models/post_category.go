package models

import "time"

type PostCategory struct {
	Id        int64
	Name      string
	CreatedAt time.Time
}

type IPostCategoryRepo interface {
	Create(category *PostCategory) (int64, error)
	AddToPost(categoryId, postId int64) (int64, error)
	Update(category *PostCategory) error
	GetByID(id int64) (*PostCategory, error)
	GetByName(name string) (*PostCategory, error)
	GetByPostID(postId int64) ([]*PostCategory, error)
	DeleteByPostID(postId int64) error
	DeleteByID(id int64) error
}

type IPostCategoryService interface {
	AddToPostByNames(names []string, postId int64) error
	GetByPostID(postId int64) ([]*PostCategory, error)
	DeleteByPostID(postId int64) error
	DeleteFromPost(id int64) error
}
