package postgres

const (
	countUsersQuery = `
		SELECT COUNT(id)
		FROM users
	`
	getUsersQuery = `
		SELECT id, login, first_name, middle_name, last_name
		FROM users
		ORDER BY login LIMIT @page_size OFFSET @offset
	`
	getUserByID = `
		SELECT id, login, is_admin, first_name, middle_name, last_name, position, created_at, updated_at
		FROM users
		WHERE id = @id
	`
	createUser = `
		INSERT INTO users(id, login, password, is_admin, first_name, middle_name, last_name, position, created_at, updated_at)
		VALUES (@id, @login, @password, @is_admin, @first_name, @middle_name, @last_name, @position, @created_at, @updated_at)
	`
	updateUser = `
		UPDATE users
		SET
		login = @login,
		first_name = @first_name,
		middle_name = @middle_name,
		last_name = @last_name,
		position = @position,
		updated_at = @updated_at
		WHERE id = @id
	`
	deleteUser = `
		DELETE FROM users
		WHERE id = @id
	`
	getProjectsOfUserQuery = `
		SELECT p.id, p.name
		FROM project p
		JOIN participants_project pp on pp.project_id = p.id
		WHERE pp.user_id = @user_id
	`
)
