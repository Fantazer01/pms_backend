package app

import (
	"context"
	"log/slog"
	"pms_backend/pms_api/internal/config"

	"github.com/labstack/echo/v4"
)

type App struct {
	// db *pgxpool.Pool
	config *config.Config
	router *echo.Echo
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.init(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) Run() error {
	go func() {
		err := a.runHttpServer()
		if err != nil {
			slog.Error("run server error", "err", err)
		}
	}()
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	//a.db.Close()
	return a.stopHttpServer(ctx)
}
