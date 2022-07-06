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
	User  models.User
	Users []models.User
	Post  models.Post
	Posts []models.Post
	// Comments           []models.Comment
	IsErr     bool
	ErrMsg    string
	ErrStatus int
}

func NewMainHandler(service *service.Service, configs *Configs) (*MainHandler, error) {
	mh := &MainHandler{
		templatesDir: configs.TemplatesDir,
		service:      service,
	}
	// templates, err := mh.newTemplate()
	// if err != nil {
	// 	return nil, err
	// }
	// mh.templates = templates
	return mh, nil
}

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

func (m *MainHandler) executeTemplate(w http.ResponseWriter, pg interface{}, names string) {
	tmpl, err := m.newView(names)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "bootstrap", pg)
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
