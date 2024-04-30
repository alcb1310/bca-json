package database

import (
	"errors"

	"github.com/google/uuid"

	"github.com/alcb1310/bca-json/internal/types"
	"github.com/alcb1310/bca-json/internal/utils"
)

func (d *database) Register(reg types.RegisterInformation) (uuid.UUID, error) {
	var companyId, userId uuid.UUID
	tx, err := d.DB.Begin()
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback()

	if reg.Employees == nil {
		var emp uint = 1
		reg.Employees = &emp
	}

	sqlQuery := "INSERT INTO company (ruc, name, employees) VALUES ($1, $2, $3) RETURNING id"
	if err = tx.QueryRow(sqlQuery, reg.RUC, reg.Name, reg.Employees).Scan(&companyId); err != nil {
		return uuid.Nil, err
	}

	password, err := utils.EncryptPasssword(reg.Password)

	sqlQuery = "INSERT INTO public.user (email, password, name, company_id, role_id) VALUES ($1, $2, $3, $4, 'a') RETURNING id"

	if err = tx.QueryRow(sqlQuery, reg.Email, string(password), reg.UserName, companyId).Scan(&userId); err != nil {
		return uuid.Nil, err
	}

	tx.Commit()
	return companyId, nil
}

func (d *database) Login(credentials types.CredentialsType) error {
	var password string
	sqlQuery := "SELECT password FROM public.user WHERE email = $1"

	if err := d.DB.QueryRow(sqlQuery, credentials.Email).Scan(&password); err != nil {
		return errors.New("Credenciales inválidas")
	}

	if err := utils.ComparePassword(password, credentials.Password); err != nil {
		return errors.New("Credenciales inválidas")
	}

	return nil
}
