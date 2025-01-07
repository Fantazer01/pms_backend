package user

import (
	"context"
	"crypto/sha256"
	"fmt"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/repository/interfaces"
	"time"

	"github.com/google/uuid"
)

type userService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(r interfaces.UserRepository) *userService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) GetUsers(ctx context.Context, pageInfo *model.PageInfo) ([]*model.UserShort, int, error) {
	users, count, err := s.userRepository.GetUsers(ctx, pageInfo)
	if err != nil {
		return nil, 0, fmt.Errorf("getting users: %w", err)
	}
	return users, count, nil
}

func (s *userService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("getting user by id: %w", err)
	}
	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, u *model.UserInserted) (*model.User, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(u.Password))
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}
	pas := hash.Sum(nil)
	now := time.Now()
	user := &model.User{
		ID:         uuid.NewString(),
		Username:   u.Username,
		Password:   pas,
		IsAdmin:    u.IsAdmin,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
		LastName:   u.LastName,
		Position:   u.Position,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, userID string, u *model.UserInserted) (*model.User, error) {
	userFromDB, err := s.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userFromDB == nil {
		return nil, apperror.NotFound
	}
	user := &model.User{
		ID:         userID,
		Username:   u.Username,
		IsAdmin:    userFromDB.IsAdmin,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
		LastName:   u.LastName,
		Position:   u.Position,
		UpdatedAt:  time.Now(),
	}
	err = s.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("updating user: %w", err)
	}
	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, userID string) error {
	userFromDB, err := s.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if userFromDB == nil {
		return apperror.NotFound
	}
	err = s.userRepository.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("deleting user: %w", err)
	}
	return nil
}

func (s *userService) GetUserProjects(ctx context.Context, userID string) ([]*model.ProjectShort, error) {
	projects, err := s.userRepository.GetUserProjects(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("getting projects of the user: %w", err)
	}
	return projects, nil
}
