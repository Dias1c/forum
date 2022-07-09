package user

import (
	"fmt"
	"forum/architecture/models"
	"strings"
	"time"
)

func (u *UserRepo) GetByNickname(nickname string) (*models.User, error) {
	row := u.db.QueryRow(`
SELECT id, nickname, email, password, created_at FROM users
WHERE nickname = ?`, nickname)
	user := &models.User{}
	strCreatedAt := ""

	err := row.Scan(&user.Id, &user.Nickname, &user.Email, &user.Password, &strCreatedAt)

	switch {
	case err == nil:
		timeCreatedAt, err := time.ParseInLocation(timeFormat, strCreatedAt, time.Local)
		if err != nil {
			return nil, fmt.Errorf("time.Parse: %w", err)
		}
		user.CreatedAt = timeCreatedAt
		return user, nil
	case strings.HasPrefix(err.Error(), "sql: no rows in result set"):
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("row.Scan: %w", err)
	}
}
