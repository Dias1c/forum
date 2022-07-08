package cookies

import (
	"errors"
	"log"
	"net/http"
)

const (
	CookieRedirectName = "redirect_to"
)

// AddRedirectCookie - sets redirect cookie if field redirectTo is not empty
func AddRedirectCookie(w http.ResponseWriter, redirectTo string) {
	if redirectTo == "" {
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:   CookieRedirectName,
			Value:  redirectTo,
			MaxAge: 3600,
		},
	)
}

// GetRedirectCookie - returns redirect cookie
func GetRedirectCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie(CookieRedirectName)
	switch {
	case errors.Is(err, http.ErrNoCookie):
	case err != nil:
		log.Printf("GetRedirectCookie: r.Cookie: %v", err)
	case cookie != nil:
		return cookie
	}
	return nil
}

// CleanRedirectCookie - removes cookie by setting maxAge -1
func CleanRedirectCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(CookieRedirectName)
	switch {
	case errors.Is(err, http.ErrNoCookie):
	case err != nil:
		log.Printf("CleanRedirectCookie: r.Cookie: %v", err)
	case cookie != nil:
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
}
