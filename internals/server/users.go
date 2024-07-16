package server

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) error {
    log.Println("GetUsers")
     _, claims, _ := jwtauth.FromContext(r.Context())
    slog.Info("GetUsers", "claims", claims)
    return json.NewEncoder(w).Encode(claims)
}
