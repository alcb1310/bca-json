package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/alcb1310/bca-json/internal/database"
	"github.com/alcb1310/bca-json/internal/utils"
)

type Server struct {
	Router *chi.Mux
	DB     database.Database
}

func NewServer(db database.Database) *Server {
	return &Server{
		Router: chi.NewRouter(),
		DB:     db,
	}
}

func (s *Server) MountHandlers() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	s.Router.Get("/foo", utils.Make(s.HandleFoo))
	s.Router.Post("/register", utils.Make(s.HandleRegister))

	s.Router.NotFound(utils.Make(NotFound))
	s.Router.MethodNotAllowed(utils.Make(MethodNotAllowed))
}
