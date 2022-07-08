package models

type Post struct {
	Id       int64
	Title    string
	Content  string
	Category *Category
	User     *User
	Comments []Comment
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
