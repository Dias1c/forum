package models

type Category struct {
	Id   int64
	Name string
}

type ICategoryRepo interface {
	Create(category *Category) (int64, error)
	Update(category *Category) error
	GetByID(id int64) (*Category, error)
	GetByName(name string) (*Category, error)
	DeleteByID(id int64) error
}

type ICategoryService interface {
	AddToPostByNames(names []string, postId int64) error
	GetByPostID(postId int64) ([]*Category, error)
	DeleteFromPost(id int64) error
}
