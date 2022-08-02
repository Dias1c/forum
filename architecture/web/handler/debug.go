package handler

import (
	"fmt"
	"forum/internal/lg"
	"net/http"
)

// DebugHandler - Handle for Testing
func (m *MainHandler) DebugHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("TestHandler", r)
	fmt.Fprintf(w, "Current state:\n")
	fmt.Fprintf(w, "Cookies count: %v; Cookies: %q;\n", len(r.Cookies()), r.Cookies())
	fmt.Fprintf(w, "Ready Endpoints:\n")
	fmt.Fprintf(w, "Your Current path:      %s%s\n", r.URL.Host, r.URL.Path)
	fmt.Fprintf(w, "Main:      %s/\n", r.URL.Host)
	fmt.Fprintf(w, "SignUp:    %s/signup\n", r.URL.Host)
	fmt.Fprintf(w, "SignIn:    %s/signin\n", r.URL.Host)
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
	lg.Info.Printf("%-30v | %-7v | %-30v \n", r.URL, r.Method, fName)
}
