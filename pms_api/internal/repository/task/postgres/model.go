package postgres

import "time"

type task struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	ProjectID   string    `db:"project_id"`
	AuthorID    string    `db:"author_id"`
	ExecutorID  string    `db:"executor_id"`
	TesterID    string    `db:"tester_id"`
	CreatedAt   time.Time `db:"created_at"`
	Deadline    time.Time `db:"deadline"`
}
