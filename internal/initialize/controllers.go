package initialize

import (
	"github.com/sirupsen/logrus"
)

type ControllersDeps struct {
	Logger *logrus.Logger
}

type Controllers struct {
}

func NewControllers(deps *ControllersDeps) *Controllers {

	return &Controllers{}
}
