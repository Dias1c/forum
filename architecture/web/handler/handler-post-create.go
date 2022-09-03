package handler

import (
	"errors"
	"fmt"
	"forum/architecture/models"
	"forum/architecture/web/handler/view"
	"forum/internal/lg"
	"net/http"
	"strings"

	spost "forum/architecture/service/post"
	scategory "forum/architecture/service/post_category"
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
		lg.Err.Println("PostCreateHandler: r.Context().Value(\"UserId\") is nil")
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

		post := &models.Post{
			Title:   r.FormValue("title"),
			Content: r.FormValue("content"),
			UserId:  userId,
		}
		_, err := m.service.Post.Create(post)
		switch {
		case err == nil:
		case errors.Is(err, spost.ErrInvalidTitleLength):
			pg := &view.Page{Error: fmt.Errorf("invalid length of title")}
			m.view.ExecuteTemplate(w, pg, "post-create.html")
			return
		default:
			lg.Err.Printf("PostCreateHandler: m.service.Post.Create: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later: %s", err)}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "post-create.html")
			return
		}

		catNames := strings.Fields(r.Form.Get("categories"))
		err = m.service.PostCategory.AddToPostByNames(catNames, post.Id)
		switch {
		case err == nil:
		case errors.Is(err, scategory.ErrCategoryLimitForPost):
			err = m.service.Post.DeleteByID(post.Id)
			if err != nil {
				lg.Err.Println("PostCreateHandler: m.service.Post.DeleteByID: %w", err)
			}

			pg := &view.Page{Warn: fmt.Errorf("post not created, invalid categies count, category limit = %v", models.MaxCategoryLimitForPost)}
			w.WriteHeader(http.StatusBadRequest)
			m.view.ExecuteTemplate(w, pg, "post-create.html")
			return
		default:
			err = m.service.Post.DeleteByID(post.Id)
			if err != nil {
				lg.Err.Println("PostCreateHandler: m.service.Post.DeleteByID: %w", err)
			}

			lg.Err.Printf("PostCreateHandler:  m.service.Category.AddToPostByNames: %s", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later: %s", err)}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "post-create.html")
			return
		}
		// Create categories
		http.Redirect(w, r, fmt.Sprintf("/post/get?id=%v", post.Id), http.StatusSeeOther)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
