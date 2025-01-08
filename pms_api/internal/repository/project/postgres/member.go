package postgres

import (
	"context"
	"errors"
	"pms_backend/pms_api/internal/pkg/model"

	"github.com/jackc/pgx/v5"
)

func (r *repository) GetProjectMembers(ctx context.Context, projectID string) ([]*model.UserShort, error) {
	rows, err := r.pool.Query(ctx, getProjectMembers, pgx.NamedArgs{"project_id": projectID})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[user])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*model.UserShort{}, nil
		}
		return nil, err
	}

	return toUserShortsFromRepo(items), nil
}

func (r *repository) AddProjectMember(ctx context.Context, projectID string, member *model.Member) error {
	_, err := r.pool.Exec(ctx, insertUserToProject, pgx.NamedArgs{
		"project_id":       projectID,
		"user_id":          member.UserID,
		"role":             member.Role,
		"is_admin_project": member.IsAdminProject,
	})
	return err
}

func (r *repository) DeleteProjectMember(ctx context.Context, projectID, userID string) error {
	_, err := r.pool.Exec(ctx, deleteUserFromProject, pgx.NamedArgs{
		"project_id": projectID,
		"user_id":    userID,
	})
	return err
}
