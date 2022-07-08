package handler

import (
	"net/http"

	"forum/architecture/web/handler/cookies"
)

// LogInHandler -
func (m *MainHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("IndexHandler", r)
	cookies.RemoveRedirectCookie(w, r)

	switch r.Method {
	case http.MethodGet:
		m.view.ExecuteTemplate(w, nil, "home.html")
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
