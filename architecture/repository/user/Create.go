package user

import (
	"fmt"
	model "forum/architecture/models"
	"strings"
)

func (u *UserRepo) Create(user *model.User) (int, error) {
	row := u.db.QueryRow(`
INSERT INTO users (nickname, email, password) VALUES
(?, ?, ?) RETURNING id`, user.Nickname, user.Email, user.Password)
	err := row.Scan(&user.Id)
	if err == nil {
		return user.Id, nil
	}

	if strings.HasPrefix(err.Error(), "UNIQUE") {
		return -1, fmt.Errorf("client: you cant create use with this nickname or email")
	}
	return -1, err
}
