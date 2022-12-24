// Package router implements the router interface abstractions and server mux
package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/blevels/weatherAPI/infrastructure/logger"
	"github.com/gorilla/mux"
)

type Mux struct {
	router *mux.Router
}

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewMux() *Mux {
	return &Mux{
		router: mux.NewRouter(),
	}
}

func (m *Mux) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.router.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (m *Mux) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.router.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (m *Mux) SERVE(port string) {
	httpServer := &http.Server{
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      m.router,
	}

	logger.NewLogrus().Infof("The HTTP server has started and is running on port: %s", port)

	log.Fatal(httpServer.ListenAndServe())
}
