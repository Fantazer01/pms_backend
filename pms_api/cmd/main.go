package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	_ "pms_backend/pms_api/docs"
	"pms_backend/pms_api/internal/app"
	"syscall"
	"time"
)

// @title           PMS API
// @version         1.0
// @description     PMS - project management system. It is course work of student team for 5 course in MEPhI.
// @termsOfService  http://swagger.io/terms/

// @host localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey Login
// @in header
// @name Authorization
// @description Type "Bearer TOKEN" to correctly set the API Key

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	server, err := app.NewApp(ctx)
	if err != nil {
		slog.Error("failed to init app: " + err.Error())
		return
	}
	err = server.Run()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to run app: %s", err))
	}
	<-ctx.Done()
	ctx, stop = context.WithTimeout(context.Background(), 5*time.Second)
	defer stop()
	err = server.Stop(ctx)
	if err != nil {
		slog.Error("stop server error", "err", err)
	}
}
