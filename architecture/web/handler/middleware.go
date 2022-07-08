package handler

import (
	"fmt"
	"forum/architecture/web/handler/cookies"
	"forum/architecture/web/handler/view"
	"log"
	"net/http"
	"time"
)

// MiddlewareCookieChecker - NOT FINISHED
func (m *MainHandler) MiddlewareCookieChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		debugLogHandler("MiddlewareCookieChecker", r)

		cookie := cookies.GetSessionCookie(w, r)
		if cookie == nil {
			cookies.AddRedirectCookie(w, r.URL.Path)
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}

		session, err := m.service.Session.GetByUuid(cookie.Value)
		switch {
		// case errors.Is(err, ssession.ErrNotFound):
		case err != nil:
			log.Printf("MiddlewareCookieChecker: m.service.Session.GetByUuid: %v\n", err)
			pg := &view.Page{Error: fmt.Errorf("something wrong, maybe try again later")}
			w.WriteHeader(http.StatusInternalServerError)
			m.view.ExecuteTemplate(w, pg, "error.html")
			return
		default:
			duration := time.Until(session.ExpiredAt)
			if duration.Seconds() <= 0 {
				cookies.AddRedirectCookie(w, r.URL.Path)
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
