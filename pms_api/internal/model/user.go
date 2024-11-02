package model

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UsersPaged struct {
	PageIndex int     `json:"page_index"`
	PageSize  int     `json:"page_size"`
	Total     int     `json:"total"`
	Users     []*User `json:"items"`
}
