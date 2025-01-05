package model

import "time"

type Task struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	ProjectID   string    `json:"project_id"`
	AuthorID    string    `json:"author_id"`
	ExecutorID  string    `json:"executor_id"`
	TesterID    string    `json:"tester_id"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
}

type TaskInserted struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	ProjectID   string    `json:"project_id"`
	AuthorID    string    `json:"author_id"`
	ExecutorID  string    `json:"executor_id"`
	TesterID    string    `json:"tester_id"`
	Deadline    time.Time `json:"deadline"`
}
