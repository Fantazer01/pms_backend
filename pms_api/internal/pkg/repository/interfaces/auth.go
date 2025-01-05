package interfaces

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}
