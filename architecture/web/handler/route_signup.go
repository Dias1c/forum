package handler

import (
	"errors"
	"fmt"
	"forum/architecture/models"
	"log"
	"net/http"

	suser "forum/architecture/service/user"
)

// SignUpHandler -
func (m *MainHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("SignUpHandler", r)

	switch r.Method {
	case http.MethodGet:
		if cookie, err := r.Cookie("session"); err == nil && cookie != nil {
			pg := &page{Warn: fmt.Errorf("you already signed in!")}
			m.executeTemplate(w, pg, "signup.html")
			return
		}
		m.executeTemplate(w, nil, "signup.html")
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Printf("SignUpHandler: r.ParseForm: %v\n", err)
			pg := &page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.executeTemplate(w, pg, "login.html")
			return
		}

		newUser := &models.User{
			Nickname: r.FormValue("nickname"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		_, err = m.service.User.Create(newUser)
		switch {
		case err == nil:
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		case errors.Is(err, suser.ErrExistNickname):
			pg := &page{Error: fmt.Errorf("nickname \"%v\" is used. Try with another nickname.", newUser.Nickname)}
			m.executeTemplate(w, pg, "signup.html")
		case errors.Is(err, suser.ErrExistEmail):
			pg := &page{Error: fmt.Errorf("email \"%v\" is used. Try with another email.", newUser.Nickname)}
			m.executeTemplate(w, pg, "signup.html")
		case errors.Is(err, suser.ErrInvalidNickname):
			pg := &page{Error: fmt.Errorf("invalid nickname \"%v\"", newUser.Nickname)}
			m.executeTemplate(w, pg, "signup.html")
		case errors.Is(err, suser.ErrInvalidEmail):
			pg := &page{Error: fmt.Errorf("invalid email \"%v\"", newUser.Email)}
			m.executeTemplate(w, pg, "signup.html")
		default:
			log.Printf("ERROR: SignUpHandler: %s", err)
			pg := &page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.executeTemplate(w, pg, "signup.html")
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
