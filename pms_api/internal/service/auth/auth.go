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

func (s *authService) Authentication(ctx context.Context, login string, password string) (*model.Tokens, error) {
	user, err := s.authRepository.GetUserByUsername(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("authentication, getting user by username: %w", err)
	}
	if user == nil {
		return nil, apperror.Unauthorized
	}
	hash := sha256.New()
	_, err = hash.Write([]byte(password))
	if err != nil {
		return nil, fmt.Errorf("authentication, hash password: %w", err)
	}
	if !bytes.Equal(user.Password, hash.Sum(nil)) {
		return nil, apperror.Unauthorized
	}

	return s.generateTokens(user.ID, login, user.IsAdmin)
}

func (s *authService) RefreshTokens(ctx context.Context, refreshToken string) (*model.Tokens, error) {
	claims := &model.AppClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil || !token.Valid || !claims.IsRefreshToken {
		return nil, apperror.Unauthorized
	}

	user, err := s.authRepository.GetUserByUsername(ctx, claims.Username)
	if err != nil {
		return nil, fmt.Errorf("refresh token, getting user by username: %w", err)
	}
	if user == nil {
		return nil, apperror.Unauthorized
	}

	return s.generateTokens(claims.Subject, claims.Username, claims.Admin)
}

func (s *authService) generateTokens(userID, username string, isAdmin bool) (*model.Tokens, error) {
	accessClaims := &model.AppClaims{
		Username:       username,
		Admin:          isAdmin,
		IsRefreshToken: false,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	accessJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedToken, err := accessJWT.SignedString(s.secretKey)
	if err != nil {
		return nil, err
	}

	refreshClaims := &model.AppClaims{
		Username:       username,
		Admin:          isAdmin,
		IsRefreshToken: true,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * 24 * time.Hour)),
		},
	}

	refreshJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshJWT.SignedString(s.secretKey)
	if err != nil {
		return nil, err
	}
	return &model.Tokens{
		AccessToken:  signedToken,
		RefreshToken: refreshToken,
	}, nil
}
