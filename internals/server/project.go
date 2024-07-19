package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/internals/validation"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

func (s *Server) CreateProject(w http.ResponseWriter, r *http.Request) error {
    if r.Body == nil {
        return &types.BCAError{
            Code:    http.StatusBadRequest,
            Message: errors.New("Missing Body"),
        }
    }
    _, claims, _ := jwtauth.FromContext(r.Context())
    companyUUID := uuid.MustParse(claims["company"].(string))

    var project types.Project
    project.CompanyID = companyUUID
    project.IsActive = true
    if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
        if strings.Contains(err.Error(), "gross_area") {
            return &types.BCAError{
                Code:    http.StatusBadRequest,
                Message: errors.New("Invalid gross area"),
            }
        }
        if strings.Contains(err.Error(), "net_area") {
            return &types.BCAError{
                Code:    http.StatusBadRequest,
                Message: errors.New("Invalid net area"),
            }
        }
        return &types.BCAError{
            Code:    http.StatusBadRequest,
            Message: err,
        }
    }

    if err := validation.ValidateString(project.Name, 3, true); err != nil {
        return &types.BCAError{
            Code:    http.StatusBadRequest,
            Message: err,
        }
    }

    if project.GrossArea < 0 {
        return &types.BCAError{
            Code:    http.StatusBadRequest,
            Message: errors.New("Invalid gross area"),
        }
    }

    if project.NetArea < 0 {
        return &types.BCAError{
            Code:    http.StatusBadRequest,
            Message: errors.New("Invalid net area"),
        }
    }

    returnedProject, err := s.DB.CreateProject(project)
    if err != nil {
        return &types.BCAError{
            Code:    http.StatusInternalServerError,
            Message: err,
        }
    }

    response := map[string]types.Project{
        "project": returnedProject,
    }
    w.WriteHeader(http.StatusCreated)
    return json.NewEncoder(w).Encode(response)
}
