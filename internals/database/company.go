package database

import (
	"errors"
	"log/slog"

	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/internals/utils"
)

func (s *service) CreateCompany(company *types.Company, user types.CreateUser) (types.User, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		slog.Error("Error creating transaction", "error", err)
		return types.User{}, err
	}
	defer  func () {
        if err := tx.Rollback(); err != nil {
            slog.Error("Error rolling back transaction", "error", err)
        }
    }()

	query := "INSERT INTO company (ruc, name, employees, is_active) VALUES ($1, $2, $3, $4) RETURNING id"
	if err := tx.QueryRow(query, company.Ruc, company.Name, company.Employees, company.IsActive).Scan(&company.ID); err != nil {
		slog.Error("Error creating company", "error", err)
		return types.User{}, err
	}

	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		slog.Error("Error hashing password", "error", err)
		return types.User{}, err
	}
	query = "INSERT INTO \"user\" (email, name, password, company_id, role_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	if err := tx.QueryRow(query, user.Email, user.Name, pass, company.ID, user.RoleID).Scan(&user.ID); err != nil {
		slog.Error("Error creating user", "error", err)
		return types.User{}, err
	}
    user.CompanyID = company.ID

    if err := tx.Commit(); err != nil {
        slog.Error("Error commiting the transaction", "error", err)
        return types.User{}, err
    }
	return user.User, nil
}
func (s *service) Login(email, password string) (types.User, error) {
    var pass string
    u := types.User{}

    query := "SELECT id, email, name, company_id, password from \"user\" WHERE email = $1"
    if err := s.DB.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.Name, &u.CompanyID, &pass); err != nil {
        slog.Error("Error querying user", "error", err)
        return u, err
    }

    if !utils.CheckPasswordHash(password, pass){
        err := errors.New("Wrong password")
        slog.Error("Error comparing passwords", "error", err)
        return types.User{}, err
    }
    return u, nil
}
