package main

import (
	"net/http"

	"github.com/alcb1310/bca-json/internal/server"
)

func main() {
	s := server.NewServer()

	s.MountHandlers()
	listenAddr := ":42069"

	http.ListenAndServe(listenAddr, s.Router)
}
