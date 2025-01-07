package profile

import "github.com/labstack/echo/v4"

func (h *handler) RegisterRoutes(router *echo.Group) {
	router.GET("/profile", h.GetProfile)
}
