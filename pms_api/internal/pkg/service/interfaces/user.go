package interfaces

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

type UserService interface {
	GetUsers(ctx context.Context, pageInfo *model.PageInfo, isAdmin *bool) ([]*model.UserShort, int, error)
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.UserInserted) (*model.User, error)
	UpdateUser(ctx context.Context, userID string, user *model.UserInserted) (*model.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetUserProjects(ctx context.Context, userID string) ([]*model.ProjectShort, error)
}
