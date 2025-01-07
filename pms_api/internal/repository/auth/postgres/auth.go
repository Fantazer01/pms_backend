package postgres

import (
	"context"
	"errors"
	"pms_backend/pms_api/internal/pkg/model"

	"github.com/jackc/pgx/v5"
)

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	rows, err := r.pool.Query(ctx, getUserByUsername, pgx.NamedArgs{"username": username})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	u, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[user])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &model.User{
		ID:         u.ID,
		Username:   u.Username,
		Password:   u.Password,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
		LastName:   u.LastName,
		Position:   u.Position,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}, nil
}
