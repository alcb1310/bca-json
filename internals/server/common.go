package server

import "net/http"

type Handler  func(w http.ResponseWriter, r *http.Request) error

func handleErrors(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("Hello, World!"))
	return err
}
