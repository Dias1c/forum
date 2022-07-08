package cookies

import (
	"net/http"
)

const (
	CookieSessionName = "session"
)

// AddSessionCookie - sets session cookie
func AddSessionCookie(w http.ResponseWriter, uuid string, durationSec int) {
	http.SetCookie(w,
		&http.Cookie{
			Name:   CookieSessionName,
			Value:  uuid,
			MaxAge: durationSec,
		},
	)
}

// GetRedirectCookie - returns session cookie
//
// Cookie Value is uuid
func GetSessionCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	return getCookieByName(w, r, CookieSessionName)
}

// CleanSessionCookie - removes cookie by setting maxAge -1
func CleanSessionCookie(w http.ResponseWriter, r *http.Request) {
	removeCookieByName(w, r, CookieSessionName)
}
