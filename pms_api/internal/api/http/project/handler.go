package project

import (
	"errors"
	"log/slog"
	"net/http"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/service/interfaces"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	internalServerError = "Internal server error"
	incorrectProjectID  = "Incorrect project id"
	bindError           = "Bind error"
	incorrectUserID     = "Incorrect user id"
	projectNotFound     = "Project not found"
)

type handler struct {
	projectService interfaces.ProjectService
}

func NewHandler(s interfaces.ProjectService) *handler {
	return &handler{
		projectService: s,
	}
}

// GetProjects
// @Tags Project
// @Summary Get projects
// @Description Get projects
// @Accept json
// @Produce json
// @Param pageIndex query string false "Page index of projects"
// @Param pageSize query string false "Page size of projects"
// @Success 200 {object} model.ProjectsPaged
// @Failure 500 {object} model.Message
// @Security Login
// @Router /projects [get]
func (h *handler) GetProjects(c echo.Context) error {
	qp := c.QueryParam("pageIndex")
	pageIndex, err := strconv.Atoi(qp)
	if err != nil || pageIndex < 1 {
		pageIndex = 1
	}
	qs := c.QueryParam("pageSize")
	pageSize, err := strconv.Atoi(qs)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	pageInfo := &model.PageInfo{
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
	projects, countProjects, err := h.projectService.GetProjectsPaged(c.Request().Context(), pageInfo)
	if err != nil {
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.JSON(http.StatusOK,
		model.ProjectsPaged{
			PageIndex: pageIndex,
			PageSize:  pageSize,
			Total:     countProjects,
			Projects:  projects,
		})
}

// GetProjectByID
// @Tags Project
// @Summary Get project by ID
// @Description Get project by ID
// @Accept json
// @Produce json
// @Param project_id path string true "Project id"
// @Success 200 {object} model.Project
// @Failure 404 {object} model.Message "Project not found"
// @Failure 422 {object} model.Message "Incorrect id of project"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id} [get]
func (h *handler) GetProjectByID(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	project, err := h.projectService.GetProjectByID(c.Request().Context(), projectID)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.JSON(http.StatusOK, project)
}

// CreateProject
// @Tags Project
// @Summary Create project
// @Description Create project
// @Accept json
// @Produce json
// @Param project body model.InsertProject true "Project"
// @Success 201 {object} model.Project
// @Failure 422 {object} model.Message "Bind error"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects [post]
func (h *handler) CreateProject(c echo.Context) error {
	insertProject := &model.InsertProject{}
	err := c.Bind(insertProject)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: bindError})
	}
	project, err := h.projectService.CreateProject(c.Request().Context(), insertProject)
	if err != nil {
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.JSON(http.StatusCreated, project)
}

// UpdateProject
// @Tags Project
// @Summary Update project
// @Description Update project
// @Accept json
// @Produce json
// @Param project_id path string true "Project id"
// @Param project body model.InsertProject true "Project"
// @Success 200 {object} model.Project
// @Failure 404 {object} model.Message "Project not found"
// @Failure 422 {object} model.Message "Incorrect id of project/Bind error"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id} [put]
func (h *handler) UpdateProject(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	updateProject := &model.InsertProject{}
	err := c.Bind(updateProject)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: bindError})
	}
	project, err := h.projectService.UpdateProject(c.Request().Context(), projectID, updateProject)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.JSON(http.StatusOK, project)
}

// DeleteProject
// @Tags Project
// @Summary Delete project
// @Description Delete project
// @Accept  json
// @Produce  json
// @Param project_id path string true "Project id"
// @Success 204
// @Failure 404 {object} model.Message "Project not found"
// @Failure 422 {object} model.Message "Incorrect id of project"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id} [delete]
func (h *handler) DeleteProject(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	err := h.projectService.DeleteProject(c.Request().Context(), projectID)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.NoContent(http.StatusNoContent)
}

