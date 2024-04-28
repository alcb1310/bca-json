package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/alcb1310/bca-json/internal/handlers"
	"github.com/alcb1310/bca-json/internal/utils"
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
	s.Router.Use(middleware.Logger)

	s.Router.Get("/foo", utils.Make(handlers.HandleFoo))

	s.Router.NotFound(handlers.HandleNotFound)
}
