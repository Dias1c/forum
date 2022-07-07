package view

import "forum/architecture/models"

type Page struct {
	User  *models.User
	Users *[]models.User
	Post  *models.Post
	Posts *[]models.Post
	// Comments           []models.Comment
	Error error
	Warn  error
}

type View struct {
	templatesDir string
}
