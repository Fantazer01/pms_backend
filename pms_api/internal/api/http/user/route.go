package user

import "github.com/labstack/echo/v4"

func (h *handler) RegisterRoutes(router *echo.Group) {
	users := router.Group("/users")
	{
		users.GET("", h.GetUsers)
		users.GET("/:user_id", h.GetUserByID)
		users.POST("", h.CreateUser)
		users.PUT("/:user_id", h.UpdateUser)
		users.DELETE("/:user_id", h.DeleteUser)
		users.GET("/:user_id/projects", h.GetUserProjects)
	}
}
