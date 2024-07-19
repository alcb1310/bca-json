package server

import (
	"encoding/json"
	"net/http"

	"github.com/alcb1310/bca-json/internals/types"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func handleErrors(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			e, ok := err.(*types.BCAError)
			if ok {
                response := map[string]string{
                    "error": e.Error(),
                }
                w.WriteHeader(e.Code)
                if err := json.NewEncoder(w).Encode(&response); err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                }
				return
			}
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
