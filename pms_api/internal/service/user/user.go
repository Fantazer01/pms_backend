package user

import (
	"context"
	"fmt"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/pms_error"
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
	now := time.Now()
	user := &model.User{
		ID:         uuid.NewString(),
		Username:   u.Username,
		Email:      u.Email,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
		LastName:   u.LastName,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	err := s.userRepository.CreateUser(ctx, user)
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
		return nil, pms_error.NotFound
	}
	user := &model.User{
		ID:         userID,
		Username:   u.Username,
		Email:      u.Email,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
		LastName:   u.LastName,
		UpdatedAt:  time.Now(),
	}
	err = s.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("updating user: %w", err)
	}
	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, userID string) error {
	err := s.userRepository.DeleteUser(ctx, userID)
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
