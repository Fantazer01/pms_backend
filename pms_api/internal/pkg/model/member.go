package model

type Member struct {
	UserID         string `json:"user_id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	Role           string `json:"role"`
	IsAdminProject bool   `json:"is_admin_project"`
}

type MemberInserted struct {
	UserID         string `json:"user_id"`
	Role           string `json:"role"`
	IsAdminProject bool   `json:"is_admin_project"`
}
