package handler

import (
	"fmt"
	"log"
	"net/http"
)

// PostCreateHandler -
func (m *MainHandler) PostCreateHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("PostCreateHandler", r)

	iUserId := r.Context().Value("UserId")
	if iUserId == nil {
		log.Println("PostCreateHandler: r.Context().Value(\"UserId\") is nil")
		return
	}

	userId := iUserId.(int64)
	user, _ := m.service.User.GetByID(userId)
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "%v", user)
	case http.MethodPost:
		fmt.Fprint(w, "%v", user)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
