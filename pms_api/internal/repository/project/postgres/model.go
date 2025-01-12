package postgres

import (
	"pms_backend/pms_api/internal/pkg/model"
	"strings"
	"time"
)

type project struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	IsActive    bool      `db:"is_active"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func toProjectShortsFromDb(projectsFromDb []project) []*model.ProjectShort {
	projects := make([]*model.ProjectShort, len(projectsFromDb))
	for i := range projects {
		projects[i] = &model.ProjectShort{
			ID:       projectsFromDb[i].ID,
			Name:     projectsFromDb[i].Name,
			IsActive: projectsFromDb[i].IsActive,
		}
	}
	return projects
}

func toProjectFromDb(projectFromDb project) *model.Project {
	return &model.Project{
		ID:          projectFromDb.ID,
		Name:        projectFromDb.Name,
		Description: projectFromDb.Description,
		IsActive:    projectFromDb.IsActive,
		CreatedAt:   projectFromDb.CreatedAt,
		UpdatedAt:   projectFromDb.UpdatedAt,
	}
}

type member struct {
	ID             string `db:"id"`
	Username       string `db:"login"`
	FirstName      string `db:"first_name"`
	MiddleName     string `db:"middle_name"`
	LastName       string `db:"last_name"`
	Role           string `db:"name_role"`
	IsAdminProject bool   `db:"is_admin_project"`
}

type role struct {
	ID   int    `db:"id"`
	Name string `db:"name_role"`
}

func toMembersFromRepo(usersFromDb []member) []*model.Member {
	users := make([]*model.Member, len(usersFromDb))
	for i := range users {
		fullNameSb := strings.Builder{}
		fullNameSb.WriteString(usersFromDb[i].FirstName)
		if usersFromDb[i].MiddleName != "" {
			fullNameSb.WriteString(" ")
			fullNameSb.WriteString(usersFromDb[i].MiddleName)
		}
		if usersFromDb[i].LastName != "" {
			fullNameSb.WriteString(" ")
			fullNameSb.WriteString(usersFromDb[i].LastName)
		}
		users[i] = &model.Member{
			UserID:         usersFromDb[i].ID,
			Username:       usersFromDb[i].Username,
			FullName:       fullNameSb.String(),
			Role:           usersFromDb[i].Role,
			IsAdminProject: usersFromDb[i].IsAdminProject,
		}
	}
	return users
}

type task struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	ProjectID   string    `db:"project_id"`
	AuthorID    string    `db:"author_id"`
	ExecutorID  string    `db:"executor_id"`
	TesterID    string    `db:"tester_id"`
	CreatedAt   time.Time `db:"created_at"`
	Deadline    time.Time `db:"deadline"`
}

func toTasksFromRepo(tasksFromDb []task) []*model.Task {
	tasks := make([]*model.Task, len(tasksFromDb))
	for i := range tasks {
		tasks[i] = &model.Task{
			ID:          tasksFromDb[i].ID,
			Name:        tasksFromDb[i].Name,
			Description: tasksFromDb[i].Description,
			Status:      tasksFromDb[i].Status,
			ProjectID:   tasksFromDb[i].ProjectID,
			AuthorID:    tasksFromDb[i].AuthorID,
			ExecutorID:  tasksFromDb[i].ExecutorID,
			TesterID:    tasksFromDb[i].TesterID,
			CreatedAt:   tasksFromDb[i].CreatedAt,
			Deadline:    tasksFromDb[i].Deadline,
		}
	}
	return tasks
}
