package user

import (
	model "forum/architecture/models"
	"strings"
)

func (u *UserRepo) Create(user *model.User) (int64, error) {
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
	return -1, err
}
