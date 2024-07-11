package types

import (
	"github.com/google/uuid"
)

type CreateCompany struct {
    Company
    CreateUser
}

type Company struct {
	ID        uuid.UUID `json:"id"`
	Ruc       string    `json:"ruc"`
	Name      string    `json:"name"`
	Employees uint      `json:"employees"`
	IsActive  bool      `json:"is_active"`
}

type User struct {
	ID     uuid.UUID `json:"id"`
	Email  string    `json:"email"`
	Name   string    `json:"name"`
	RoleID string    `json:"role_id"`
}

type CreateUser struct {
	User
	Password string `json:"password"`
}

