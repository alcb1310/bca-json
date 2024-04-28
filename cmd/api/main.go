package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/alcb1310/bca-json/internal/database"
	"github.com/alcb1310/bca-json/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "error", err)
	}

	db := database.Connect()
	if err := db.Health(); err != nil {
		slog.Error("Error connecting to database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database connected")

	if err := db.CreateSchema(); err != nil {
		slog.Error("Error creating schema", "error", err)
		os.Exit(1)
	}
	s := server.NewServer(db)

	s.MountHandlers()
	listenAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	slog.Info("Starting server on", "addr", listenAddr)
	http.ListenAndServe(listenAddr, s.Router)
}
