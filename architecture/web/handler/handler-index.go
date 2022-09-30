package handler

import (
	"errors"
	"net/http"

	"github.com/Dias1c/forum/architecture/models"
	"github.com/Dias1c/forum/architecture/web/handler/cookies"
	"github.com/Dias1c/forum/architecture/web/handler/view"
	"github.com/Dias1c/forum/internal/lg"

	ssession "github.com/Dias1c/forum/architecture/service/session"
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
		posts, err := m.service.Post.GetAll(0, models.SqlLimitInfinity)
		if err != nil {
			lg.Err.Printf("IndexHandler: Post.GetAll: %v\n", err)
		}

		cookie := cookies.GetSessionCookie(w, r)
		if cookie == nil {
			err = m.service.FillPosts(posts, 0)
			if err != nil {
				lg.Err.Printf("IndexHandler: FillPosts: %v\n", err)
			}
			pg := &view.Page{Posts: posts}
			m.view.ExecuteTemplate(w, pg, "home.html")
			return
		}

		session, err := m.service.Session.GetByUuid(cookie.Value)
		switch {
		case err == nil:
		case errors.Is(err, ssession.ErrExpired) || errors.Is(err, ssession.ErrNotFound):
			err = m.service.FillPosts(posts, 0)
			if err != nil {
				lg.Err.Printf("IndexHandler: FillPosts: %v\n", err)
			}
			pg := &view.Page{Posts: posts}
			m.view.ExecuteTemplate(w, pg, "home.html")
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

		err = m.service.FillPosts(posts, user.Id)
		if err != nil {
			lg.Err.Printf("IndexHandler: FillPosts: %v\n", err)
		}
		pg := &view.Page{Posts: posts, User: user}
		m.view.ExecuteTemplate(w, pg, "home.html")
		return
	}
}
