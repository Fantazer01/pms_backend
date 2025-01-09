package project

import (
	"context"
	"fmt"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/repository/interfaces"
	"time"

	"github.com/google/uuid"
)

type projectService struct {
	projectRepository interfaces.ProjectRepository
}

func NewProjectService(r interfaces.ProjectRepository) *projectService {
	return &projectService{
		projectRepository: r,
	}
}

func (s *projectService) GetProjectsPaged(ctx context.Context, pageInfo *model.PageInfo) ([]*model.ProjectShort, int, error) {
	projects, countProjects, err := s.projectRepository.GetProjectsPaged(ctx, pageInfo)
	if err != nil {
		return nil, 0, fmt.Errorf("getting projects: %w", err)
	}
	return projects, countProjects, nil
}

func (s *projectService) GetProjectByID(ctx context.Context, projectID string) (*model.Project, error) {
	project, err := s.projectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("getting project by id: %w", err)
	}
	return project, nil
}

func (s *projectService) CreateProject(ctx context.Context, insertProject *model.InsertProject) (*model.Project, error) {
	now := time.Now()
	project := &model.Project{
		ID:          uuid.NewString(),
		Name:        insertProject.Name,
		Description: insertProject.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	err := s.projectRepository.CreateProject(ctx, project)
	if err != nil {
		return nil, fmt.Errorf("creating project: %w", err)
	}
	return project, nil
}

func (s *projectService) UpdateProject(ctx context.Context, projectID string, updateProject *model.InsertProject) (*model.Project, error) {
	projectFromDb, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, err
	}
	if projectFromDb == nil {
		return nil, apperror.NotFound
	}
	project := &model.Project{
		ID:          projectID,
		Name:        updateProject.Name,
		Description: updateProject.Description,
		IsActive:    projectFromDb.IsActive,
		UpdatedAt:   time.Now(),
	}
	err = s.projectRepository.UpdateProject(ctx, project)
	if err != nil {
		return nil, fmt.Errorf("updating project: %w", err)
	}
	return project, nil
}

func (s *projectService) DeleteProject(ctx context.Context, projectID string) error {
	projectFromDb, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if projectFromDb == nil {
		return apperror.NotFound
	}
	err = s.projectRepository.DeleteProject(ctx, projectID)
	if err != nil {
		return fmt.Errorf("deleting project: %w", err)
	}
	return nil
}

func (s *projectService) GetArchivedProjectsPaged(ctx context.Context, pageInfo *model.PageInfo) ([]*model.ProjectShort, int, error) {
	projects, countProjects, err := s.projectRepository.GetArchivedProjectsPaged(ctx, pageInfo)
	if err != nil {
		return nil, 0, fmt.Errorf("getting archive projects: %w", err)
	}
	return projects, countProjects, nil
}

func (s *projectService) ArchiveProject(ctx context.Context, projectID string) error {
	projectFromDb, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if projectFromDb == nil {
		return apperror.NotFound
	}
	err = s.projectRepository.ArchiveProject(ctx, projectID)
	if err != nil {
		return fmt.Errorf("archiving project: %w", err)
	}
	return nil
}

func (s *projectService) UnarchiveProject(ctx context.Context, projectID string) error {
	projectFromDb, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if projectFromDb == nil {
		return apperror.NotFound
	}
	err = s.projectRepository.UnarchiveProject(ctx, projectID)
	if err != nil {
		return fmt.Errorf("unarchive project: %w", err)
	}
	return nil
}

func (s *projectService) GetProjectMembers(ctx context.Context, projectID string) ([]*model.UserShort, error) {
	projectFromDb, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, err
	}
	if projectFromDb == nil {
		return nil, apperror.NotFound
	}
	members, err := s.projectRepository.GetProjectMembers(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("getting project members: %w", err)
	}
	return members, nil
}

func (s *projectService) AddProjectMember(ctx context.Context, projectID string, member *model.Member) error {
	projectFromDb, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if projectFromDb == nil {
		return apperror.NotFound
	}
	err = s.projectRepository.AddProjectMember(ctx, projectID, member)
	if err != nil {
		return fmt.Errorf("adding member to project: %w", err)
	}
	return nil
}

func (s *projectService) DeleteProjectMember(ctx context.Context, projectID, userID string) error {
	projectFromDb, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if projectFromDb == nil {
		return apperror.NotFound
	}
	err = s.projectRepository.DeleteProjectMember(ctx, projectID, userID)
	if err != nil {
		return fmt.Errorf("deleting member from project: %w", err)
	}
	return nil
}

func (s *projectService) GetProjectTasks(ctx context.Context, projectID string) ([]*model.Task, error) {
	projectFromDb, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, err
	}
	if projectFromDb == nil {
		return nil, apperror.NotFound
	}
	tasks, err := s.projectRepository.GetProjectTasks(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("getting project tasks by project id: %w", err)
	}
	return tasks, nil
}
