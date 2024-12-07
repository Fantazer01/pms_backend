package postgres

import (
	"pms_backend/pms_api/internal/pkg/model"
	"time"
)

type project struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func toProjectShortsFromDb(projectsFromDb []project) []*model.ProjectShort {
	projects := make([]*model.ProjectShort, len(projectsFromDb))
	for i := range projects {
		projects[i] = &model.ProjectShort{
			ID:   projectsFromDb[i].ID,
			Name: projectsFromDb[i].Name,
		}
	}
	return projects
}

func toProjectFromDb(projectFromDb project) *model.Project {
	return &model.Project{
		ID:          projectFromDb.ID,
		Name:        projectFromDb.Name,
		Description: projectFromDb.Description,
		CreatedAt:   projectFromDb.CreatedAt,
		UpdatedAt:   projectFromDb.UpdatedAt,
	}
}
