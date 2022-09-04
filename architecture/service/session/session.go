package session

import "github.com/Dias1c/forum/architecture/models"

type SessionService struct {
	repo models.ISessionRepo
}

func NewSessionService(repo models.ISessionRepo) *SessionService {
	return &SessionService{repo}
}
