package postgres

import (
	"context"
	"errors"
	"pms_backend/pms_api/internal/pkg/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(p *pgxpool.Pool) *userRepository {
	return &userRepository{
		pool: p,
	}
}

func (r *userRepository) GetUsers(ctx context.Context, pageInfo *model.PageInfo) ([]*model.UserShort, int, error) {
	var countUsers int
	err := r.pool.QueryRow(ctx, countUsersQuery).Scan(&countUsers)
	if err != nil {
		return nil, 0, err
	}
	args := pgx.NamedArgs{
		"page_size":   pageInfo.PageSize,
		"page_offset": pageInfo.GetOffset(),
	}
	rows, err := r.pool.Query(ctx, getUsersQuery, args)
	if err != nil {
		return nil, 0, err
	}
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[userShort])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	return toUserShortsFromRepo(items), countUsers, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	rows, err := r.pool.Query(ctx, getUserByID, pgx.NamedArgs{"id": userID})
	if err != nil {
		return nil, err
	}
	item, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[user])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return toUserFromRepo(item), nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.pool.Exec(ctx, createUser, pgx.NamedArgs{
		"id":          user.ID,
		"login":       user.Username,
		"password":    user.Password,
		"first_name":  user.FirstName,
		"middle_name": user.MiddleName,
		"last_name":   user.LastName,
		"position":    user.Position,
		"created_at":  user.CreatedAt,
		"updated_at":  user.UpdatedAt,
	})
	return err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	_, err := r.pool.Exec(ctx, updateUser, pgx.NamedArgs{
		"id":          user.ID,
		"login":       user.Username,
		"first_name":  user.FirstName,
		"middle_name": user.MiddleName,
		"last_name":   user.LastName,
		"position":    user.Position,
		"updated_at":  user.UpdatedAt,
	})
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, userID string) error {
	_, err := r.pool.Exec(ctx, deleteUser, pgx.NamedArgs{"id": userID})
	return err
}

func (r *userRepository) GetUserProjects(ctx context.Context, userID string) ([]*model.ProjectShort, error) {
	rows, err := r.pool.Query(ctx, getProjectsOfUserQuery, pgx.NamedArgs{"user_id": userID})
	if err != nil {
		return nil, err
	}
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[projectShort])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return toProjectShortsFromDb(items), nil
}
