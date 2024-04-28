package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/alcb1310/bca-json/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "error", err)
	}

	s := server.NewServer()

	s.MountHandlers()
	listenAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	slog.Info("Starting server on", "addr", listenAddr)
	http.ListenAndServe(listenAddr, s.Router)
}
