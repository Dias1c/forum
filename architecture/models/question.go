package models

type Question struct {
	Id    int
	Title string
	Text  string
	//? Tags     []*Tag
	//? Answers  []*Answer
	UserId int
}

type IQuestionService interface {
	Create(user *Question) error
	Update(user *Question) error
	GetByID(id int) (*Question, error)
	DeleteByID(id int) error
}

type IQuestionRepo interface {
	Create(user *Question) error
	Update(user *Question) error
	GetByID(id int) (*Question, error)
	DeleteByID(id int) error
}
