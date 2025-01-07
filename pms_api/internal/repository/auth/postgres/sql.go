package postgres

const (
	getUserByUsername = `
		SELECT id, login, password, first_name, middle_name, last_name, position, created_at, updated_at
		FROM users
		WHERE login = @username
	`
)
