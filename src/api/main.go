package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"github.com/alcb1310/bca-json/internals/database"
	"github.com/alcb1310/bca-json/internals/server"
)

func main() {
	slog.Info("Starting BCA JSON API Server")

	db := database.New()
    db.CreateTables()
	s := server.New(db)

	port := os.Getenv("PORT")
	slog.Info("Starting server", "port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), s.Server); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
