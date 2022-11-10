package initialize

import (
	"github.com/yeahyeahcore/pena-tasks/internal/controller"

	"github.com/sirupsen/logrus"
)

type ControllersDeps struct {
	Logger *logrus.Logger
}

type Controllers struct {
	ListenerController *controller.ListenerController
}

func NewControllers(deps *ControllersDeps) *Controllers {
	return &Controllers{
		ListenerController: controller.NewListenerController(deps.Logger),
	}
}
