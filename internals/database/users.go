package database

import (
	"github.com/google/uuid"

	"github.com/alcb1310/bca-json/internals/types"
	"github.com/alcb1310/bca-json/internals/utils"
)

func (s *service) GetUsers(companyID uuid.UUID) ([]types.User, error) {
	users := []types.User{}

	query := "SELECT id, email, name, role_id, company_id FROM \"user\" WHERE company_id = $1"
	rows, err := s.DB.Query(query, companyID)
	if err != nil {
		return []types.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		user := types.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.RoleID, &user.CompanyID); err != nil {
			return []types.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *service) GetUserByID(id, companyID uuid.UUID) (types.User, error) {
	user := types.User{}
	query := "SELECT id, email, name, role_id, company_id FROM \"user\" WHERE id = $1 and company_id = $2"
	err := s.DB.QueryRow(query, id, companyID).Scan(&user.ID, &user.Email, &user.Name, &user.RoleID, &user.CompanyID)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (s *service) CreateUser(user types.CreateUser) (types.User, error) {
	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		return types.User{}, err
	}
	user.Password = pass

	query := "INSERT INTO \"user\" (email, password, name, role_id, company_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	if err := s.DB.
		QueryRow(
			query,
			user.Email,
			user.Password,
			user.Name,
			user.RoleID,
			user.CompanyID,
		).Scan(&user.ID); err != nil {
		return types.User{}, err
	}
	return types.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		RoleID:    user.RoleID,
		CompanyID: user.CompanyID,
	}, nil
}

func (s *service) DeleteUser(id, companyID uuid.UUID) error {
    query := "DELETE FROM \"user\" WHERE id = $1 and company_id = $2"
    _, err := s.DB.Exec(query, id, companyID)
    if err != nil {
        return err
    }
    return nil
}

func (s *service) UpdateUser(user types.CreateUser) error {
    if user.Password == "" {
        query := "UPDATE \"user\" SET email = $1, name = $2 WHERE id = $3 and company_id = $4"
        if _, err := s.DB.Exec(query, user.Email, user.Name, user.ID, user.CompanyID); err != nil {
            return err
        }
        return nil
    }

    pass, err := utils.HashPassword(user.Password)
    if err != nil {
        return err
    }
    query := "UPDATE \"user\" SET email = $1, password = $2, name = $3 WHERE id = $4 and company_id = $5"
    if _, err := s.DB.Exec(query, user.Email, pass, user.Name, user.ID, user.CompanyID); err != nil {
        return err
    }
    return nil
}
