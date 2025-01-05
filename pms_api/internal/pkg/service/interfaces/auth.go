package interfaces

import "context"

type AuthService interface {
	Authentication(ctx context.Context, login string, password string) (token string, err error)
}
