package app

import (
	"fmt"

	"github.com/yeahyeahcore/pena-tasks/internal/models"
	"github.com/yeahyeahcore/pena-tasks/internal/server"
)

// RunHTTPDeps - dependencies for running HTTP server
type RunHTTPDeps struct {
	httpServer *server.HTTP
	config     *models.HTTPConfiguration
}

func runHTTP(deps *RunHTTPDeps) {
	connectionString := fmt.Sprintf("%s:%s", deps.config.Host, deps.config.Port)
	startServerMessage := fmt.Sprintf("Starting HTTP Server on %s", connectionString)

	deps.httpServer.Logger.Infoln(startServerMessage)

	if err := deps.httpServer.Listen(connectionString); err != nil {
		deps.httpServer.Logger.Infoln("HTTP Listen error: ", err)
		panic(err)
	}
}
