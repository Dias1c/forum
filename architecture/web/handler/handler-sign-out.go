package handler

import (
	"net/http"

	"github.com/Dias1c/forum/architecture/web/handler/cookies"
)

// SignOutHandler -
func (m *MainHandler) SignOutHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("SignOutHandler", r)

	switch r.Method {
	case http.MethodGet:
		cookies.RemoveSessionCookie(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
