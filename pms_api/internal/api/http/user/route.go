package user

import "github.com/labstack/echo/v4"

func (h *handler) RegisterRoutes(router *echo.Group) {
	users := router.Group("/users")
	{
		users.GET("", h.GetUsers)
		users.GET("/:id", h.GetUserByID)
		users.POST("", h.CreateUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
		users.GET("/:id/projects", h.GetUserProjects)
	}
}
