package app

import (
	"context"
	"log/slog"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "2M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
)

func (a *App) runHttpServer() error {
	slog.Info("starting http server", "address", a.config.GetAddress())
	a.router.Server.ReadTimeout = readTimeout
	a.router.Server.WriteTimeout = writeTimeout
	a.router.Server.MaxHeaderBytes = maxHeaderBytes

	return a.router.Start(a.config.GetAddress())
}

func (a *App) stopHttpServer(ctx context.Context) error {
	return a.router.Shutdown(ctx)
}
