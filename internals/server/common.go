package server

import (
	"encoding/json"
	"net/http"
)

type Handler  func(w http.ResponseWriter, r *http.Request) error

func handleErrors(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func ContentTypeJSON(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func Home(w http.ResponseWriter, r *http.Request) error {
    res := map[string]string{
        "message": "Hello World!",
    }

    err := json.NewEncoder(w).Encode(res)
	return err
}
