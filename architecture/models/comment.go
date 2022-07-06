package models

type Comment struct {
	Id      int64
	Content string
	User    *User
}
