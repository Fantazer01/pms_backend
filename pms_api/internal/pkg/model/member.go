package model

type Member struct {
	UserID         string `json:"user_id"`
	Role           string `json:"role"`
	IsAdminProject bool   `json:"is_admin_project"`
}
