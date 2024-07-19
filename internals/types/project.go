package types

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	IsActive    bool       `json:"is_active"`
	GrossArea   float64    `json:"gross_area"`
	NetArea     float64    `json:"net_area"`
	LastClosure *time.Time `json:"last_closure"`
	CompanyID   uuid.UUID  `json:"company_id"`
}
