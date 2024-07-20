package database

import (
	"github.com/alcb1310/bca-json/internals/types"
	"github.com/google/uuid"
)

func (s *service) GetProjects(companyID uuid.UUID) ([]types.Project, error) {
    query := "SELECT id, name, is_active, gross_area, net_area, company_id FROM project WHERE company_id = $1"
    rows, err := s.DB.Query(query, companyID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    projects := []types.Project{}
    for rows.Next() {
        project := types.Project{}
        if err := rows.Scan(&project.ID, &project.Name, &project.IsActive, &project.GrossArea, &project.NetArea, &project.CompanyID); err != nil {
            return nil, err
        }
        projects = append(projects, project)
    }

    return projects, nil
}

func (s *service) CreateProject(project types.Project) (types.Project, error) {
    query := "INSERT INTO project (name, is_active, gross_area, net_area, company_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
    if err := s.DB.QueryRow(query, project.Name, project.IsActive, project.GrossArea, project.NetArea, project.CompanyID).Scan(&project.ID); err != nil {
        return types.Project{}, err
    }

    return project, nil
}
