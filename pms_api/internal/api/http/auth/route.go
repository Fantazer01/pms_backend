package auth

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group) {
	h := NewHandler()
	e.POST("/login", h.Login)
	e.POST("/refresh", h.Refresh)
	e.POST("/logout", h.Logout)
}
