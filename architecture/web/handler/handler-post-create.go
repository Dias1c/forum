package handler

import (
	"errors"
	"fmt"
	"forum/architecture/models"
	"forum/architecture/web/handler/view"
	"log"
	"net/http"
	"strings"

	spost "forum/architecture/service/post"
)

// PostCreateHandler -
func (m *MainHandler) PostCreateHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("PostCreateHandler", r)

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
		log.Println("PostCreateHandler: r.Context().Value(\"UserId\") is nil")
		pg := &view.Page{Error: fmt.Errorf("internal server error, maybe try again later")}
		w.WriteHeader(http.StatusInternalServerError)
		m.view.ExecuteTemplate(w, pg, "post-create.html")
		return
	}

	userId := iUserId.(int64)
	user, _ := m.service.User.GetByID(userId)
	switch r.Method {
	case http.MethodGet:
		pg := &view.Page{User: user}
		m.view.ExecuteTemplate(w, pg, "post-create.html")
		return
	case http.MethodPost:
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Printf("categories: %+v\n", strings.Split(r.Form.Get("categories"), " "))

		post := &models.Post{
			Title:   r.FormValue("title"),
			Content: r.FormValue("content"),
			UserId:  userId,
		}
		postId, err := m.service.Post.Create(post)
		switch {
		case err == nil:
		case errors.Is(err, spost.ErrInvalidTitleLength):
			pg := &view.Page{Error: fmt.Errorf("invalid length of title")}
			m.view.ExecuteTemplate(w, pg, "post-create.html")
			return
		default:
			log.Printf("ERROR: PostCreateHandler: m.service.Post.Create: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later: %s", err)}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "post-create.html")
			return
		}

		// Create categories
		log.Printf("LOGIC: %d %s", postId, err)
		pg := &view.Page{Error: fmt.Errorf("logic not finished")}
		m.view.ExecuteTemplate(w, pg, "post-create.html")
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
