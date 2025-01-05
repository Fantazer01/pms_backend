package interfaces

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

type TaskRepository interface {
	GetTaskByID(ctx context.Context, taskID string) (*model.Task, error)
	CreateTask(ctx context.Context, task *model.Task) error
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, taskID string) error
}
