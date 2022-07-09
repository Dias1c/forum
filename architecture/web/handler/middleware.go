package handler

import (
	"context"
	"errors"
	ssession "forum/architecture/service/session"
	"forum/architecture/web/handler/cookies"
	"log"
	"net/http"
)

func (m *MainHandler) MiddlewareMethodChecker(next http.Handler, allowedMthods map[string]bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		debugLogHandler("MiddlewareMethodChecker", r)
		if !allowedMthods[r.Method] {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// MiddlewareCookieChecker - NOT FINISHED
func (m *MainHandler) MiddlewareCookieChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		debugLogHandler("MiddlewareCookieChecker", r)

		cookie := cookies.GetSessionCookie(w, r)
		if cookie == nil {
			cookies.AddRedirectCookie(w, r.URL.Path)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		session, err := m.service.Session.GetByUuid(cookie.Value)
		switch {
		case err == nil:
		case errors.Is(err, ssession.ErrExpired) || errors.Is(err, ssession.ErrNotFound):
			cookies.AddRedirectCookie(w, r.URL.Path)
			cookies.RemoveSessionCookie(w, r)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		case err != nil:
			log.Printf("MiddlewareCookieChecker: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "UserId", session.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
