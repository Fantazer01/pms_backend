package postgres

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

func (r *repository) GetProjectMembers(ctx context.Context, projectID string) ([]*model.User, error) {
	return nil, nil
}

func (r *repository) AddProjectMember(ctx context.Context, projectID, userID, roleID string) error {
	return nil
}

func (r *repository) DeleteProjectMember(ctx context.Context, projectID, userID, roleID string) error {
	return nil
}
