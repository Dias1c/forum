package models

type Comment struct {
	Id      int
	Content string
	User    *User
}
