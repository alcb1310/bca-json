package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
    Server *chi.Mux
}

func New() *Server {
    chiServer := chi.NewRouter()

    s := &Server{
        Server: chiServer,
    }

    // Middlewares
    s.Server.Use(middleware.Logger)
    s.Server.Use(middleware.CleanPath)
    s.Server.Use(middleware.StripSlashes)
    s.Server.Use(middleware.AllowContentType("application/json"))
    s.Server.Use(ContentTypeJSON)

    // Routes
    s.Server.Get("/", handleErrors(Home))

    return s
}
