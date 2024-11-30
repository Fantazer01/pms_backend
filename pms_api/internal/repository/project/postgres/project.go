package postgres

import (
	"context"
	"errors"
	"pms_backend/pms_api/internal/pkg/model"

	"github.com/jackc/pgx/v5"
)

func (r *repository) GetProjectsPaged(ctx context.Context, pageInfo *model.PageInfo) ([]*model.ProjectShort, int, error) {
	args := pgx.NamedArgs{
		"page_size":   pageInfo.PageSize,
		"page_offset": pageInfo.GetOffset(),
	}
	rows, err := r.pool.Query(ctx, getProjectsQuery, args)
	if err != nil {
		return nil, 0, err
	}
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[project])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*model.ProjectShort{}, 0, nil
		}
		return nil, 0, err
	}
	return toProjectSliceFromDb(items), 0, nil
}

func (r *repository) GetProjectByID(ctx context.Context, projectID string) (*model.Project, error) {
	rows, err := r.pool.Query(ctx, getProjectByID, pgx.NamedArgs{"project_id": projectID})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	project, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[project])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return toProjectFromDb(project), nil
}

func (r *repository) CreateProject(ctx context.Context, project *model.Project) error {
	_, err := r.pool.Exec(ctx, createProject, pgx.NamedArgs{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
		"created_at":  project.CreatedAt,
		"updated_at":  project.UpdatedAt,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateProject(ctx context.Context, project *model.Project) error {
	return nil
}

func (r *repository) DeleteProject(ctx context.Context, projectID string) error {
	return nil
}

func (r *repository) GetArchivedProjectsPaged(ctx context.Context, pageInfo *model.PageInfo) ([]*model.ProjectShort, int, error) {
	return nil, 0, nil
}

func (r *repository) ArchiveProject(ctx context.Context, projectID string) error {
	return nil
}

func (r *repository) UnarchiveProject(ctx context.Context, projectID string) error {
	return nil
}
