package interfaces

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

type AuthService interface {
	Authentication(ctx context.Context, login string, password string) (*model.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*model.Tokens, error)
}