// GetArchivedProjects
// @Tags Project
// @Summary Get archived projects
// @Description Get archived projects
// @Produce json
// @Param pageIndex query string false "Page index of projects"
// @Param pageSize query string false "Page size of projects"
// @Success 200 {object} model.ProjectsPaged
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/archived [get]
func (h *handler) GetArchivedProjects(c echo.Context) error {
	qp := c.QueryParam("pageIndex")
	pageIndex, err := strconv.Atoi(qp)
	if err != nil || pageIndex < 1 {
		pageIndex = 1
	}
	qs := c.QueryParam("pageSize")
	pageSize, err := strconv.Atoi(qs)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	pageInfo := &model.PageInfo{
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
	projects, countProjects, err := h.projectService.GetArchivedProjectsPaged(c.Request().Context(), pageInfo)
	if err != nil {
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.JSON(http.StatusOK,
		model.ProjectsPaged{
			PageIndex: pageIndex,
			PageSize:  pageSize,
			Total:     countProjects,
			Projects:  projects,
		})
}

// ArchiveProject
// @Tags Project
// @Summary Archive the project by id
// @Description Archive the project by id
// @Produce json
// @Param project_id path string true "Project id"
// @Success 204
// @Failure 404 {object} model.Message "Project not found"
// @Failure 422 {object} model.Message "Incorrect id of project"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id}/archive [put]
func (h *handler) ArchiveProject(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	err := h.projectService.ArchiveProject(c.Request().Context(), projectID)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.NoContent(http.StatusNoContent)
}

// UnarchiveProject
// @Tags Project
// @Summary Unarchive the project by id
// @Description Unarchive the project by id
// @Produce json
// @Param project_id path string true "Project id"
// @Success 204
// @Failure 404 {object} model.Message "Project not found"
// @Failure 422 {object} model.Message "Incorrect id of project"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id}/unarchive [put]
func (h *handler) UnarchiveProject(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	err := h.projectService.UnarchiveProject(c.Request().Context(), projectID)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.NoContent(http.StatusNoContent)
}

// GetProjectMembers
// @Tags Project
// @Summary Get project members
// @Description Get project members
// @Accept json
// @Produce json
// @Param project_id path string true "Project id"
// @Success 200 {object} []model.UserShort
// @Failure 404 {object} model.Message "Project not found"
// @Failure 422 {object} model.Message "Incorrect id of project"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id}/members [get]
func (h *handler) GetProjectMembers(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	users, err := h.projectService.GetProjectMembers(c.Request().Context(), projectID)
	if err != nil {
		slog.Error(err.Error())
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.JSON(http.StatusOK, users)
}

// AddProjectMember
// @Tags Project
// @Summary Add project member
// @Description Add project member
// @Accept json
// @Produce json
// @Param project_id path string true "Project id"
// @Param project_member body model.Member true "Project member"
// @Success 204
// @Failure 404 {object} model.Message "Project not found/User not found"
// @Failure 422 {object} model.Message "Incorrect id of project/Incorrect id of user"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id}/members/{user_id} [post]
func (h *handler) AddProjectMember(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	projectMember := &model.Member{}
	err := c.Bind(projectMember)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: bindError})
	}
	if err := uuid.Validate(projectMember.UserID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectUserID})
	}
	err = h.projectService.AddProjectMember(c.Request().Context(), projectID, projectMember)
	if err != nil {
		slog.Error(err.Error())
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.NoContent(http.StatusNoContent)
}

// DeleteProjectMember
// @Tags Project
// @Summary Delete project member
// @Description Delete project member
// @Accept json
// @Produce json
// @Param project_id path string true "Project id"
// @Param user_id path string true "User id"
// @Success 204
// @Failure 404 {object} model.Message "Project not found/User not found"
// @Failure 422 {object} model.Message "Incorrect id of project/Incorrect id of user"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id}/members/{user_id} [delete]
func (h *handler) DeleteProjectMember(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	userID := c.Param("user_id")
	if err := uuid.Validate(userID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectUserID})
	}
	err := h.projectService.DeleteProjectMember(c.Request().Context(), projectID, userID)
	if err != nil {
		slog.Error(err.Error())
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.NoContent(http.StatusNoContent)
}

// GetProjectTasks
// @Tags Project
// @Summary Get project tasks
// @Description Get project tasks
// @Produce json
// @Param project_id path string true "Project id"
// @Success 200 {object} []model.Task
// @Failure 404 {object} model.Message "Project not found"
// @Failure 422 {object} model.Message "Incorrect id of project"
// @Failure 500 {object} model.Message "Internal server error"
// @Security Login
// @Router /projects/{project_id}/tasks [get]
func (h *handler) GetProjectTasks(c echo.Context) error {
	projectID := c.Param("project_id")
	if err := uuid.Validate(projectID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectProjectID})
	}
	tasks, err := h.projectService.GetProjectTasks(c.Request().Context(), projectID)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: projectNotFound})
		}
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.JSON(http.StatusOK, tasks)
}
