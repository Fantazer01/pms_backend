package postgres

const (
	countProjectsQuery = `
		SELECT count(id)
		FROM project
		WHERE is_active = true
	`
	getProjectsQuery = `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM public.project
		WHERE is_active = true
		ORDER BY name OFFSET @offset LIMIT @page_size
	`
	getProjectByID = `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM public.project
		WHERE id = @project_id
	`
	createProject = `
		INSERT INTO project (id, name, description, is_active, created_at, updated_at)
		VALUES (@id, @name, @description, true, @created_at, @updated_at)
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
		SELECT id, name, description, is_active, created_at, updated_at
		FROM public.project
		WHERE is_active = false
		ORDER BY name OFFSET @offset LIMIT @page_size
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
	getProjectMembers = `
		SELECT u.id, u.login, u.first_name, u.middle_name, u.last_name, r.name_role, pp.is_admin_project
		FROM users u
		JOIN participants_project pp on pp.user_id = u.id
		JOIN role r on r.id = pp.role_id
		WHERE pp.project_id = @project_id
	`
	getProjectMemberByID = `
		SELECT u.id, u.login, u.first_name, u.middle_name, u.last_name, r.name_role, pp.is_admin_project
		FROM participants_project pp
		JOIN users u on u.id = pp.user_id
		JOIN role r on r.id = pp.role_id
		WHERE pp.project_id = @project_id AND pp.user_id = @user_id
	`
	insertMemberToProject = `
		INSERT INTO participants_project(user_id, project_id, role_id, is_admin_project)
		VALUES (@user_id, @project_id, @role_id, @is_admin_project)
	`
	updateMemberToProject = `
		UPDATE participants_project
		SET 
		role_id = @role_id,
		is_admin_project = @is_admin_project
		WHERE project_id = @project_id AND user_id = @user_id
	`
	deleteUserFromProject = `
		DELETE FROM participants_project
		WHERE project_id = @project_id AND user_id = @user_id
	`
	getProjectTasks = `
		SELECT id, name, description, status, project_id, author_id, executor_id, tester_id, created_at, deadline
		FROM task
		WHERE project_id = @project_id
	`
)
