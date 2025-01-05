package postgres

import (
	"context"
	"errors"
	"pms_backend/pms_api/internal/pkg/model"

	"github.com/jackc/pgx/v5"
)

func (r *repository) GetTaskByID(ctx context.Context, taskID string) (*model.Task, error) {
	rows, err := r.pool.Query(ctx, getTaskByID, pgx.NamedArgs{"task_id": taskID})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	task, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[task])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &model.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		ProjectID:   task.ProjectID,
		AuthorID:    task.AuthorID,
		ExecutorID:  task.ExecutorID,
		TesterID:    task.TesterID,
		CreatedAt:   task.CreatedAt,
		Deadline:    task.Deadline,
	}, nil
}

func (r *repository) CreateTask(ctx context.Context, t *model.Task) error {
	_, err := r.pool.Exec(ctx, createTask, pgx.NamedArgs{
		"id":          t.ID,
		"name":        t.Name,
		"description": t.Description,
		"status":      t.Status,
		"project_id":  t.ProjectID,
		"author_id":   t.AuthorID,
		"executor_id": t.ExecutorID,
		"tester_id":   t.TesterID,
		"created_at":  t.CreatedAt,
		"deadline":    t.Deadline,
	})
	return err
}

func (r *repository) UpdateTask(ctx context.Context, t *model.Task) error {
	_, err := r.pool.Exec(ctx, updateTask, pgx.NamedArgs{
		"id":          t.ID,
		"name":        t.Name,
		"description": t.Description,
		"status":      t.Status,
		"author_id":   t.AuthorID,
		"executor_id": t.ExecutorID,
		"tester_id":   t.TesterID,
		"deadline":    t.Deadline,
	})
	return err
}

func (r *repository) DeleteTask(ctx context.Context, taskID string) error {
	_, err := r.pool.Exec(ctx, deleteTask, pgx.NamedArgs{"task_id": taskID})
	return err
}
