package project

import "github.com/labstack/echo/v4"

func (h *handler) RegirterRoutes(router *echo.Group) {
	projects := router.Group("/projects")
	{
		projects.GET("", h.GetProjects)
		projects.GET("/:project_id", h.GetProjectByID)
		projects.POST("", h.CreateProject)
		projects.PUT("/:project_id", h.UpdateProject)
		projects.DELETE("/:project_id", h.DeleteProject)

		projects.GET("/archived", h.GetArchivedProjects)
		projects.PUT("/:project_id/archive", h.ArchiveProject)
		projects.PUT("/:project_id/unarchive", h.UnarchiveProject)

		projects.GET("/:project_id/members", h.GetProjectMembers)
		projects.POST("/:project_id/members/:user_id", h.AddProjectMember)
		projects.DELETE("/:project_id/members/:user_id", h.DeleteProjectMember)

		projects.GET("/:project_id/tasks", h.GetProjectTasks)
	}
}
