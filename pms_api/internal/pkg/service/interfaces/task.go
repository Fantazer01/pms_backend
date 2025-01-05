package interfaces

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

type TaskService interface {
	GetTaskByID(ctx context.Context, taskID string) (*model.Task, error)
	CreateTask(ctx context.Context, task *model.TaskInserted) (*model.Task, error)
	UpdateTask(ctx context.Context, taskID string, task *model.TaskInserted) (*model.Task, error)
	DeleteTask(ctx context.Context, taskID string) error
}
