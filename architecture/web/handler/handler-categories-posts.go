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
		// TODO: 1. Get Array of catNames (Max 5)
		// TODO: 2. Add dropdown menu category posts in navbar
		names := []string{"dc", "marvel"}
		categories, err := m.service.Category.GetByNames(names)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("CategoriesPostsHandler: Category.GetByNames: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
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

		cookie := cookies.GetSessionCookie(w, r)
		if cookie == nil {
			err = m.service.FillPosts(posts, 0)
			if err != nil {
				lg.Err.Printf("CategoriesPostsHandler: FillPosts: %v\n", err)
			}
			pg := &view.Page{Posts: posts, Categories: categories}
			m.view.ExecuteTemplate(w, pg, "categories-posts.html") // TODO: Replace PG
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
			pg := &view.Page{Posts: posts, Categories: categories}
			m.view.ExecuteTemplate(w, pg, "categories-posts.html") // TODO: Replace PG
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
		pg := &view.Page{Posts: posts, User: user, Categories: categories}
		m.view.ExecuteTemplate(w, pg, "categories-posts.html") // TODO: Replace PG
		return
	}
}
