package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"

	"github.com/alcb1310/bca-json/internals/database"
)

type Server struct {
	Server    *chi.Mux
	DB        database.Service
	TokenAuth *jwtauth.JWTAuth
}

func New(db database.Service, secret string) *Server {
	chiServer := chi.NewRouter()

	s := &Server{
		Server:    chiServer,
		DB:        db,
		TokenAuth: jwtauth.New("HS256", []byte(secret), nil),
	}

	// Middlewares
	s.Server.Use(middleware.Logger)
	s.Server.Use(middleware.CleanPath)
	s.Server.Use(middleware.StripSlashes)
	s.Server.Use(middleware.AllowContentType("application/json"))
	s.Server.Use(ContentTypeJSON)
	s.Server.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	// Routes
	s.Server.Get("/", handleErrors(Home))

	s.Server.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(s.TokenAuth))
		r.Use(jwtauth.Authenticator(s.TokenAuth))

		r.Route("/api/v2/bca", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", handleErrors(s.GetUsers))
				r.Post("/", handleErrors(s.CreateUser))
				r.Get("/me", handleErrors(s.GetCurrentUser))
				r.Get("/{userID}", handleErrors(s.GetUserByID))
				r.Delete("/{userID}", handleErrors(s.DeleteUser))
				r.Put("/{userID}", handleErrors(s.UpdateUser))
			})

            r.Route("/projects", func(r chi.Router) {
                r.Get("/", handleErrors(s.GetProjects))
                r.Post("/", handleErrors(s.CreateProject))
                r.Get("/{projectID}", handleErrors(s.GetProjectByID))
            })
		})
	})

	s.Server.Group(func(r chi.Router) {
		r.Route("/api/v2", func(r chi.Router) {
			r.Post("/companies", handleErrors(s.CreateCompany))
			r.Post("/login", handleErrors(s.Login))
		})
	})

	return s
}
