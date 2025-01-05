package task

import (
	"errors"
	"log/slog"
	"net/http"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/service/interfaces"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	incorrectTaskID = "Incorrect task id"
	internalError   = "Internal server error"
	taskNotFound    = "Task not found"
	bindError       = "Bind error"
)

type handler struct {
	taskService interfaces.TaskService
}

func NewHandler(s interfaces.TaskService) *handler {
	return &handler{
		taskService: s,
	}
}

// GetTaskByID
// @Tags Task
// @Summary Get task by ID
// @Description Get task by ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} model.Task
// @Failure 422 {object} model.Message "Incorrect id"
// @Failure 404 {object} model.Message "Task not found"
// @Failure 500 {object} model.Message "Internal error"
// @Security Login
// @Router /task/{id} [get]
func (h *handler) GetTaskByID(c echo.Context) error {
	taskID := c.Param("task_id")
	if err := uuid.Validate(taskID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectTaskID})
	}
	task, err := h.taskService.GetTaskByID(c.Request().Context(), taskID)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: taskNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalError})
	}
	return c.JSON(http.StatusOK, task)
}

// CreateTask
// @Tags Task
// @Summary Create task
// @Description Create task
// @Accept json
// @Produce json
// @Param task body model.TaskInserted true "Task"
// @Success 201 {object} model.Task
// @Failure 422 {object} model.Message "Incorrect request body (bind error)"
// @Failure 500 {object} model.Message "Internal error"
// @Security Login
// @Router /task [post]
func (h *handler) CreateTask(c echo.Context) error {
	taskInsert := &model.TaskInserted{}
	err := c.Bind(taskInsert)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: bindError})
	}
	task, err := h.taskService.CreateTask(c.Request().Context(), taskInsert)
	if err != nil {
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalError})
	}
	return c.JSON(http.StatusOK, task)
}

// UpdateTask
// @Tags Task
// @Summary Update task
// @Description Update task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body model.TaskInserted true "Task"
// @Success 200 {object} model.Task
// @Failure 404 {object} model.Message "Task not found"
// @Failure 422 {object} model.Message "Incorrect id/Incorrect request body (bind error)"
// @Failure 500 {object} model.Message "Internal error"
// @Security Login
// @Router /task/{id} [put]
func (h *handler) UpdateTask(c echo.Context) error {
	taskID := c.Param("task_id")
	if err := uuid.Validate(taskID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectTaskID})
	}
	taskUpdate := &model.TaskInserted{}
	err := c.Bind(taskUpdate)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: bindError})
	}
	task, err := h.taskService.UpdateTask(c.Request().Context(), taskID, taskUpdate)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: taskNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalError})
	}
	return c.JSON(http.StatusOK, task)
}

// DeleteTask
// @Tags Task
// @Summary Delete task
// @Description Delete task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404 {object} model.Message "Task not found"
// @Failure 422 {object} model.Message "Incorrect id"
// @Failure 500 {object} model.Message "Internal error"
// @Security Login
// @Router /task/{id} [delete]
func (h *handler) DeleteTask(c echo.Context) error {
	taskID := c.Param("task_id")
	if err := uuid.Validate(taskID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: incorrectTaskID})
	}
	err := h.taskService.DeleteTask(c.Request().Context(), taskID)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: taskNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalError})
	}
	return c.NoContent(http.StatusNoContent)
}
