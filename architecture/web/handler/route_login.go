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
		if cookie, err := r.Cookie("session"); err == nil && cookie != nil {
			pg := &page{Warn: fmt.Errorf("you already signed in!")}
			m.executeTemplate(w, pg, "login.html")
			return
		}
		m.executeTemplate(w, nil, "login.html")
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Printf("LogInHandler: r.ParseForm: %v\n", err)
			pg := &page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.executeTemplate(w, pg, "login.html")
			return
		}

		usr, err := m.service.User.GetByNicknameOrEmail(r.FormValue("login"))
		switch {
		case err == nil:
		case errors.Is(err, suser.ErrNotFound):
			pg := &page{Error: fmt.Errorf("user with login \"%v\" not found", r.FormValue("login"))}
			// w.WriteHeader(http.StatusNotFound)
			m.executeTemplate(w, pg, "login.html")
			return
		case errors.Is(err, suser.ErrInvalidEmail):
			pg := &page{Error: fmt.Errorf("invalid email %v", r.FormValue("login"))}
			// w.WriteHeader(http.StatusBadRequest)
			m.executeTemplate(w, pg, "login.html")
			return
		case errors.Is(err, suser.ErrInvalidNickname):
			pg := &page{Error: fmt.Errorf("invalid nickname %v", r.FormValue("login"))}
			// w.WriteHeader(http.StatusBadRequest)
			m.executeTemplate(w, pg, "login.html")
			return
		default:
			log.Printf("ERROR: LogInHandler: User.GetByNicknameOrEmail: %s", err)
			pg := &page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.executeTemplate(w, pg, "login.html")
			return
		}

		areEqual, err := usr.CompareHashAndPassword(r.FormValue("password"))
		switch {
		case err != nil:
			log.Printf("ERROR: LogInHandler: user.CompareHashAndPassword: %s", err)
			pg := &page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.executeTemplate(w, pg, "login.html")
			return
		case !areEqual:
			pg := &page{Error: fmt.Errorf("invalid password for login \"%s\"", r.FormValue("login"))}
			m.executeTemplate(w, pg, "login.html")
			return
		}

		session, err := m.service.Session.Record(usr.Id)
		if err != nil {
			log.Printf("ERROR: LogInHandler: Session.Record: %s", err)
			pg := &page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.executeTemplate(w, pg, "login.html")
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
	}
}
