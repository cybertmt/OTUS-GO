package internalhttp

import (
	"context"
	"github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/app"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

type Server struct {
	host   string
	port   string
	logger app.Logger
	server *http.Server
}

type Application interface { // TODO
}

func NewServer(logger app.Logger, app *app.App, host, port string) *Server {
	server := &Server{
		host:   host,
		port:   port,
		logger: logger,
		server: nil,
	}

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: loggingMiddleware(NewRouter(app), logger),
	}

	server.server = httpServer

	return server
}

func NewRouter(app *app.App) http.Handler {
	handlers := NewServerHandlers(app)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HelloWorld).Methods("GET")
	r.HandleFunc("/events", handlers.CreateEvent).Methods("POST")
	r.HandleFunc("/events/{id}", handlers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", handlers.DeleteEvent).Methods("DELETE")
	r.HandleFunc("/events", handlers.ListEvents).Methods("GET")

	return r
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("HTTP server listen and serve %s:%s", s.host, s.port)
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
