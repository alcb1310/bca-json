package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/joho/godotenv/autoload"

	"github.com/alcb1310/bca-json/internals/database"
	"github.com/alcb1310/bca-json/internals/server"
)

var (
    host     = os.Getenv("DB_HOST")
    port     = os.Getenv("DB_PORT")
    username = os.Getenv("DB_USER")
    password = os.Getenv("DB_PASSWORD")
    databaseName = os.Getenv("DB_NAME")
    secret = os.Getenv("JWT_SECRET")
)

func main() {
    if host == "" {
        slog.Error("DB_HOST is not set")
        os.Exit(1)
    }
    if port == "" {
        slog.Error("DB_PORT is not set")
        os.Exit(1)
    }
    if username == "" {
        slog.Error("DB_USER is not set")
        os.Exit(1)
    }
    if password == "" {
        slog.Error("DB_PASSWORD is not set")
        os.Exit(1)
    }
    if databaseName == "" {
        slog.Error("DB_NAME is not set")
        os.Exit(1)
    }
    if secret == "" {
        slog.Error("JWT_SECRET is not set")
        os.Exit(1)
    }

	slog.Info("Starting BCA JSON API Server")

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, databaseName)
	db := database.New(connStr)
    if err :=  db.LoadScript(filepath.Join(".", "scripts", "tables.sql")); err != nil {
        os.Exit(1)
    }
    slog.Info("Tables created")
	s := server.New(db, secret)

	port := os.Getenv("PORT")
	slog.Info("Starting server", "port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), s.Server); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
