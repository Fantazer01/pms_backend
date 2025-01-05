package auth

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterRoutes(e *echo.Group) {
	e.POST("/login", h.Login)
	e.POST("/refresh", h.Refresh)
	e.POST("/logout", h.Logout)
}
