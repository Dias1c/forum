package view

import (
	"net/http"

	"github.com/Dias1c/forum/internal/lg"
)

func NewView(templatesDir string) *View {
	return &View{templatesDir: templatesDir}
}

func (v *View) ExecuteTemplate(w http.ResponseWriter, pg interface{}, names ...string) {
	tmpl, err := v.getTemplate(names...)
	if err != nil {
		lg.Err.Printf("m.newView: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "bootstrap", pg)
	if err != nil {
		lg.Err.Printf("tmpl.ExecuteTemplate: %v", err)
		return
	}
}
