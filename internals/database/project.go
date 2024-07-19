package database

import "github.com/alcb1310/bca-json/internals/types"

func (s *service) CreateProject(project types.Project) (types.Project, error) {
    query := "INSERT INTO project (name, is_active, gross_area, net_area, company_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
    if err := s.DB.QueryRow(query, project.Name, project.IsActive, project.GrossArea, project.NetArea, project.CompanyID).Scan(&project.ID); err != nil {
        return types.Project{}, err
    }

    return project, nil
}
