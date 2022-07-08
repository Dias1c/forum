package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	suser "forum/architecture/service/user"
	"forum/architecture/web/handler/cookies"
	"forum/architecture/web/handler/view"
)

// SignInHandler -
func (m *MainHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("SignInHandler", r)

	switch r.Method {
	case http.MethodGet:
		cookies.AddRedirectCookie(w, r.URL.Query().Get("redirect_to"))

		if cookie := cookies.GetSessionCookie(w, r); cookie != nil {
			pg := &view.Page{Warn: fmt.Errorf("you already signed in!")}
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		}
		m.view.ExecuteTemplate(w, nil, "signin.html")
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Printf("SignInHandler: r.ParseForm: %v\n", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		}

		usr, err := m.service.User.GetByNicknameOrEmail(r.FormValue("login"))
		switch {
		case err == nil:
		case errors.Is(err, suser.ErrNotFound):
			pg := &view.Page{Error: fmt.Errorf("user with login \"%v\" not found", r.FormValue("login"))}
			// w.WriteHeader(http.StatusNotFound)
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		case errors.Is(err, suser.ErrInvalidEmail):
			pg := &view.Page{Error: fmt.Errorf("invalid email %v", r.FormValue("login"))}
			// w.WriteHeader(http.StatusBadRequest)
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		case errors.Is(err, suser.ErrInvalidNickname):
			pg := &view.Page{Error: fmt.Errorf("invalid nickname %v", r.FormValue("login"))}
			// w.WriteHeader(http.StatusBadRequest)
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		default:
			log.Printf("ERROR: SignInHandler: User.GetByNicknameOrEmail: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		}

		areEqual, err := usr.CompareHashAndPassword(r.FormValue("password"))
		switch {
		case err != nil:
			log.Printf("ERROR: SignInHandler: user.CompareHashAndPassword: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		case !areEqual:
			pg := &view.Page{Error: fmt.Errorf("invalid password for login \"%s\"", r.FormValue("login"))}
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		}

		session, err := m.service.Session.Record(usr.Id)
		if err != nil {
			log.Printf("ERROR: SignInHandler: Session.Record: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "signin.html")
			return
		}
		expiresAfterSeconds := time.Until(session.ExpiredAt).Seconds()
		cookies.AddSessionCookie(w, session.Uuid, int(expiresAfterSeconds))

		// TODO: Добавить в главном хендлере удаление этого куки
		if cookie := cookies.GetRedirectCookie(w, r); cookie != nil {
			cookies.RemoveRedirectCookie(w, r)
			http.Redirect(w, r, cookie.Value, http.StatusMovedPermanently)
			return
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
