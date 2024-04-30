package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/alcb1310/bca-json/internal/types"
	"github.com/alcb1310/bca-json/internal/utils"
)

func (s *Server) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	clear(errorResponse)
	w.Header().Set("Content-Type", "application/json")
	var credentials types.CredentialsType

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		if err.Error() == "EOF" {
			errorResponse["error"] = "No se recibió JSON"
			slog.Error("No se recibió JSON")
			w.WriteHeader(http.StatusBadRequest)
			return json.NewEncoder(w).Encode(errorResponse)
		}
	}

	if err := utils.ValidateEmail(credentials.Email); err != nil {
		error := make(map[string]interface{})
		error["field"] = "email"
		error["message"] = err.Error()
		errorResponse["error"] = error

		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errorResponse)
	}

	if err := utils.ValidatePassword(credentials.Password); err != nil {
		error := make(map[string]interface{})
		error["field"] = "password"
		error["message"] = err.Error()
		errorResponse["error"] = error

		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errorResponse)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
