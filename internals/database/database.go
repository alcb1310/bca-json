package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/alcb1310/bca-json/internals/types"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface{
    CreateTables()

    // Company methods
    CreateCompany(company *types.Company, user types.CreateUser) (types.User, error)
}

type service struct {
	DB *sql.DB
}

var (
    host     = os.Getenv("DB_HOST")
    port     = os.Getenv("DB_PORT")
    username = os.Getenv("DB_USER")
    password = os.Getenv("DB_PASSWORD")
    database = os.Getenv("DB_NAME")
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

func (s *service) CreateTables() {
    data, err := os.OpenFile("./scripts/tables.sql", os.O_RDONLY, 0644)
    if err != nil {
        slog.Error("Unable to open scripts file", "err", err)
        os.Exit(1)
    }
    defer data.Close()

	info, _ := data.Stat()
	bs := make([]byte, info.Size())
	if _, err := bufio.NewReader(data).Read(bs); err != nil {
		slog.Error("Unable to read file", "err", err)
		os.Exit(1)
	}

    queries := strings.Split(string(bs), ";")

    tx, err := s.DB.Begin()
    if err != nil {
        slog.Error("Unable to create transaction", "err", err)
        os.Exit(1)
    }
    defer tx.Rollback()

    for _, query := range queries {
        if _, err := tx.Exec(query); err != nil {
            slog.Error("Unable to create tables", "err", err)
            os.Exit(1)
        }
    }

    if err := tx.Commit(); err != nil {
        slog.Error("Unable to commit transaction", "err", err)
        os.Exit(1)
    }
    slog.Info("Tables created")
}
