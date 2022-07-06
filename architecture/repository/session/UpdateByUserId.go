package session

import (
	"fmt"
	"forum/architecture/models"
)

func (s *SessionRepo) UpdateByUserId(userId int64, session *models.Session) error {
	row := s.db.QueryRow(`
UPDATE sessions 
SET uuid = ?, expired_at = ?
WHERE user_id = ?
RETURNING id`, session.Uuid, session.ExpiredAt, session.UserId)

	err := row.Scan(&session.Id)
	switch {
	case err == nil:
		return nil
	}
	return fmt.Errorf("row.Scan: %w", err)
}
