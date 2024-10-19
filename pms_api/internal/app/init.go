package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"pms_backend/pms_api/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *App) init(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLog,
		a.initDb,
		a.initUseCase,
		a.initRouter,
		a.initSwagger,
		a.initMiddleware,
		a.registerRoutes,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			slog.Error("init deps", "error message", err.Error())
			return err
		}
	}
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	a.config = &config.Config{
		Port: ":8080",
	}
	return nil
}

func (a *App) initLog(ctx context.Context) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	slog.Info("config", "val", a.config)
	return nil
}

func (a *App) initDb(ctx context.Context) error {
	return nil
}

func (a *App) initUseCase(ctx context.Context) error {
	return nil
}

func (a *App) initRouter(_ context.Context) error {
	a.router = echo.New()
	a.router.Logger.SetLevel(log.INFO)
	a.router.Pre(middleware.RemoveTrailingSlash())
	a.router.HideBanner = true
	return nil
}

func (a *App) initSwagger(ctx context.Context) error {
	a.router.GET("/swagger/*", echoSwagger.WrapHandler)

	a.router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	return nil
}

func (a *App) initMiddleware(ctx context.Context) error {
	return nil
}

func (a *App) registerRoutes(ctx context.Context) error {
	a.router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	return nil
}
