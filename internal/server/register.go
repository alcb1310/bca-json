package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/alcb1310/bca-json/internal/types"
	"github.com/alcb1310/bca-json/internal/utils"
)

var errorResponse = make(map[string]interface{})

func (s *Server) HandleRegister(w http.ResponseWriter, r *http.Request) error {
	clear(errorResponse)
	w.Header().Set("Content-Type", "application/json")
	var registerData types.RegisterInformation

	if err := json.NewDecoder(r.Body).Decode(&registerData); err != nil {
		if err.Error() == "EOF" {
			errorResponse["error"] = "No se recibió JSON"
			slog.Error("No se recibió JSON")
			w.WriteHeader(http.StatusBadRequest)
			return json.NewEncoder(w).Encode(errorResponse)
		}

		if strings.Contains(err.Error(), "cannot unmarshal number") {
			error := make(map[string]interface{})
			error["field"] = "employees"
			error["message"] = "Los empleados deben ser un número positivo válido"
			errorResponse["error"] = error

			w.WriteHeader(http.StatusBadRequest)
			return json.NewEncoder(w).Encode(errorResponse)
		}

		if strings.Contains(err.Error(), "cannot unmarshal string") {
			error := make(map[string]interface{})
			error["field"] = "employees"
			error["message"] = "Los empleados deben ser un número positivo válido"
			errorResponse["error"] = error

			w.WriteHeader(http.StatusBadRequest)
			return json.NewEncoder(w).Encode(errorResponse)
		}

		errorResponse["error"] = fmt.Sprintf("Error decoding JSON: %v", err)
		slog.Error("Error decoding JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errorResponse)
	}

	if registerData.RUC == "" {
		error := make(map[string]interface{})
		error["field"] = "ruc"
		error["message"] = "Ingrese un RUC"
		errorResponse["error"] = error

		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errorResponse)
	}

	if registerData.Name == "" {
		error := make(map[string]interface{})
		error["field"] = "nombre"
		error["message"] = "Ingrese un Nombre"
		errorResponse["error"] = error

		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errorResponse)
	}

	if err := utils.ValidateEmail(registerData.Email); err != nil {
		error := make(map[string]interface{})
		error["field"] = "email"
		error["message"] = err.Error()
		errorResponse["error"] = error

		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errorResponse)
	}

	if err := utils.ValidatePassword(registerData.Password); err != nil {
		error := make(map[string]interface{})
		error["field"] = "password"
		error["message"] = err.Error()
		errorResponse["error"] = error

		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errorResponse)
	}

	if registerData.UserName == "" {
		error := make(map[string]interface{})
		error["field"] = "username"
		error["message"] = "El nombre de usuario es obligatorio"
		errorResponse["error"] = error

		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errorResponse)
	}

	id, err := s.DB.Register(registerData)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			if strings.Contains(err.Error(), "company_name_unique") {
				slog.Error(fmt.Sprintf("Error: %v", err))
				error := make(map[string]interface{})
				error["field"] = "name"
				error["message"] = "Ya existe una empresa con este nombre"
				errorResponse["error"] = error

				w.WriteHeader(http.StatusConflict)
				return json.NewEncoder(w).Encode(errorResponse)
			}

			if strings.Contains(err.Error(), "company_ruc_unique") {
				slog.Error(fmt.Sprintf("Error: %v", err))
				error := make(map[string]interface{})
				error["field"] = "ruc"
				error["message"] = "Ya existe una empresa con este ruc"
				errorResponse["error"] = error

				w.WriteHeader(http.StatusConflict)
				return json.NewEncoder(w).Encode(errorResponse)
			}

			if strings.Contains(err.Error(), "user_email_unique") {
				slog.Error(fmt.Sprintf("Error: %v", err))
				error := make(map[string]interface{})
				error["field"] = "email"
				error["message"] = "Ya existe un usuario con este correo"
				errorResponse["error"] = error

				w.WriteHeader(http.StatusConflict)
				return json.NewEncoder(w).Encode(errorResponse)
			}
		}

		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(err.Error())
	}

	_ = id

	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(registerData)
}
