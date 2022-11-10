package main

import (
	"github.com/yeahyeahcore/pena-tasks/internal/app"
	"github.com/yeahyeahcore/pena-tasks/internal/models"
	"github.com/yeahyeahcore/pena-tasks/pkg/env"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	config, err := env.Read[models.Config]()
	if err != nil {
		logger.Fatalf("config read is failed: %s", err)
	}

	app.Run(config, logger)
}
