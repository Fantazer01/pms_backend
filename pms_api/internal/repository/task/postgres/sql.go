package postgres

const (
	getTaskByID = `
		SELECT id, name, description, status, project_id, author_id, executor_id, tester_id, created_at, deadline
		FROM task
		WHERE id = @id
	`
	createTask = `
		INSERT INTO task (id, name, description, status, project_id, author_id, executor_id, tester_id, created_at, deadline)
		VALUES (@id, @name, @description, @status, @project_id, @author_id, @executor_id, @tester_id, @created_at, @deadline)
	`
	updateTask = `
		UPDATE task
		SET
		name = @name,
		description = @description,
		author_id = @author_id,
		executor_id = @executor_id,
		tester_id = @tester_id,
		deadline = @deadline
		WHERE id = @id
	`
	deleteTask = `
		DELETE FROM task
		WHERE id = @id
	`
)
