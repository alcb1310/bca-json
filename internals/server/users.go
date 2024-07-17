package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alcb1310/bca-json/internals/types"
	"github.com/go-chi/chi/v5"
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

func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request) error {
    id := chi.URLParam(r, "userID")

    parsedID, err := uuid.Parse(id)
    if err!=nil {
        return &types.BCAError{
            Code: http.StatusBadRequest,
            Message: err,
        }
    }

    user, err := s.DB.GetUserByID(parsedID)
    if err != nil {
        if err == sql.ErrNoRows {
            return &types.BCAError{
                Code:    http.StatusNotFound,
                Message: errors.New("User not found"),
            }
        }
        return &types.BCAError{
            Code:    http.StatusInternalServerError,
            Message: err,
        }
    }

    w.WriteHeader(http.StatusOK)
    return json.NewEncoder(w).Encode(user)
}
