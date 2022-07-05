package models

type Post struct {
	Id       int
	Title    string
	Content  string
	Category *Category
	User     *User
	Comments []Comment
}

type IPostService interface {
	Create(user *Post) error
	Update(user *Post) error
	GetByID(id int) (*Post, error)
	DeleteByID(id int) error
}

type IPostRepo interface {
	Create(user *Post) error
	Update(user *Post) error
	GetByID(id int) (*Post, error)
	DeleteByID(id int) error
}
