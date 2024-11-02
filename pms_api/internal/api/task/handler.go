package task

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

// GetTaskByID
// @Tags Task
// @Summary Get task by ID
// @Description Get task by ID
// @Accept json
// @Produce json
// @Success 200 {object} model.Task
// @Router /task/{id} [get]
func (h *handler) GetTaskByID(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// CreateTask
// @Tags Task
// @Summary Create task
// @Description Create task
// @Accept json
// @Produce json
// @Success 201 {object} model.Task
// @Router /task [post]
func (h *handler) CreateTask(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// UpdateTask
// @Tags Task
// @Summary Update task
// @Description Update task
// @Accept json
// @Produce json
// @Success 200 {object} model.Task
// @Router /task/{id} [put]
func (h *handler) UpdateTask(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

// DeleteTask
// @Tags Task
// @Summary Delete task
// @Description Delete task
// @Accept json
// @Produce json
// @Success 204
// @Router /task/{id} [delete]
func (h *handler) DeleteTask(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
