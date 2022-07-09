package handler

import (
	"forum/architecture/service"
	"forum/architecture/web/handler/view"
	"net/http"
)

type Configs struct {
	TemplatesDir   string `cenv:"templates"`
	StaticFilesDir string `cenv:"static_files"`
}

type MainHandler struct {
	// templates    *template.Template
	view    view.View
	service *service.Service
}

func NewMainHandler(service *service.Service, configs *Configs) (*MainHandler, error) {
	mh := &MainHandler{
		view:    *view.NewView(configs.TemplatesDir),
		service: service,
	}
	return mh, nil
}

func (m *MainHandler) InitRoutes(configs *Configs) http.Handler {
	mux := http.NewServeMux()
	// HERE IS ALL ROUTES
	fsStatic := http.FileServer(http.Dir(configs.StaticFilesDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	// AnyRoutes
	mux.HandleFunc("/debug", m.DebugHandler)

	mux.HandleFunc("/", m.IndexHandler)
	mux.Handle("/signup", m.MiddlewareMethodChecker(
		http.HandlerFunc(m.SignUpHandler),
		map[string]bool{"GET": true, "POST": true}),
	)
	mux.Handle("/signin", m.MiddlewareMethodChecker(
		http.HandlerFunc(m.SignInHandler),
		map[string]bool{"GET": true, "POST": true}),
	)
	mux.Handle("/signout", m.MiddlewareMethodChecker(
		http.HandlerFunc(m.SignOutHandler),
		map[string]bool{"GET": true}),
	)
	// mux.HandleFunc("/post/get", m.PostGet)
	mux.Handle("/post/create", m.MiddlewareMethodChecker(
		m.MiddlewareCookieChecker(http.HandlerFunc(m.PostCreateHandler)),
		map[string]bool{"GET": true, "POST": true}),
	)

	return mux
}
