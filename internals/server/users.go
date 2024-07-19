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
	var companyUUID uuid.UUID
	_, claims, _ := jwtauth.FromContext(r.Context())

	companyUUID = uuid.MustParse(claims["company"].(string))
	id := chi.URLParam(r, "userID")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	user, err := s.DB.GetUserByID(parsedID, companyUUID)
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

func (s *Server) GetCurrentUser(w http.ResponseWriter, r *http.Request) error {
	var companyUUID, userUUID uuid.UUID
	_, claims, _ := jwtauth.FromContext(r.Context())

	companyUUID = uuid.MustParse(claims["company"].(string))
	userUUID = uuid.MustParse(claims["id"].(string))
	user, _ := s.DB.GetUserByID(userUUID, companyUUID)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(user)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Body == nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: errors.New("Missing Body"),
		}
	}

	var user types.CreateUser
	_, claims, _ := jwtauth.FromContext(r.Context())

	user.CompanyID = uuid.MustParse(claims["company"].(string))
	user.RoleID = "a"
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	if err := validation.ValidateEmail(user.Email, true); err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	if err := validation.ValidatePassword(user.Password, 3, true); err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	if err := validation.ValidateString(user.Name, 3, true); err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	userResponse, err := s.DB.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return &types.BCAError{
				Code:    http.StatusConflict,
				Message: errors.New("User already exists"),
			}
		}

		return &types.BCAError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	res := map[string]types.User{
		"user": userResponse,
	}
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(res)
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	_, claims, _ := jwtauth.FromContext(r.Context())
	companyUUID := uuid.MustParse(claims["company"].(string))
	id := chi.URLParam(r, "userID")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	if uuid.MustParse(claims["id"].(string)) == parsedID {
		return &types.BCAError{
			Code:    http.StatusForbidden,
			Message: errors.New("You can not delete yourself"),
		}
	}

	user, err := s.DB.GetUserByID(parsedID, companyUUID)
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

	if err := s.DB.DeleteUser(user.ID, user.CompanyID); err != nil {
		return &types.BCAError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	var user types.CreateUser
	_, claims, _ := jwtauth.FromContext(r.Context())
	companyUUID, _ := uuid.Parse(claims["company"].(string))
	id := chi.URLParam(r, "userID")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	user.CompanyID = companyUUID
	user.RoleID = "a"

	if r.Body != nil {
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			return &types.BCAError{
				Code:    http.StatusBadRequest,
				Message: err,
			}
		}
	}

	if err := validation.ValidateEmail(user.Email, false); err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	if err := validation.ValidateString(user.Name, 3, false); err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	if err := validation.ValidatePassword(user.Password, 3, false); err != nil {
		return &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	currentUser, err := s.DB.GetUserByID(parsedID, companyUUID)
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

	if user.Email == "" {
		user.Email = currentUser.Email
	}

	if user.Name == "" {
		user.Name = currentUser.Name
	}

	if err := s.DB.UpdateUser(user); err != nil {
		return &types.BCAError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	userResponse := map[string]types.User{
		"user": {
			ID:        currentUser.ID,
			Name:      user.Name,
			Email:     user.Email,
			RoleID:    user.RoleID,
			CompanyID: user.CompanyID,
		},
	}
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(userResponse)
}
