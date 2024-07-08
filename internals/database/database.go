package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Service interface{}

type service struct {
	DB *sql.DB
}

var (
    host     = os.Getenv("DB_HOST")
    port     = os.Getenv("DB_PORT")
    username = os.Getenv("DB_USER")
    password = os.Getenv("DB_PASSWORD")
    database = os.Getenv("DB_DATABASE")
)

func New() Service {
    db := service{}
    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
    d, err := sql.Open("pgx", connStr)
    if err != nil {
        slog.Error("Error connecting to the database", "error", err)
        os.Exit(1)
    }
    db.DB = d

    if err := db.DB.Ping(); err != nil {
        slog.Error("Error connecting to the database", "error", err)
        os.Exit(1)
    }

    return &db
}
