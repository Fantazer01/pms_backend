package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

// GetUsers
// @Tags User
// @Summary Get users
// @Description Get users
// @Accept json
// @Produce json
// @Param pageIndex query int false "Page index"
// @Param pageSize query int false "Page size"
// @Success 200 {object} model.UsersPaged
// @Router /users [get]
func (h *handler) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// GetUserByID
// @Tags User
// @Summary Get user by ID
// @Description Get user by ID
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func (h *handler) GetUserByID(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// CreateUser
// @Tags User
// @Summary Create user
// @Description Create user
// @Accept json
// @Produce json
// @Success 201 {object} model.User
// @Router /users [post]
func (h *handler) CreateUser(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// UpdateUser
// @Tags User
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Failure 400
// @Router /users/{id} [put]
func (h *handler) UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// DeleteUser
// @Tags User
// @Summary Delete user
// @Description Delete User
// @Accept  json
// @Produce  json
// @Success 204
// @Failure 400
// @Router /users/{id} [delete]
func (h *handler) DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
