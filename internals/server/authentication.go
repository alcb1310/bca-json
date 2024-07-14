package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/internals/validation"
)

func (s *Server) CreateCompany(w http.ResponseWriter, r *http.Request) error {
	company := &types.Company{}
	user := &types.CreateUser{}

	if r.Body == nil {
		e := &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: errors.New("Missing Body"),
		}
		return e
	}
	data := &types.CreateCompany{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		slog.Error("CreateCompany: Request body", "error", err)
		return err
	}

	if err := validation.ValidateRuc(data.Ruc, true); err != nil {
		e := &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
		return e
	} else {
		company.Ruc = data.Ruc
	}

	if err := validation.ValidateString(data.Name, 3, true); err != nil {
		e := &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
		return e
	} else {
		company.Name = data.Company.Name
	}

	if err := validation.ValidateEmail(data.Email, true); err != nil {
		e := &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
		return e
	} else {
		user.Email = data.Email
	}

	if err := validation.ValidateString(data.UserName, 3, true); err != nil {
		e := &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
		return e
	} else {
		user.Name = data.UserName
	}

	if err := validation.ValidatePassword(data.Password, 3, true); err != nil {
		e := &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
		return e
	} else {
		user.Password = data.Password
	}

	company.Employees = data.Employees
	if company.Employees == 0 {
		company.Employees = 1
	}
	company.IsActive = true
	user.RoleID = "a"

	u, err := s.DB.CreateCompany(company, *user)
	if err != nil {
		e := &types.BCAError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return e
	}

	ret := map[string]interface{}{
		"company": company,
		"user":    u,
	}

	if err := json.NewEncoder(w).Encode(ret); err != nil {
		e := &types.BCAError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return e
	}

	return nil
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) error {
	data := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		slog.Error("Login: Request body", "error", err)
		return err
	}

	if err := validation.ValidateEmail(data.Email, true); err != nil {
		e := &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
		return e
	}

	if err := validation.ValidatePassword(data.Password, 3, true); err != nil {
		e := &types.BCAError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
		return e
	}

	u, err := s.DB.Login(data.Email, data.Password)
	if err != nil {
		e := &types.BCAError{
			Code:    http.StatusUnauthorized,
			Message: errors.New("Invalid credentials"),
		}
		return e
	}
	_, tokenString, _ := s.TokenAuth.Encode(map[string]interface{}{
		"id":      u.ID,
		"name":    u.Name,
		"email":   u.Email,
		"company": u.CompanyID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
