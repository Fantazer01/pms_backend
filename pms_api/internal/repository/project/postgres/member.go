package postgres

import (
	"context"
	"errors"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"

	"github.com/jackc/pgx/v5"
)

func (r *repository) GetProjectMembers(ctx context.Context, projectID string) ([]*model.Member, error) {
	rows, err := r.pool.Query(ctx, getProjectMembers, pgx.NamedArgs{"project_id": projectID})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[member])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*model.Member{}, nil
		}
		return nil, err
	}

	return toMembersFromRepo(items), nil
}

func (r *repository) getProjectMember(ctx context.Context, projectID string, userID string) (*member, error) {
	rows, err := r.pool.Query(ctx, getProjectMemberByID, pgx.NamedArgs{
		"project_id": projectID,
		"user_id":    userID,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	item, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[member])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (r *repository) getRoles(ctx context.Context) ([]role, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name_role FROM role")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[role])
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *repository) AddProjectMember(ctx context.Context, projectID string, member *model.MemberInserted) error {
	roles, err := r.getRoles(ctx)
	if err != nil {
		return err
	}
	roleID := -1
	for _, v := range roles {
		if member.Role == v.Name {
			roleID = v.ID
		}
	}
	if roleID == -1 {
		return errors.Join(apperror.InvalidValue, errors.New("incorrect value of role"))
	}
	_, err = r.pool.Exec(ctx, insertMemberToProject, pgx.NamedArgs{
		"project_id":       projectID,
		"user_id":          member.UserID,
		"role_id":          roleID,
		"is_admin_project": member.IsAdminProject,
	})
	return err
}

func (r *repository) UpdateProjectMember(ctx context.Context, projectID string, member *model.MemberInserted) error {
	memberFromDb, err := r.getProjectMember(ctx, projectID, member.UserID)
	if err != nil {
		return err
	}
	if memberFromDb == nil {
		return apperror.NotFound
	}
	roles, err := r.getRoles(ctx)
	if err != nil {
		return err
	}
	roleID := -1
	for _, v := range roles {
		if member.Role == v.Name {
			roleID = v.ID
		}
	}
	if roleID == -1 {
		return errors.Join(apperror.InvalidValue, errors.New("incorrect value of role"))
	}
	_, err = r.pool.Exec(ctx, updateMemberToProject, pgx.NamedArgs{
		"project_id":       projectID,
		"user_id":          member.UserID,
		"role_id":          roleID,
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
