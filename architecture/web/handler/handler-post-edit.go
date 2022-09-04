package handler

import (
	"errors"
	"fmt"
	"forum/architecture/web/handler/cookies"
	"forum/architecture/web/handler/view"
	"forum/internal/lg"
	"net/http"
	"strconv"

	spost "forum/architecture/service/post"
	suser "forum/architecture/service/user"
)

// PostCreateHandler -
func (m *MainHandler) PostEditHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("PostEditHandler", r)

	// Allowed Methods
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	iUserId := r.Context().Value("UserId")
	if iUserId == nil {
		lg.Err.Println("PostEditHandler: r.Context().Value(\"UserId\") is nil")
		pg := &view.Page{Error: fmt.Errorf("internal server error, maybe try again later")}
		w.WriteHeader(http.StatusInternalServerError)
		m.view.ExecuteTemplate(w, pg, "post-edit.html")
		return
	}

	userId := iUserId.(int64)
	user, err := m.service.User.GetByID(userId) // TODO: Проверь это
	switch {
	case err == nil:
	case errors.Is(err, suser.ErrNotFound):
		cookies.RemoveSessionCookie(w, r)
		cookies.AddRedirectCookie(w, r.RequestURI)
		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
	case err != nil:
		lg.Err.Printf("PostEditHandler: m.service.User.GetByID: %v\n", err)
		pg := &view.Page{Error: fmt.Errorf("internal server error, maybe try again later")}
		w.WriteHeader(http.StatusInternalServerError)
		m.view.ExecuteTemplate(w, pg, "alert.html") // TODO: Custom Error Page
	}

	switch r.Method {
	case http.MethodGet:
		strPostId := r.URL.Query().Get("id")
		postId, err := strconv.ParseInt(strPostId, 10, 64)
		if err != nil || postId < 1 {
			http.Error(w, "Invalid query id", http.StatusBadRequest)
			return
		}
		post, err := m.service.Post.GetByID(postId)
		switch {
		case err == nil:
		case errors.Is(err, spost.ErrNotFound):
			// TODO: error page
			http.Error(w, "Post Not Found", http.StatusNotFound)
			return
		}

		if post.UserId != user.Id {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		categories, err := m.service.PostCategory.GetByPostID(post.Id)
		switch {
		case err == nil:
			post.WCategories = categories
		default:
			lg.Err.Printf("PostEditHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		pg := &view.Page{User: user, Post: post}
		m.view.ExecuteTemplate(w, pg, "post-edit.html")
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
