package auth

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/repository/interfaces"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
	secretKey      []byte
	authRepository interfaces.AuthRepository
}

func NewAuthService(secretKey string, r interfaces.AuthRepository) *authService {
	return &authService{
		secretKey:      []byte(secretKey),
		authRepository: r,
	}
}

func (s *authService) Authentication(ctx context.Context, login string, password string) (accessToken string, err error) {
	user, err := s.authRepository.GetUserByUsername(ctx, login)
	if err != nil {
		return "", fmt.Errorf("authentication, getting user by username: %w", err)
	}
	if user == nil {
		return "", apperror.Unauthorized
	}
	hash := sha256.New()
	_, err = hash.Write([]byte(password))
	if err != nil {
		return "", fmt.Errorf("authentication, hash password: %w", err)
	}
	if !bytes.Equal(user.Password, hash.Sum(nil)) {
		return "", apperror.Unauthorized
	}

	claims := &model.AppClaims{
		Username: login,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
