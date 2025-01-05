package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	project_handler "pms_backend/pms_api/internal/api/http/project"
	task_handler "pms_backend/pms_api/internal/api/http/task"
	user_handler "pms_backend/pms_api/internal/api/http/user"
	"pms_backend/pms_api/internal/config"
	"pms_backend/pms_api/internal/pkg/model"
	project_repository "pms_backend/pms_api/internal/repository/project/postgres"
	task_repository "pms_backend/pms_api/internal/repository/task/postgres"
	user_repository "pms_backend/pms_api/internal/repository/user/postgres"
	project_service "pms_backend/pms_api/internal/service/project"
	task_service "pms_backend/pms_api/internal/service/task"
	user_service "pms_backend/pms_api/internal/service/user"

	"github.com/jackc/pgx/v5/pgxpool"
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
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	a.config = config
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
	db, err := pgxpool.New(ctx, a.config.Database.ConnectionString)
	if err != nil {
		return fmt.Errorf("initialize db connection: %w", err)
	}
	a.db = db
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
	a.router.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		c.Logger().Error(err)
		err = c.JSON(code, model.Message{Message: "Internal server error"})
		if err != nil {
			c.Logger().Error(err)
		}
	}

	return nil
}

func (a *App) registerRoutes(ctx context.Context) error {
	api := a.router.Group(a.config.Http.BasePath)

	handlers := []Handler{
		project_handler.NewHandler(
			project_service.NewProjectService(project_repository.NewRepository(a.db)),
		),

		user_handler.NewHandler(
			user_service.NewUserService(user_repository.NewUserRepository(a.db)),
		),

		task_handler.NewHandler(
			task_service.NewTaskService(task_repository.NewRepository(a.db)),
		),
	}

	for _, h := range handlers {
		h.RegisterRoutes(api)
	}

	return nil
}
