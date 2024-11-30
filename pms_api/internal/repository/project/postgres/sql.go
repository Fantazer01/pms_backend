package postgres

const (
	getProjectsQuery = `
		SELECT id, name, description, created_at, updated_at
		FROM public.project;
	`
	getProjectByID = `
		SELECT id, name, description, created_at, updated_at
		FROM public.project
		WHERE id = @project_id;
	`
	createProject = `
		INSERT INTO project (id, name, description, created_at, updated_at)
		VALUES (@id, @name, @description, @created_at, @updated_at)
	`
)
