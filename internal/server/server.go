package server

import (
	"github.com/go-chi/chi/v5"

	"github.com/alcb1310/bca-json/internal/handlers"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (s *Server) MountHandlers() {
	s.Router.Get("/foo", handlers.HandleFoo)
}
