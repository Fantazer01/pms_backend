package model

import "time"

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProjectShort struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type InsertProject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectsPaged struct {
	PageIndex int             `json:"page_index"`
	PageSize  int             `json:"page_size"`
	Total     int             `json:"total"`
	Projects  []*ProjectShort `json:"items"`
}
