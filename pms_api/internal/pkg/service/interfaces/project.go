package interfaces

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
)

type ProjectService interface {
	GetProjectsPaged(ctx context.Context, pageInfo *model.PageInfo) ([]*model.Project, error)
	GetProjectByID(ctx context.Context, projectID string) (model.Project, error)
	CreateProject(ctx context.Context, project *model.InsertProject) error
	UpdateProject(ctx context.Context, projectID string, project *model.InsertProject) error
	DeleteProject(ctx context.Context, projectID string) error

	GetArchivedProjectsPaged(ctx context.Context, pageInfo *model.PageInfo) ([]*model.Project, error)
	ArchiveProject(ctx context.Context, projectID string) error
	UnarchiveProject(ctx context.Context, projectID string) error

	GetProjectMembers(ctx context.Context, projectID string) ([]*model.User, error)
	AddProjectMember(ctx context.Context, projectID, userID string) error
	DeleteProjectMember(ctx context.Context, projectID, userID string) error

	GetProjectTasks(ctx context.Context, projectID string) ([]*model.Task, error)
}