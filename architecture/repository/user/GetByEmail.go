package user

import (
	"errors"
	"forum/architecture/models"

	"github.com/mattn/go-sqlite3"
)

func (u *UserRepo) GetByEmail(email string) (*models.User, error) {
	row := u.db.QueryRow(`
SELECT id, nickname, email, password FROM user
WHERE email = ?`)
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Nickname, &user.Email, &user.Password)
	switch {
	case err == nil:
		return user, nil
	case errors.Is(err, sqlite3.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
