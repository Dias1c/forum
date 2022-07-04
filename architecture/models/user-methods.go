package models

import (
	"fmt"
	"net/mail"
)

func (u *User) ValidateNickname() error {
	if lng := len(u.Nickname); lng < 1 || 32 < lng {
		return fmt.Errorf("nickname: invalid lenght (%d)", lng)
	}
	for _, c := range u.Nickname {
		if c < 33 || 125 < c {
			return fmt.Errorf("nickname: invalid character '%c'", c)
		}
	}
	return nil
}

func (u *User) ValidateEmail() error {
	if lng := len(u.Email); lng < 1 || 320 < lng {
		return fmt.Errorf("email: invalid lenght (%d)", lng)
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ValidatePassword(confirm string) error {
	if u.Password != confirm {
		return fmt.Errorf("password: different passwords %q != %q", u.Password, confirm)
	}
	return nil
}
