package session

import (
	"errors"
	"fmt"
	"time"

	"github.com/Dias1c/forum/architecture/models"

	rsession "github.com/Dias1c/forum/architecture/repository/session"

	uuid "github.com/satori/go.uuid"
)

func (s *SessionService) Record(userId int64) (*models.Session, error) {
	uid := uuid.NewV4()
	session := &models.Session{
		Uuid:      uid.String(),
		UserId:    userId,
		ExpiredAt: time.Now().Add(models.SessionExpiredAfter),
	}

	_, err := s.repo.Create(session)
	switch {
	case err == nil:
		return session, nil
	case errors.Is(err, rsession.ErrSessionExists):
		err := s.repo.UpdateByUserId(session.UserId, session)
		if err != nil {
			return nil, fmt.Errorf("s.repo.UpdateByUserId: %w", err)
		}
		return session, nil
	default:
		return nil, fmt.Errorf("s.repo.Create: %w", err)
	}
}
