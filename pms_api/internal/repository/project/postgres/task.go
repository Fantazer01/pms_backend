package postgres

import (
	"context"
	"errors"
	"pms_backend/pms_api/internal/pkg/model"

	"github.com/jackc/pgx/v5"
)

func (r *repository) GetProjectTasks(ctx context.Context, projectID string) ([]*model.Task, error) {
	rows, err := r.pool.Query(ctx, getProjectTasks, pgx.NamedArgs{"project_id": projectID})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[task])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*model.Task{}, nil
		}
		return nil, err
	}
	return toTasksFromRepo(items), nil
}
