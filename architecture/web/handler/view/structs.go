package view

import "forum/architecture/models"

type Page struct {
	User  *models.User
	Users *[]models.User
	Post  *models.Post
	Posts *[]models.Post
	// Comments           []models.Comment
	Error error // Error - Notification Error
	Warn  error // Warn - Notification Warning
	Info  error // Info - Notification Info
}

type View struct {
	templatesDir string
}
