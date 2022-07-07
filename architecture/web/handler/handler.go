package handler

import (
	"fmt"
	"forum/architecture/models"
	"forum/architecture/service"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type Configs struct {
	TemplatesDir   string `templates`
	StaticFilesDir string `static_files`
}

type MainHandler struct {
	// templates    *template.Template
	templatesDir string
	service      *service.Service
}

type page struct {
	User  *models.User
	Users *[]models.User
	Post  *models.Post
	Posts *[]models.Post
	// Comments           []models.Comment
	Error error
	Warn  error
}

func NewMainHandler(service *service.Service, configs *Configs) (*MainHandler, error) {
	mh := &MainHandler{
		templatesDir: configs.TemplatesDir,
		service:      service,
	}
	return mh, nil
}

func (m *MainHandler) InitRoutes(configs *Configs) http.Handler {
	mux := http.NewServeMux()
	// HERE IS ALL ROUTES
	fsStatic := http.FileServer(http.Dir(configs.StaticFilesDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	// AnyRoutes
	mux.HandleFunc("/", m.TestHandler)
	mux.HandleFunc("/signup", m.SignUpHandler)
	mux.HandleFunc("/login", m.LogInHandler)
	return mux
}

func (m *MainHandler) newView(names ...string) (*template.Template, error) {
	paths := []string{m.templatesDir + "/bootstrap.html", m.templatesDir + "/navbar.html"}
	for _, name := range names {
		paths = append(paths, m.templatesDir+"/"+name)
	}

	t, err := template.ParseFiles(paths...)
	if err != nil {
		return nil, fmt.Errorf("template.ParseFiles: %w", err)
	}
	return t, nil
}

func (m *MainHandler) executeTemplate(w http.ResponseWriter, pg interface{}, names ...string) {
	tmpl, err := m.newView(names...)
	if err != nil {
		log.Printf("m.newView: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "bootstrap", pg)
	if err != nil {
		log.Printf("tmpl.ExecuteTemplate: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UNUSING

// newTemplate - returns combined files template which in path m.templatesDir
func (m *MainHandler) newTemplate() (*template.Template, error) {
	// Gets All Templates in folder templates
	filepaths, err := filepath.Glob(m.templatesDir + "/*.html")
	if err != nil {
		return nil, fmt.Errorf("newTemplate: %w", err)
	}
	files, err := template.ParseFiles(filepaths...)
	if err != nil {
		return nil, fmt.Errorf("newTemplate: %w", err)
	}
	return template.Must(files, nil), nil
}
