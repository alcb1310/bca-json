package database

import (
	"github.com/google/uuid"

	"github.com/alcb1310/bca-json/internals/types"
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
