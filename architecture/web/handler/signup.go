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
		m.executeTemplate(w, nil, "signup.html")
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "r.ParseForm: %v\n", err)
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
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		case errors.Is(err, suser.ErrExistNickname):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Is(err, suser.ErrExistEmail):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Is(err, suser.ErrInvalidNickname):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Is(err, suser.ErrInvalidEmail):
			w.WriteHeader(http.StatusBadRequest)
		default:
			// w.WriteHeader(http.StatusInternalServerError)
			log.Printf("ERROR: SignUpHandler: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintln(w, err)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %v not allowed\n", r.Method)
	}
}
