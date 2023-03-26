package server

import (
	"context"
	"log"
	"net/http"
)

type Config struct {
	Port int16
}

type Server struct {
	server *http.Server
}

func New(router http.Handler) *Server {
	srv := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	return &Server{
		server: srv,
	}
}

func (s *Server) ServeHTTPHandler() {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[server] failed to listen and serve: %s", err.Error())
	}
}

func (s *Server) ShutdownHTTPHandler(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("[server] failed to shutdown: %s", err.Error())
	}
}
