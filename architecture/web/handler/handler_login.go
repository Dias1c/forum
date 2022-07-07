package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	suser "forum/architecture/service/user"
	"forum/architecture/web/handler/view"
)

// LogInHandler -
func (m *MainHandler) LogInHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("LogInHandler", r)

	switch r.Method {
	case http.MethodGet:
		if cookie, err := r.Cookie("session"); err == nil && cookie != nil {
			pg := &view.Page{Warn: fmt.Errorf("you already signed in!")}
			m.view.ExecuteTemplate(w, pg, "login.html")
			return
		}
		fmt.Println(r.URL.Query())
		if redirect := r.URL.Query().Get("redirect_to"); redirect != "" {
			http.SetCookie(w,
				&http.Cookie{
					Name:   "redirect_to",
					Value:  redirect,
					MaxAge: 3600,
				},
			)
		}
		m.view.ExecuteTemplate(w, nil, "login.html")
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Printf("LogInHandler: r.ParseForm: %v\n", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "login.html")
			return
		}

		usr, err := m.service.User.GetByNicknameOrEmail(r.FormValue("login"))
		switch {
		case err == nil:
		case errors.Is(err, suser.ErrNotFound):
			pg := &view.Page{Error: fmt.Errorf("user with login \"%v\" not found", r.FormValue("login"))}
			// w.WriteHeader(http.StatusNotFound)
			m.view.ExecuteTemplate(w, pg, "login.html")
			return
		case errors.Is(err, suser.ErrInvalidEmail):
			pg := &view.Page{Error: fmt.Errorf("invalid email %v", r.FormValue("login"))}
			// w.WriteHeader(http.StatusBadRequest)
			m.view.ExecuteTemplate(w, pg, "login.html")
			return
		case errors.Is(err, suser.ErrInvalidNickname):
			pg := &view.Page{Error: fmt.Errorf("invalid nickname %v", r.FormValue("login"))}
			// w.WriteHeader(http.StatusBadRequest)
			m.view.ExecuteTemplate(w, pg, "login.html")
			return
		default:
			log.Printf("ERROR: LogInHandler: User.GetByNicknameOrEmail: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "login.html")
			return
		}

		areEqual, err := usr.CompareHashAndPassword(r.FormValue("password"))
		switch {
		case err != nil:
			log.Printf("ERROR: LogInHandler: user.CompareHashAndPassword: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "login.html")
			return
		case !areEqual:
			pg := &view.Page{Error: fmt.Errorf("invalid password for login \"%s\"", r.FormValue("login"))}
			m.view.ExecuteTemplate(w, pg, "login.html")
			return
		}

		session, err := m.service.Session.Record(usr.Id)
		if err != nil {
			log.Printf("ERROR: LogInHandler: Session.Record: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "login.html")
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

		if cookie, err := r.Cookie("redirect_to"); err == nil && cookie != nil {
			redirectTo := cookie.Value
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)
			http.Redirect(w, r, redirectTo, http.StatusMovedPermanently)
			return
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
