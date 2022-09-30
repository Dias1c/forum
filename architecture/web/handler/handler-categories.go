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

// CategoriesHandler -
func (m *MainHandler) CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("CategoriesHandler", r)

	// Allowed Methods
	switch r.Method {
	case http.MethodGet:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Logic
	switch r.Method {
	case http.MethodGet:
		pCategories, err := m.service.Category.GetAll(0, models.SqlLimitInfinity)
		if err != nil {
			lg.Err.Printf("CategoriesHandler: PostCategory.GetAll: %v\n", err)
		}

		cookie := cookies.GetSessionCookie(w, r)
		if cookie == nil {
			pg := &view.Page{Categories: pCategories}
			m.view.ExecuteTemplate(w, pg, "categories.html")
			return
		}

		session, err := m.service.Session.GetByUuid(cookie.Value)
		switch {
		case err == nil:
		case errors.Is(err, ssession.ErrExpired) || errors.Is(err, ssession.ErrNotFound):
			cookies.RemoveSessionCookie(w, r)
			pg := &view.Page{Categories: pCategories}
			m.view.ExecuteTemplate(w, pg, "categories.html")
			return
		case err != nil:
			lg.Err.Printf("CategoriesHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		user, err := m.service.User.GetByID(session.UserId)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("CategoriesHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		pg := &view.Page{Categories: pCategories, User: user}
		m.view.ExecuteTemplate(w, pg, "categories.html")
		return
	}
}
