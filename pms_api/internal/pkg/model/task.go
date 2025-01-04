package model

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID   string `json:"project_id"`
	UserID      string `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TaskInserted struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID   string `json:"project_id"`
	UserID      string `json:"user_id"`
}
