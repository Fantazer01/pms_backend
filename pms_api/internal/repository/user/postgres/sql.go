package postgres

const (
	countUsersQuery = `
		SELECT COUNT(id)
		FROM users
	`
	getUsersQuery = `
		SELECT id, login, first_name, middle_name, last_name
		FROM users
		ORDER BY login OFFSET @offset LIMIT @page_size
	`
	getUserByID = `
		SELECT id, login, email, first_name, middle_name, last_name, created_at, updated_at
		FROM users
		WHERE id = @id
	`
	createUser = `
		INSERT INTO users(id, login, email, first_name, middle_name, last_name, created_at, updated_at)
		VALUES (@id, @login, @email, @first_name, @middle_name, @last_name, @created_at, @updated_at)
	`
	updateUser = `
		UPDATE users
		SET
		login = @login,
		email = @email,
		first_name = @first_name,
		middle_name = @middle_name,
		last_name = @last_name,
		updated_at = @updated_at
		WHERE id = @id
	`
	deleteUser = `
		DELETE FROM users
		WHERE id = @id
	`
	getProjectsOfUserQuery = `
	`
)
