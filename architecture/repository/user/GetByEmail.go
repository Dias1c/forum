package user

import (
	"fmt"
	"forum/architecture/models"
	"strings"
)

func (u *UserRepo) GetByEmail(email string) (*models.User, error) {
	row := u.db.QueryRow(`
SELECT id, nickname, email, password FROM users
WHERE email = ?`, email)
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Nickname, &user.Email, &user.Password)
	switch {
	case err == nil:
		return user, nil
	case strings.HasPrefix(err.Error(), "sql: no rows in result set"):
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("row.Scan: %w", err)
	}
}
