package handler

import (
	"errors"
	"net/http"

	"forum/architecture/web/handler/cookies"
	"forum/architecture/web/handler/view"
	"forum/internal/lg"

	ssession "forum/architecture/service/session"
)

// IndexHandler -
func (m *MainHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("IndexHandler", r)

	// Allowed Methods
	switch r.Method {
	case http.MethodGet:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Logic
	cookies.RemoveRedirectCookie(w, r)
	switch r.Method {
	case http.MethodGet:
		cookie := cookies.GetSessionCookie(w, r)
		if cookie == nil {
			m.view.ExecuteTemplate(w, nil, "home.html")
			return
		}

		session, err := m.service.Session.GetByUuid(cookie.Value)
		switch {
		case err == nil:
		case errors.Is(err, ssession.ErrExpired) || errors.Is(err, ssession.ErrNotFound):
			m.view.ExecuteTemplate(w, nil, "home.html")
			return
		case err != nil:
			lg.Err.Printf("IndexHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		user, err := m.service.User.GetByID(session.UserId)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("IndexHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		pg := &view.Page{User: user}
		m.view.ExecuteTemplate(w, pg, "home.html")
		return
	}
}
