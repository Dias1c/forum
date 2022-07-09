package user

import (
	"fmt"
	"forum/architecture/models"
	"strings"

	model "forum/architecture/models"
)

func (u *UserRepo) GetByID(id int64) (*model.User, error) {
	row := u.db.QueryRow(`
SELECT id, nickname, email FROM users
WHERE id = ?`, id)
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Nickname, &user.Email)
	switch {
	case err == nil:
		return user, nil
	case strings.HasPrefix(err.Error(), "sql: no rows in result set"):
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("row.Scan: %w", err)
	}
}
