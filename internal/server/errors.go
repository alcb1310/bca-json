package server

import (
	"encoding/json"
	"net/http"
)

var responses = make(map[string]interface{})

func NotFound(w http.ResponseWriter, r *http.Request) error {
	clear(responses)
	responses["error"] = "No se encuentra la ruta"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	return json.NewEncoder(w).Encode(responses)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) error {
	clear(responses)
	responses["error"] = "Método no permitido"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	return json.NewEncoder(w).Encode(responses)
}
