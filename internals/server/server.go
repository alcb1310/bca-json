package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/alcb1310/bca-json/internals/database"
)

type Server struct {
	Server *chi.Mux
	DB     database.Service
}

func New(db database.Service) *Server {
	chiServer := chi.NewRouter()

	s := &Server{
		Server: chiServer,
		DB:     db,
	}

	// Middlewares
	s.Server.Use(middleware.Logger)
	s.Server.Use(middleware.CleanPath)
	s.Server.Use(middleware.StripSlashes)
	s.Server.Use(middleware.AllowContentType("application/json"))
	s.Server.Use(ContentTypeJSON)

	// Routes
	s.Server.Get("/", handleErrors(Home))

    s.Server.Route("/api/v2", func(r chi.Router) {
        r.Post("/companies", handleErrors(s.CreateCompany))
    })

	return s
}
