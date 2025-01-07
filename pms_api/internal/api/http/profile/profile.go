package profile

import (
	"errors"
	"log/slog"
	"net/http"
	"pms_backend/pms_api/internal/pkg/apperror"
	"pms_backend/pms_api/internal/pkg/model"
	"pms_backend/pms_api/internal/pkg/service/interfaces"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	userNotFound        = "User not found"
	internalServerError = "Internal server error"
)

type handler struct {
	profileService interfaces.ProfileService
}

func NewHandler(s interfaces.ProfileService) *handler {
	return &handler{
		profileService: s,
	}
}

// GetProfile
// @Tags Profile
// @Summary Send information about authorized user
// @Description Send information about authorized user
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Failure 500 {object} model.Message
// @Security Login
// @Router /profile [get]
func (h *handler) GetProfile(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.AppClaims)

	user, err := h.profileService.GetUserByID(c.Request().Context(), claims.Subject)
	if err != nil {
		if errors.Is(err, apperror.NotFound) {
			return c.JSON(http.StatusNotFound, model.Message{Message: userNotFound})
		}
		slog.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Message{Message: internalServerError})
	}
	return c.JSON(http.StatusOK, user)
}
