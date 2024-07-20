package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"

	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/internals/validation"
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
		if strings.Contains(err.Error(), "23505") {
			return &types.BCAError{
				Code:    http.StatusConflict,
				Message: errors.New("Project already exists"),
			}
		}
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

func (s *Server) GetProjects(w http.ResponseWriter, r *http.Request) error {
    _, claims, _ := jwtauth.FromContext(r.Context())
    if claims == nil {
        return &types.BCAError{
            Code:    http.StatusUnauthorized,
            Message: errors.New("Unauthorized"),
        }
    }

    companyUUID := uuid.MustParse(claims["company"].(string))

    responseProjects, err := s.DB.GetProjects(companyUUID)
    if err != nil {
        return &types.BCAError{
            Code:    http.StatusInternalServerError,
            Message: err,
        }
    }

    w.WriteHeader(http.StatusOK)
    return json.NewEncoder(w).Encode(responseProjects)
}

func (s *Server) GetProjectByID(w http.ResponseWriter, r *http.Request) error {
    _, claims, _ := jwtauth.FromContext(r.Context())
    if claims == nil {
        return &types.BCAError{
            Code:    http.StatusUnauthorized,
            Message: errors.New("Unauthorized"),
        }
    }

    companyUUID := uuid.MustParse(claims["company"].(string))
    id := chi.URLParam(r, "projectID")
    projectUUID, err := uuid.Parse(id)
    if err != nil {
        return &types.BCAError{
            Code:    http.StatusBadRequest,
            Message: err,
        }
    }

    responseProject, err := s.DB.GetProjectByID(projectUUID, companyUUID)
    if err != nil {
        if err == sql.ErrNoRows {
            return &types.BCAError{
                Code:    http.StatusNotFound,
                Message: errors.New("Project not found"),
            }
        }
        return &types.BCAError{
            Code:    http.StatusInternalServerError,
            Message: err,
        }
    }

    res := map[string]types.Project{
        "project": responseProject,
    }
    w.WriteHeader(http.StatusOK)
    return json.NewEncoder(w).Encode(res)
}
