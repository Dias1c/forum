package handler

import (
	"fmt"
	"forum/architecture/web/handler/view"
	"log"
	"net/http"
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
		return
	}

	userId := iUserId.(int64)
	user, _ := m.service.User.GetByID(userId)
	switch r.Method {
	case http.MethodGet:
		pg := &view.Page{User: user, Error: fmt.Errorf("hi")}
		m.view.ExecuteTemplate(w, pg, "post-create.html")
	case http.MethodPost:
		fmt.Fprint(w, "%v", user)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
