package user

import (
	"fmt"
	"forum/architecture/models"
	"strings"
)

func (u *UserRepo) Create(user *models.User) (int64, error) {
	row := u.db.QueryRow(`
INSERT INTO users (nickname, email, password) VALUES
(?, ?, ?) RETURNING id`, user.Nickname, user.Email, user.Password)

	err := row.Scan(&user.Id)
	switch {
	case err == nil:
		return user.Id, nil
	case strings.HasPrefix(err.Error(), "UNIQUE constraint failed"):
		switch {
		case strings.Contains(err.Error(), "nickname"):
			return -1, ErrExistNickname
		case strings.Contains(err.Error(), "email"):
			return -1, ErrExistEmail
		}
	}
	return -1, fmt.Errorf("row.Scan: %w", err)
}
