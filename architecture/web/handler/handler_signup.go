package handler

import (
	"errors"
	"fmt"
	"forum/architecture/models"
	"forum/architecture/web/handler/cookies"
	"forum/architecture/web/handler/view"
	"log"
	"net/http"

	suser "forum/architecture/service/user"
)

// SignUpHandler -
func (m *MainHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("SignUpHandler", r)

	switch r.Method {
	case http.MethodGet:
		cookies.AddRedirectCookie(w, r.URL.Query().Get("redirect_to"))
		if cookie := cookies.GetSessionCookie(w, r); cookie != nil {
			pg := &view.Page{Warn: fmt.Errorf("you already signed in!")}
			m.view.ExecuteTemplate(w, pg, "signup.html")
			return
		}
		m.view.ExecuteTemplate(w, nil, "signup.html")
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Printf("SignUpHandler: r.ParseForm: %v\n", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "login.html")
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
			pg := &view.Page{Error: fmt.Errorf("nickname \"%v\" is used. Try with another nickname.", newUser.Nickname)}
			m.view.ExecuteTemplate(w, pg, "signup.html")
		case errors.Is(err, suser.ErrExistEmail):
			pg := &view.Page{Error: fmt.Errorf("email \"%v\" is used. Try with another email.", newUser.Nickname)}
			m.view.ExecuteTemplate(w, pg, "signup.html")
		case errors.Is(err, suser.ErrInvalidNickname):
			pg := &view.Page{Error: fmt.Errorf("invalid nickname \"%v\"", newUser.Nickname)}
			m.view.ExecuteTemplate(w, pg, "signup.html")
		case errors.Is(err, suser.ErrInvalidEmail):
			pg := &view.Page{Error: fmt.Errorf("invalid email \"%v\"", newUser.Email)}
			m.view.ExecuteTemplate(w, pg, "signup.html")
		default:
			log.Printf("ERROR: SignUpHandler: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "signup.html")
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
