package project

import "github.com/labstack/echo/v4"

func (h *handler) RegirterRoutes(router *echo.Group) {
	projects := router.Group("/projects")
	{
		projects.GET("", h.GetProjects)
		projects.GET("/:id", h.GetProjectByID)
		projects.POST("", h.CreateProject)
		projects.PUT("/:id", h.UpdateProject)
		projects.DELETE("/:id", h.DeleteProject)

		projects.GET("/:id/members", h.GetProjectMembers)
		projects.POST("/:id/members/:user_id", h.AddProjectMember)
		projects.DELETE("/:id/members/:user_id", h.DeleteProjectMember)

		projects.GET("/:id/tasks", h.GetProjectTasks)
	}
}
