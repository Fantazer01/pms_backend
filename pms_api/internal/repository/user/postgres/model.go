package postgres

import "time"

type userShort struct {
	ID         string `db:"id"`
	Username   string `db:"login"`
	IsAdmin    bool   `db:"is_admin"`
	FirstName  string `db:"first_name"`
	MiddleName string `db:"middle_name"`
	LastName   string `db:"last_name"`
}

type user struct {
	ID         string    `db:"id"`
	Username   string    `db:"login"`
	IsAdmin    bool      `db:"is_admin"`
	FirstName  string    `db:"first_name"`
	MiddleName string    `db:"middle_name"`
	LastName   string    `db:"last_name"`
	Position   string    `db:"position"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type projectShort struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	IsActive bool   `db:"is_active"`
}
