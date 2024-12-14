package postgres

const (
	countProjectsQuery = `
		SELECT count(id)
		FROM project
		WHERE is_active = true
	`
	getProjectsQuery = `
		SELECT id, name, description, created_at, updated_at
		FROM public.project
		WHERE is_active = true
	`
	getProjectByID = `
		SELECT id, name, description, created_at, updated_at
		FROM public.project
		WHERE id = @project_id
	`
	createProject = `
		INSERT INTO project (id, name, description, created_at, updated_at)
		VALUES (@id, @name, @description, @created_at, @updated_at)
	`
	updateProject = `
		UPDATE project
		SET
		name = @name,
		description = @description,
		updated_at = @updated_at
		WHERE id = @id
	`
	deleteProject = `
		DELETE FROM project
		WHERE id = @id
	`
	countArchiveProjectsQuery = `
		SELECT count(id)
		FROM project
		WHERE is_active = false
	`
	getArchiveProjectsQuery = `
		SELECT id, name, description, created_at, updated_at
		FROM public.project
		WHERE is_active = false
	`
	archiveProject = `
		UPDATE project
		SET is_active = false
		WHERE id = @id
	`
	unarchiveProject = `
		UPDATE project
		SET is_active = true
		WHERE id = @id
	`
)
