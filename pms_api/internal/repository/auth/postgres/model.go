package postgres

import "time"

type user struct {
	ID         string    `db:"id"`
	Username   string    `db:"login"`
	Password   []byte    `db:"password"`
	IsAdmin    bool      `db:"is_admin"`
	FirstName  string    `db:"first_name"`
	MiddleName string    `db:"middle_name"`
	LastName   string    `db:"last_name"`
	Position   string    `db:"position"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
