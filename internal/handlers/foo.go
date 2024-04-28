package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["message"] = "Hello, World!"
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
