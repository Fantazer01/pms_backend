package model

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProjectsPaged struct {
	PageIndex int        `json:"page_index"`
	PageSize  int        `json:"page_size"`
	Total     int        `json:"total"`
	Projects  []*Project `json:"items"`
}
