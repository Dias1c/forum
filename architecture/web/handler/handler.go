package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"forum/architecture/service"
)

type Configs struct {
	Templates   string `templates`
	StaticFiles string `static_files`
}

type MainHandler struct {
	templates *template.Template
	service   *service.Service
}

func NewMainHandler(service *service.Service, configs *Configs) (*MainHandler, error) {
	templates, err := newTemplate(configs)
	if err != nil {
		return nil, err
	}
	return &MainHandler{
		templates: templates,
		service:   service,
	}, nil
}

func newTemplate(configs *Configs) (*template.Template, error) {
	// Gets All Templates in folder templates
	filepaths, err := filepath.Glob(configs.Templates + "/*.html")
	files, err := template.ParseFiles(filepaths...)
	if err != nil {
		return nil, fmt.Errorf("newTemplate: %w", err)
	}
	return template.Must(files, nil), nil
}

func (m *MainHandler) InitRoutes(configs *Configs) http.Handler {
	mux := http.NewServeMux()
	// HERE IS ALL ROUTES
	fsStatic := http.FileServer(http.Dir(configs.StaticFiles))
	mux.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	// AnyRoutes
	mux.HandleFunc("/test", m.TestHandler)
	return mux
}
