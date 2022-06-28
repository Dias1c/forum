package handler

import (
	"fmt"
	"log"
	"net/http"
)

// TestHandler - Handle for Testing
func (m *MainHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	err := m.debugRefreshTemplates()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// err = v.templates.ExecuteTemplate(w, "pg-index", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-signup", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-login", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-question", nil)
	err = m.templates.ExecuteTemplate(w, "bootstrap", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-question-create", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-tags", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-user", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-users", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//? debugRefreshTemplates -
func (m *MainHandler) debugRefreshTemplates() error {
	templates, err := newTemplate(&Configs{Templates: "web/templates", StaticFiles: "web/static"})
	if err != nil {
		return fmt.Errorf("debugRefreshTemplates: %w", err)
	}
	m.templates = templates
	return nil
}
