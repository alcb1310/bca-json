package server

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/joho/godotenv/autoload"

	"github.com/alcb1310/bca-json/internals/database"
)

type Server struct {
	Server    *chi.Mux
	DB        database.Service
	TokenAuth *jwtauth.JWTAuth
}

func New(db database.Service) *Server {
	chiServer := chi.NewRouter()

	s := &Server{
		Server:    chiServer,
		DB:        db,
		TokenAuth: jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil),
	}

	// Middlewares
	s.Server.Use(middleware.Logger)
	s.Server.Use(middleware.CleanPath)
	s.Server.Use(middleware.StripSlashes)
	s.Server.Use(middleware.AllowContentType("application/json"))
	s.Server.Use(ContentTypeJSON)

	// Routes
	s.Server.Get("/", handleErrors(Home))



	s.Server.Group(func(r chi.Router) {
		r.Route("/api/v2", func(r chi.Router) {
			r.Post("/companies", handleErrors(s.CreateCompany))
			r.Post("/login", handleErrors(s.Login))
		})
	})

	return s
}
