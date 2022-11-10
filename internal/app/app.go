package app

import (
	"context"

	"github.com/yeahyeahcore/pena-tasks/internal/initialize"
	"github.com/yeahyeahcore/pena-tasks/internal/models"
	"github.com/yeahyeahcore/pena-tasks/internal/server"

	"github.com/sirupsen/logrus"
)

func Run(config *models.Config, logger *logrus.Logger) {
	ctx, cancel := context.WithCancel(context.Background())
	controllers := initialize.NewControllers(&initialize.ControllersDeps{Logger: logger})
	httpServer := server.New(logger).Register(controllers)

	// TODO: обработать панику при ошибке запуска сервера
	go runHTTP(&RunHTTPDeps{
		httpServer: httpServer,
		config:     &config.HTTP,
	})

	gracefulShutdown(ctx, &gracefulShutdownDeps{
		httpServer: httpServer,
		cancel:     cancel,
	})
}
