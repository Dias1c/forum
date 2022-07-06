package handler

import (
	"fmt"
	"log"
	"net/http"
)

// TestHandler - Handle for Testing
func (m *MainHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("TestHandler", r)
	fmt.Fprintf(w, "Current state:\n")
	fmt.Fprintf(w, "Cookies count: %v; Cookies: %q;\n", len(r.Cookies()), r.Cookies())
	fmt.Fprintf(w, "Ready Endpoints:\n")
	fmt.Fprintf(w, "Main:      /\n")
	fmt.Fprintf(w, "SignUp:    /signup\n")
	fmt.Fprintf(w, "LogIn:     /login\n")
}

//? debugRefreshTemplates -
// func (m *MainHandler) debugRefreshTemplates() error {
// 	templates, err := m.newTemplate()
// 	if err != nil {
// 		return fmt.Errorf("debugRefreshTemplates: %w", err)
// 	}
// 	m.templates = templates
// 	return nil
// }

//? debugLogHandler -
func debugLogHandler(fName string, r *http.Request) {
	log.Printf("%-20v | %-7v | %-20v \n", r.URL, r.Method, fName)
}
