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
)

func main() {
	slog.Info("Starting BCA JSON API Server")

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, databaseName)
	db := database.New(connStr)
    if err :=  db.LoadScript(filepath.Join(".", "scripts", "tables.sql")); err != nil {
        os.Exit(1)
    }
    slog.Info("Tables created")
	s := server.New(db)

	port := os.Getenv("PORT")
	slog.Info("Starting server", "port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), s.Server); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
