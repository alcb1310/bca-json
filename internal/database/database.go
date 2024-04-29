package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/alcb1310/bca-json/internal/types"
)

type database struct {
	DB *sql.DB
}

type Database interface {
	Health() error
	CreateSchema() error
	Register(reg types.RegisterInformation) (uuid.UUID, error)
}

type config struct {
	Database string
	User     string
	Host     string
	Password string
	Port     string
}

var c = config{
	Database: os.Getenv("PGDATABASE"),
	User:     os.Getenv("PGUSER"),
	Host:     os.Getenv("PGHOST"),
	Password: os.Getenv("PGPASSWORD"),
	Port:     os.Getenv("PGPORT"),
}

func Connect() Database {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		slog.Error("Error connecting to database", err)
		panic(err)
	}

	return &database{
		DB: db,
	}
}

func (d *database) Health() error {
	return d.DB.Ping()
}

func (d *database) CreateSchema() error {
	slog.Info("Creating schema")

	file, err := os.ReadFile("./internal/database/schema.sql")
	if err != nil {
		slog.Error("Error reading schema file", "error", err)
		return err
	}

	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		if _, err := tx.Exec(request); err != nil {
			slog.Error("Error creating schema", "error", err)
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	slog.Info("Schema created")
	return nil
}
