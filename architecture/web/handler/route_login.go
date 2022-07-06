package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

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
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "USER NOT FOUND", err)
			return
		case errors.Is(err, suser.ErrInvalidEmail):
			// 200 Не правильный email
			fmt.Fprintln(w, "INVALID EMAIL", err)
			return
		case errors.Is(err, suser.ErrInvalidNickname):
			// 200 Не правильный nickname
			fmt.Fprintln(w, "INVALID NICKNAME", err)
			return
		default:
			log.Printf("ERROR: LogInHandler: User.GetByNicknameOrEmail: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		areEqual, err := usr.CompareHashAndPassword(r.FormValue("password"))
		switch {
		case err != nil:
			log.Printf("ERROR: LogInHandler: user.CompareHashAndPassword: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		case !areEqual:
			// 200 Пароли разные
			log.Println("Password not equal")
			fmt.Fprintln(w, "PASSWORD NOT EQUAL")
			return
		}

		session, err := m.service.Session.Record(usr.Id)
		if err != nil {
			log.Printf("ERROR: LogInHandler: Session.Record: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		expiresAfterSeconds := time.Until(session.ExpiredAt).Seconds()
		http.SetCookie(w,
			&http.Cookie{
				Name:   "session",
				Value:  session.Uuid,
				MaxAge: int(expiresAfterSeconds),
			},
		)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %v not allowed\n", r.Method)
	}
}
