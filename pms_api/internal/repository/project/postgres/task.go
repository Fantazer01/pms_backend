package postgres

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

func (r *repository) GetProjectTasks(ctx context.Context, projectID string) ([]*model.Task, error) {
	return nil, nil
}
