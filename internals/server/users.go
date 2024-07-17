package server

import (
	"encoding/json"
	"net/http"

	"github.com/alcb1310/bca-json/internals/types"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) error {
    var companyUUID uuid.UUID
     _, claims, _ := jwtauth.FromContext(r.Context())

    companyUUID, _ = uuid.Parse(claims["company"].(string))

    users, err := s.DB.GetUsers(companyUUID)
    if err != nil {
        return &types.BCAError{
            Code:    http.StatusInternalServerError,
            Message: err,
        }
    }

    w.WriteHeader(http.StatusOK)
    return json.NewEncoder(w).Encode(users)
}
