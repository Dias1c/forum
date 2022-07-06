package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	suser "forum/architecture/service/user"
)

// LogInHandler -
func (m *MainHandler) LogInHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("LogInHandler", r)

	switch r.Method {
	case http.MethodGet:
		m.executeTemplate(w, nil, "login.html")
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "r.ParseForm: %v\n", err)
			return
		}

		usr, err := m.service.User.GetByNicknameOrEmail(r.FormValue("login"))
		switch {
		case err == nil:
		case errors.Is(err, suser.ErrNotFound):
			w.WriteHeader(http.StatusNotFound) //
		case errors.Is(err, suser.ErrInvalidEmail):
			// 200 Не правильный email
		case errors.Is(err, suser.ErrInvalidNickname):
			// 200 Не правильный nickname
		default:
			log.Printf("ERROR: LogInHandler: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		areEqual, err := usr.CompareHashAndPassword(r.FormValue("password"))
		if err != nil {
			log.Printf("ERROR: LogInHandler: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !areEqual {
			// 200 Пароли разные
			return
		}

		// Создать сессию, Записать в БД
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %v not allowed\n", r.Method)
	}
}
