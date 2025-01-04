package task

import (
	"context"
	"fmt"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/repository/interfaces"
	"time"

	"github.com/google/uuid"
)

type taskService struct {
	taskRepository interfaces.TaskRepository
}

func NewTaskService(r interfaces.TaskRepository) *taskService {
	return &taskService{
		taskRepository: r,
	}
}

func (s *taskService) GetTaskByID(ctx context.Context, taskID string) (*model.Task, error) {
	task, err := s.taskRepository.GetTaskByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("getting task by id: %w", err)
	}
	return task, nil
}

func (s *taskService) CreateTask(ctx context.Context, t *model.TaskInserted) (*model.Task, error) {
	now := time.Now()
	task := &model.Task{
		ID:          uuid.NewString(),
		Name:        t.Name,
		Description: t.Description,
		ProjectID:   t.ProjectID,
		UserID:      t.UserID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	err := s.taskRepository.CreateTask(ctx, task)
	if err != nil {
		return nil, fmt.Errorf("creating task: %w", err)
	}
	return task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, taskID string, t *model.TaskInserted) (*model.Task, error) {
	taskFromDb, err := s.GetTaskByID(ctx, taskID)
	if err != nil {
		return nil, err
	}
	if taskFromDb == nil {
		return nil, apperror.NotFound
	}
	task := &model.Task{
		ID:          taskID,
		Name:        t.Name,
		Description: t.Description,
		ProjectID:   taskFromDb.ProjectID,
		UserID:      taskFromDb.UserID,
		CreatedAt:   taskFromDb.CreatedAt,
		UpdatedAt:   time.Now(),
	}
	err = s.taskRepository.UpdateTask(ctx, task)
	if err != nil {
		return nil, fmt.Errorf("updating task: %w", err)
	}
	return task, nil
}

func (s *taskService) DeleteTask(ctx context.Context, taskID string) error {
	taskFromDb, err := s.GetTaskByID(ctx, taskID)
	if err != nil {
		return err
	}
	if taskFromDb == nil {
		return apperror.NotFound
	}
	err = s.taskRepository.DeleteTask(ctx, taskID)
	if err != nil {
		return fmt.Errorf("deleting task: %w", err)
	}
	return nil
}
