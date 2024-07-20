package database

import (
	"bufio"
	"database/sql"
	"log/slog"
	"os"
	"strings"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/alcb1310/bca-json/internals/types"
)

type Service interface {
	LoadScript(fileName string) error
	GetRole(name string) (types.Role, error)

	// Company methods
	CreateCompany(company *types.Company, user types.CreateUser) (types.User, error)
	Login(email, password string) (types.User, error)

    // User methods
    GetUsers(companyID uuid.UUID) ([]types.User, error)
    GetUserByID(id, companyID uuid.UUID) (types.User, error)
    CreateUser(user types.CreateUser) (types.User, error)
    DeleteUser(id, companyID uuid.UUID) error
    UpdateUser(user types.CreateUser) error

    // Project methods
    GetProjects(companyID uuid.UUID) ([]types.Project, error)
    CreateProject(project types.Project) (types.Project, error)
}

type service struct {
	DB *sql.DB
}

func New(connStr string) Service {
	db := service{}
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

func (s *service) LoadScript(fileName string) error {
	commited := false
	data, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		slog.Error("Unable to open scripts file", "err", err)
		return err
	}
	defer data.Close()

	info, _ := data.Stat()
	bs := make([]byte, info.Size())
	if _, err := bufio.NewReader(data).Read(bs); err != nil {
		slog.Error("Unable to read file", "err", err)
		return err
	}

	queries := strings.Split(string(bs), ";")

	tx, err := s.DB.Begin()
	if err != nil {
		slog.Error("Unable to create transaction", "err", err)
		return err
	}
	defer func() {
		if commited {
			return
		}
		if err := tx.Rollback(); err != nil {
			slog.Error("Error rolling back the transaction", "error", err)
		}
	}()

	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			slog.Error("Unable to create tables", "err", err)
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		slog.Error("Unable to commit transaction", "err", err)
		return err
	}
	commited = true
	return nil
}

func (s *service) GetRole(name string) (types.Role, error) {
	r := types.Role{}

	query := "SELECT id, name FROM role WHERE name = $1"
	if err := s.DB.QueryRow(query, name).Scan(&r.ID, &r.Name); err != nil {
		return r, err
	}

	return r, nil
}
