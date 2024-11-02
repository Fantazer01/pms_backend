package task

import "github.com/labstack/echo/v4"

func (h *handler) RegirterRoutes(router *echo.Group) {
	task := router.Group("/task")
	{
		task.GET("/:id", h.GetTaskByID)
		task.POST("", h.CreateTask)
		task.PUT("/:id", h.UpdateTask)
	}
}
