package model

import "time"

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserShort struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type UserInserted struct {
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type UsersPaged struct {
	PageIndex int          `json:"page_index"`
	PageSize  int          `json:"page_size"`
	Total     int          `json:"total"`
	Users     []*UserShort `json:"items"`
}
