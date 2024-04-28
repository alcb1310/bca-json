package main

import (
	"log/slog"
	"net/http"

	"github.com/alcb1310/bca-json/internal/server"
)

func main() {
	s := server.NewServer()

	s.MountHandlers()
	listenAddr := ":42069"

	slog.Info("Starting server on", "addr", listenAddr)
	http.ListenAndServe(listenAddr, s.Router)
}
