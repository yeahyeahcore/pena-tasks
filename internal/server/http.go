package server

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/pena-tasks/internal/initialize"
)

type HTTP struct {
	Logger *logrus.Entry
	server *http.Server
	echo   *echo.Echo
}

func New(logger *logrus.Logger) *HTTP {
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())
	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	return &HTTP{
		echo:   echo,
		Logger: logrus.NewEntry(logger),
		server: &http.Server{
			Handler:        echo,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (receiver *HTTP) Listen(address string) error {
	receiver.server.Addr = address

	return receiver.server.ListenAndServe()
}

func (receiver *HTTP) Stop(ctx context.Context) {
	receiver.server.Shutdown(ctx)
}

func (receiver *HTTP) Register(controllers *initialize.Controllers) *HTTP {
	receiver.echo.GET("/listen", controllers.ListenerController.Subscribe)
	receiver.echo.POST("/say", controllers.ListenerController.Say)

	return receiver
}
