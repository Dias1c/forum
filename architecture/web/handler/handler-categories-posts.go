package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Dias1c/forum/architecture/models"
	"github.com/Dias1c/forum/architecture/web/handler/cookies"
	"github.com/Dias1c/forum/architecture/web/handler/view"
	"github.com/Dias1c/forum/internal/lg"

	ssession "github.com/Dias1c/forum/architecture/service/session"
)

// CategoriesPostsHandler -
func (m *MainHandler) CategoriesPostsHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("CategoriesPostsHandler", r)
	// TODO: Add Categories Posts Page
	// Allowed Methods
	switch r.Method {
	case http.MethodGet:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	switch r.Method {
	case http.MethodGet:
		value := r.URL.Query().Get("categories")
		categoryNames := strings.Fields(value)
		if len(categoryNames) > 5 {
			http.Error(w, "Max Category names is 5", http.StatusBadRequest)
			return
		}

		categories, err := m.service.Category.GetByNames(categoryNames)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("CategoriesPostsHandler: Category.GetByNames: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		var infoMsg error
		if len(categories) != len(categoryNames) {
			infoMsg = fmt.Errorf("Looking for only contaned categories")
		}

		catIDs := make([]int64, len(categories))
		for i, v := range categories {
			catIDs[i] = v.Id
		}

		postIDs, err := m.service.Category.GetPostIDsContainedCatIDs(catIDs, 0, models.SqlLimitInfinity)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("CategoriesPostsHandler: Category.GetPostIDsContainedCatIDs: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		// TODO: Rename Ids -> IDs
		posts, err := m.service.Post.GetByIds(postIDs)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("CategoriesPostsHandler: Post.GetByIds: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		pg := &view.Page{Posts: posts, Categories: categories, Info: infoMsg}
		cookie := cookies.GetSessionCookie(w, r)
		if cookie == nil {
			err = m.service.FillPosts(posts, 0)
			if err != nil {
				lg.Err.Printf("CategoriesPostsHandler: FillPosts: %v\n", err)
			}
			m.view.ExecuteTemplate(w, pg, "categories-posts.html")
			return
		}

		session, err := m.service.Session.GetByUuid(cookie.Value)
		switch {
		case err == nil:
		case errors.Is(err, ssession.ErrExpired) || errors.Is(err, ssession.ErrNotFound):
			cookies.RemoveSessionCookie(w, r)
			err = m.service.FillPosts(posts, 0)
			if err != nil {
				lg.Err.Printf("CategoriesPostsHandler: FillPosts: %v\n", err)
			}
			m.view.ExecuteTemplate(w, pg, "categories-posts.html")
			return
		case err != nil:
			lg.Err.Printf("CategoriesPostsHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		user, err := m.service.User.GetByID(session.UserId)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("CategoriesPostsHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		err = m.service.FillPosts(posts, user.Id)
		if err != nil {
			lg.Err.Printf("CategoriesPostsHandler: FillPosts: %v\n", err)
		}
		pg.User = user
		m.view.ExecuteTemplate(w, pg, "categories-posts.html")
		return
	}
}
