package project

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

// GetProjects
// @Tags Project
// @Summary Get projects
// @Description Get projects
// @Accept json
// @Produce json
// @Success 200 {object} model.ProjectsPaged
// @Router /projects [get]
func (h *handler) GetProjects(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// GetProjectByID
// @Tags Project
// @Summary Get project by ID
// @Description Get project by ID
// @Accept json
// @Produce json
// @Success 200 {object} model.Project
// @Router /projects/{id} [get]
func (h *handler) GetProjectByID(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// CreateProject
// @Tags Project
// @Summary Create project
// @Description Create project
// @Accept json
// @Produce json
// @Success 201 {object} model.Project
// @Router /projects [post]
func (h *handler) CreateProject(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// UpdateProject
// @Tags Project
// @Summary Update project
// @Description Update project
// @Accept json
// @Produce json
// @Success 200 {object} model.Project
// @Router /projects/{id} [put]
func (h *handler) UpdateProject(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// DeleteProject
// @Tags Project
// @Summary Delete project
// @Description Delete project
// @Accept  json
// @Produce  json
// @Success 204
// @Failure 400
// @Router /projects/{id} [delete]
func (h *handler) DeleteProject(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// GetProjectMembers
// @Tags Project
// @Summary Get project members
// @Description Get project members
// @Accept json
// @Produce json
// @Success 200 {object} model.UsersPaged
// @Router /projects/{id}/members [get]
func (h *handler) GetProjectMembers(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// AddProjectMember
// @Tags Project
// @Summary Add project member
// @Description Add project member
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /projects/{project_id}/members/{user_id} [post]
func (h *handler) AddProjectMember(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// DeleteProjectMember
// @Tags Project
// @Summary Delete project member
// @Description Delete project member
// @Accept json
// @Produce json
// @Success 204
// @Failure 400
// @Router /projects/{id}/members/{user_id} [delete]
func (h *handler) DeleteProjectMember(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// GetProjectTasks
// @Tags Project
// @Summary Get project tasks
// @Description Get project tasks
// @Accept json
// @Produce json
// @Success 200 {object} model.Tasks
// @Router /projects/{id}/tasks [get]
func (h *handler) GetProjectTasks(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
