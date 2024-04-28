package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) HandleFoo(w http.ResponseWriter, r *http.Request) error {
	response := make(map[string]interface{})
	response["message"] = "Hello, World!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(response)
}
