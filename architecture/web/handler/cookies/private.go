package cookies

import (
	"errors"
	"forum/internal/lg"
	"net/http"
)

// getCookieByName - returns cookie
func getCookieByName(w http.ResponseWriter, r *http.Request, name string) *http.Cookie {
	cookie, err := r.Cookie(name)
	switch {
	case errors.Is(err, http.ErrNoCookie):
	case err != nil:
		lg.Err.Printf("GetRedirectCookie: r.Cookie: %v", err)
	case cookie != nil:
		return cookie
	}
	return nil
}

// removeCookieByName - remove cookie by setting maxAge -1
func removeCookieByName(w http.ResponseWriter, r *http.Request, name string) {
	cookie, err := r.Cookie(name)
	switch {
	case errors.Is(err, http.ErrNoCookie):
	case err != nil:
		lg.Err.Printf("removeCookieByName: r.Cookie: %v", err)
	case cookie != nil:
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
}
