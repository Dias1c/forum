package handler

import (
	"fmt"
	"forum/architecture/models"
	"log"
	"net/http"
	"strings"
)

// TestHandler - Handle for Testing
func (m *MainHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("SignUpHandler", r)
	err := m.debugRefreshTemplates()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err = m.templates.ExecuteTemplate(w, "bootstrap", nil)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err = r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "r.ParseForm: %v\n", err)
			return
		}

		newUser := &models.User{
			Nickname: r.FormValue("nickname"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		userId, err := m.service.User.Create(newUser)
		if err == nil {
			// http.StatusPermanentRedirect
			http.Redirect(w, r, "/login", userId)
			return
		}

		if strings.HasPrefix(err.Error(), "client: ") {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			err = m.templates.ExecuteTemplate(w, "bootstrap", nil)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err) //? debug
		fmt.Fprintln(w, "internal server error")

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %v not allowed\n", r.Method)
		return
	}
}
