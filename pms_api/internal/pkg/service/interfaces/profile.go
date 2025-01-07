package interfaces

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

type ProfileService interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}
