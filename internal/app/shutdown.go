package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/yeahyeahcore/pena-tasks/internal/server"
)

type gracefulShutdownDeps struct {
	httpServer *server.HTTP
	cancel     context.CancelFunc
}

func gracefulShutdown(ctx context.Context, deps *gracefulShutdownDeps) {
	defer deps.httpServer.Stop(ctx)
	defer deps.cancel()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	deps.httpServer.Logger.Infoln("Shutting down server ...")
}
