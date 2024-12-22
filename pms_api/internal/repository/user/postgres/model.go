package postgres

import "time"

type userShort struct {
	ID         string `db:"id"`
	Username   string `db:"login"`
	FirstName  string `db:"first_name"`
	MiddleName string `db:"middle_name"`
	LastName   string `db:"last_name"`
}

type user struct {
	ID         string    `db:"id"`
	Username   string    `db:"login"`
	FirstName  string    `db:"first_name"`
	MiddleName string    `db:"middle_name"`
	LastName   string    `db:"last_name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type projectShort struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
