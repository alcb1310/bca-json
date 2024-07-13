package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/internals/utils"
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
		return nil
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
		user.Password, err = utils.HashPassword(data.Password)
		if err != nil {
			d := &types.BCAError{
				Code:    http.StatusInternalServerError,
				Message: err,
			}
			return d
		}
	}

	company.Employees = data.Employees
	company.IsActive = true
	user.RoleID = "a"

	if err := s.DB.CreateCompany(company, user); err != nil {
		e := &types.BCAError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return e
	}

	ret := map[string]interface{}{
		"company": company,
		"user":    user.User,
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
