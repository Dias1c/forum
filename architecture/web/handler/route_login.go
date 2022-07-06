package handler

import (
	"fmt"
	"net/http"
)

// LogInHandler -
func (m *MainHandler) LogInHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("LogInHandler", r)

	switch r.Method {
	case http.MethodGet:
		m.executeTemplate(w, nil, "login.html")
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "r.ParseForm: %v\n", err)
			return
		}

		// usr := &models.User{
		// 	Nickname: r.FormValue("login"),
		// 	Password: r.FormValue("password"),
		// }

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %v not allowed\n", r.Method)
	}
}
