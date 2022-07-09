package session

import (
	"fmt"
	"forum/architecture/models"
	"strings"
	"time"
)

func (s *SessionRepo) GetByUuid(uuid string) (*models.Session, error) {
	row := s.db.QueryRow(`
SELECT id, uuid, expired_at, user_id FROM sessions
WHERE uuid = ?`, uuid)
	session := &models.Session{}
	strExpiredAt := ""
	err := row.Scan(&session.Id, &session.Uuid, &strExpiredAt, &session.UserId)

	timeExpiredAt, pErr := time.ParseInLocation(timeFormat, strExpiredAt, time.Local)
	if pErr != nil {
		return nil, fmt.Errorf("time.Parse: %w", pErr)
	}
	session.ExpiredAt = timeExpiredAt

	switch {
	case err == nil:
		return session, nil
	case strings.HasPrefix(err.Error(), "sql: no rows in result set"):
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("row.Scan: %w", err)
	}
}
