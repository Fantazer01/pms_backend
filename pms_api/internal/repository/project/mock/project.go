package mock

import (
	"context"
	"pms_backend/pms_api/internal/pkg/model"
	"time"
)

type mockProjectRepository struct {
}

func NewMockProjectRepository() *mockProjectRepository {
	return &mockProjectRepository{}
}

func (r *mockProjectRepository) GetProjectsPaged(ctx context.Context, pageInfo *model.PageInfo) ([]*model.Project, error) {
	return []*model.Project{
		{
			ID:          "d3ac397c-537d-4cf3-bff7-c2c371535b1b",
			Name:        "Test project",
			Description: "Test project",
			CreatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
			UpdatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
		},
		{
			ID:          "58184169-f297-4b26-8fbc-91dfc8fb8468",
			Name:        "Test project 2",
			Description: "Test project 2",
			CreatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
			UpdatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
		},
	}, nil
}

func (r *mockProjectRepository) GetProjectByID(ctx context.Context, projectID string) (*model.Project, error) {
	return &model.Project{
		ID:          "d3ac397c-537d-4cf3-bff7-c2c371535b1b",
		Name:        "Test project",
		Description: "Test project",
		CreatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
		UpdatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
	}, nil
}

func (r *mockProjectRepository) CreateProject(ctx context.Context, project *model.Project) error {
	return nil
}

func (r *mockProjectRepository) UpdateProject(ctx context.Context, project *model.Project) error {
	return nil
}

func (r *mockProjectRepository) DeleteProject(ctx context.Context, projectID string) error {
	return nil
}

func (r *mockProjectRepository) GetArchivedProjectsPaged(ctx context.Context, pageInfo *model.PageInfo) ([]*model.Project, error) {
	return []*model.Project{
		{
			ID:          "d3ac397c-537d-4cf3-bff7-c2c371535b1b",
			Name:        "Test project",
			Description: "Test project",
			CreatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
			UpdatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
		},
		{
			ID:          "58184169-f297-4b26-8fbc-91dfc8fb8468",
			Name:        "Test project 2",
			Description: "Test project 2",
			CreatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
			UpdatedAt:   time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local),
		},
	}, nil
}

func (r *mockProjectRepository) ArchiveProject(ctx context.Context, projectID string) error {
	return nil
}

func (r *mockProjectRepository) UnarchiveProject(ctx context.Context, projectID string) error {
	return nil
}

func (r *mockProjectRepository) GetProjectMembers(ctx context.Context, projectID string) ([]*model.User, error) {
	return []*model.User{
		{
			ID:         "97bbc8bc-cf27-4e16-9cb3-a77235cec203",
			Username:   "test_project_admin",
			Email:      "testadmin@pms.ru",
			FirstName:  "Иван",
			MiddleName: "Иванович",
			LastName:   "Иванов",
			CreatedAt:  time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local).String(),
			UpdatedAt:  time.Date(2024, 11, 19, 12, 25, 0, 0, time.Local).String(),
		},
	}, nil
}

func (r *mockProjectRepository) AddProjectMember(ctx context.Context, projectID, userID string) error {
	return nil
}

func (r *mockProjectRepository) DeleteProjectMember(ctx context.Context, projectID, userID string) error {
	return nil
}

func (r *mockProjectRepository) GetProjectTasks(ctx context.Context, projectID string) ([]*model.Task, error) {
	return []*model.Task{
		{},
	}, nil
}
