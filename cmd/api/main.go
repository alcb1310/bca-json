package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/foo", handleFoo)
	listenAddr := ":42069"
	http.ListenAndServe(listenAddr, router)
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
