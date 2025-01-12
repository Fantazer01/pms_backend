package task

import (
	"context"
	"errors"
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
	if time.Since(t.Deadline) > 0 {
		return nil, errors.New("trying to establish the deadline's elapsed time")
	}
	task := &model.Task{
		ID:          uuid.NewString(),
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
		ProjectID:   t.ProjectID,
		AuthorID:    t.AuthorID,
		ExecutorID:  t.ExecutorID,
		TesterID:    t.TesterID,
		CreatedAt:   time.Now(),
		Deadline:    t.Deadline,
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
	if time.Since(t.Deadline) > 0 {
		return nil, errors.New("trying to establish the deadline's elapsed time")
	}
	task := &model.Task{
		ID:          taskID,
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
		ProjectID:   taskFromDb.ProjectID,
		AuthorID:    taskFromDb.AuthorID,
		ExecutorID:  taskFromDb.ExecutorID,
		TesterID:    taskFromDb.TesterID,
		CreatedAt:   taskFromDb.CreatedAt,
		Deadline:    t.Deadline,
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
