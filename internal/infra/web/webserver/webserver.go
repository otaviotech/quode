package webserver

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type WebserverHandlerInterface interface {
	GetMethod() string
	GetPath() string
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Webserver struct {
	Router   chi.Router
	Handlers []WebserverHandlerInterface
	Port     string
}

func NewWebserver(port string) *Webserver {
	return &Webserver{
		Router:   chi.NewRouter(),
		Handlers: []WebserverHandlerInterface{},
		Port:     ":" + port,
	}
}

func (w *Webserver) AddHandler(handler WebserverHandlerInterface) {
	w.Handlers = append(w.Handlers, handler)
}

func (w *Webserver) setupHandlers() {
	for _, handler := range w.Handlers {
		w.Router.Method(handler.GetMethod(), handler.GetPath(), handler)
	}
}

func (w *Webserver) Start() {
	w.Router.Use(middleware.Logger)

	w.setupHandlers()

	log.Default().Printf("Webserver started on port %s\n", w.Port)
	http.ListenAndServe(w.Port, w.Router)
}
