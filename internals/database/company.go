package database

import (
	"log/slog"

	"github.com/alcb1310/bca-json/internals/types"
)


func (s *service) CreateCompany(company *types.Company, user *types.CreateUser) error {
    tx, err := s.DB.Begin()
    if err!=nil {
        slog.Error("Error creating transaction", "error", err)
        return err
    }
    defer tx.Rollback()

    query := "INSERT INTO company (ruc, name, employees, is_active) VALUES ($1, $2, $3, $4) RETURNING id"
    if err := tx.QueryRow(query, company.Ruc, company.Name, company.Employees, company.IsActive).Scan(&company.ID); err != nil {
        slog.Error("Error creating company", "error", err)
        return err
    }

    query = "INSERT INTO \"user\" (email, name, password, company_id, role_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
    if err := tx.QueryRow(query, user.Email, user.Name, user.Password, company.ID, user.RoleID).Scan(&user.ID); err != nil {
        slog.Error("Error creating user", "error", err)
        return err
    }

    tx.Commit()
    return nil
}
