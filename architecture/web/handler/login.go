package handler

import (
	"fmt"
	"log"
	"net/http"
)

// LogInHandler -
func (m *MainHandler) LogInHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("LogInHandler", r)

	switch r.Method {
	case http.MethodGet:
		// tmpl, err := m.templates.ParseFiles("web/templates/bootstrap.html", "web/templates/navbar.html", "web/templates/login.html")
		tmpl, err := m.newView("login.html")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "bootstrap", nil)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "r.ParseForm: %v\n", err)
			return
		}

		// newUser := &models.User{
		// 	Nickname: r.FormValue("login"),
		// 	Password: r.FormValue("password"),
		// }

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %v not allowed\n", r.Method)
	}
}
