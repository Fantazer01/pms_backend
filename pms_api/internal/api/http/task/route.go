package task

import "github.com/labstack/echo/v4"

func (h *handler) RegisterRoutes(router *echo.Group) {
	task := router.Group("/task")
	{
		task.GET("/:task_id", h.GetTaskByID)
		task.POST("", h.CreateTask)
		task.PUT("/:task_id", h.UpdateTask)
		task.DELETE("/:task_id", h.DeleteTask)
	}
}
