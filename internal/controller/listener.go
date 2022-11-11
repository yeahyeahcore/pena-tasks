package controller

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/pena-tasks/pkg/json"
)

type SetStateRequest struct {
	Word string `json:"word"`
}

type ListenerController struct {
	logger *logrus.Logger
	mState *sync.RWMutex
	state  string
}

func NewListenerController(logger *logrus.Logger) *ListenerController {
	return &ListenerController{
		logger: logger,
		state:  "test",
		mState: &sync.RWMutex{},
	}
}

func (receiver *ListenerController) Say(ctx echo.Context) error {
	request, err := json.Read[SetStateRequest](ctx.Request().Body)
	if err != nil {
		receiver.logger.Errorln("read /say body error: ", err.Error())

		return ctx.String(http.StatusBadRequest, "failed read body")
	}

	receiver.mState.Lock()

	receiver.state = request.Word

	receiver.mState.Unlock()

	return ctx.String(http.StatusOK, "success changed state")
}

func (receiver *ListenerController) Subscribe(ctx echo.Context) error {
	writer := ctx.Response().Writer

	writer.Header().Set(echo.HeaderContentType, "text/event-stream")
	writer.Header().Set(echo.HeaderCacheControl, "no-cache")
	writer.Header().Set(echo.HeaderConnection, "keep-alive")

	ticker := time.NewTicker(time.Second)

	flusher, ok := writer.(http.Flusher)
	if !ok {
		ctx.String(http.StatusInternalServerError, "flusher not init")
	}

	for {
		select {
		case <-ctx.Request().Context().Done():
			return nil
		case <-ticker.C:
			receiver.mState.RLock()
			fmt.Fprintf(writer, "data: %s\n\n", receiver.state)
			receiver.mState.RUnlock()
			flusher.Flush()
		}
	}
}
