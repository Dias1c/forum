package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dias1c/forum/architecture/web/handler/cookies"
	"github.com/Dias1c/forum/architecture/web/handler/view"
	"github.com/Dias1c/forum/internal/lg"

	spost "github.com/Dias1c/forum/architecture/service/post"
	suser "github.com/Dias1c/forum/architecture/service/user"
)

// PostDeleteHandler -
func (m *MainHandler) PostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("PostDeleteHandler", r)

	// Allowed Methods
	switch r.Method {
	case http.MethodGet:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	iUserId := r.Context().Value("UserId")
	if iUserId == nil {
		lg.Err.Println("PostDeleteHandler: r.Context().Value(\"UserId\") is nil")
		pg := &view.Page{Error: fmt.Errorf("internal server error, maybe try again later")}
		w.WriteHeader(http.StatusInternalServerError)
		m.view.ExecuteTemplate(w, pg, "alert.html") // TODO: Custom Error Page
		return
	}

	userId := iUserId.(int64)
	user, err := m.service.User.GetByID(userId)
	switch {
	case err == nil:
	case errors.Is(err, suser.ErrNotFound):
		cookies.RemoveSessionCookie(w, r)
		cookies.AddRedirectCookie(w, r.RequestURI)
		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
	case err != nil:
		lg.Err.Printf("PostDeleteHandler: m.service.User.GetByID: %v\n", err)
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

		err = m.service.Post.DeleteByID(post.Id)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("PostDeleteHandler: m.service.Post.DeleteByID: %v\n", err)
			pg := &view.Page{Error: fmt.Errorf("internal server error, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "alert.html") // TODO: Custom Error Page
			return
		}

		http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
