package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

// Login
// @Tags Auth
// @Summary Login
// @Description Login
// @Accept json
// @Produce json
// @Success 200 {object} model.Tokens
// @Router /login [post]
func (h *handler) Login(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// Refresh
// @Tags Auth
// @Summary Refresh
// @Description Refresh
// @Accept json
// @Produce json
// @Success 200 {object} model.Tokens
// @Router /refresh [post]
func (h *handler) Refresh(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// Logout
// @Tags Auth
// @Summary Logout
// @Description Logout
// @Accept json
// @Produce json
// @Success 204
// @Router /logout [post]
func (h *handler) Logout(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
