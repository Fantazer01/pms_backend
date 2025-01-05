package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/service/interfaces"

	"github.com/labstack/echo/v4"
)

const (
	bindError           = "Bind error"
	unauthorized        = "Unauthorized"
	failedGenerateToken = "Failed to generate token"
)

type handler struct {
	authService interfaces.AuthService
}

func NewHandler(s interfaces.AuthService) *handler {
	return &handler{
		authService: s,
	}
}

// Login
// @Tags Auth
// @Summary Login
// @Description Login
// @Accept json
// @Produce json
// @Param form body model.AuthForm true "Authentication form"
// @Success 200 {object} model.Tokens
// @Failure 401 {object} model.Message "Unauthorized"
// @Failure 422 {object} model.Message "Incorrect request body (bind error)"
// @Router /login [post]
func (h *handler) Login(c echo.Context) error {
	var user *model.AuthForm
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Message{Message: bindError})
	}

	signedToken, err := h.authService.Authentication(c.Request().Context(), user.Login, user.Password)
	if err != nil {
		slog.Error(err.Error())
		if errors.Is(err, apperror.Unauthorized) {
			return c.JSON(http.StatusUnauthorized, model.Message{Message: unauthorized})
		}
		return c.JSON(http.StatusInternalServerError, model.Message{Message: failedGenerateToken})
	}
	return c.JSON(http.StatusOK, model.Tokens{AccessToken: signedToken})
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
