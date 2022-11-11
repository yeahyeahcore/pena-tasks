package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type testSayCase struct {
	name                 string
	input                string
	exceptedStatusCode   int
	exceptedResponseBody string
	exceptedStateValue   string
	isSayRequestSended   bool
}

func Test_Say(t *testing.T) {
	testCases := []testSayCase{
		{
			name:                 "OK",
			input:                `{"word": "test1"}`,
			exceptedStatusCode:   http.StatusOK,
			exceptedResponseBody: "success changed state",
			exceptedStateValue:   "test1",
			isSayRequestSended:   true,
		},
		{
			name:                 "Empty Body",
			input:                "",
			exceptedStatusCode:   http.StatusBadRequest,
			exceptedResponseBody: "failed read body",
			exceptedStateValue:   "",
			isSayRequestSended:   true,
		},
		{
			name:               "Request not sended",
			exceptedStateValue: "",
			isSayRequestSended: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			echoInstance := echo.New()
			request := httptest.NewRequest(http.MethodPost, "/say", strings.NewReader(testCase.input))
			recorder := httptest.NewRecorder()

			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			echoContext := echoInstance.NewContext(request, recorder)
			controller := &ListenerController{
				logger: logrus.New(),
				state:  "",
				mState: &sync.RWMutex{},
			}

			if testCase.isSayRequestSended == false {
				assert.Equal(t, testCase.exceptedStateValue, controller.state)
			}

			if testCase.isSayRequestSended == true && assert.NoError(t, controller.Say(echoContext)) {
				assert.Equal(t, testCase.exceptedStatusCode, recorder.Code)
				assert.Equal(t, testCase.exceptedResponseBody, recorder.Body.String())
				assert.Equal(t, testCase.exceptedStateValue, controller.state)
			}
		})
	}
}

func Test_Subscribe(t *testing.T) {
	testCases := []testSayCase{
		{
			name:                 "OK",
			input:                `{"word": "test1"}`,
			exceptedStatusCode:   http.StatusOK,
			exceptedResponseBody: "data: test1\n\n",
			exceptedStateValue:   "test1",
			isSayRequestSended:   true,
		},
		{
			name:                 "Request not sended",
			exceptedStatusCode:   http.StatusOK,
			exceptedResponseBody: "data: \n\n",
			exceptedStateValue:   "",
			isSayRequestSended:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			echoInstance := echo.New()
			requestListen := httptest.NewRequest(http.MethodGet, "/listen", nil)
			requestSay := httptest.NewRequest(http.MethodPost, "/say", strings.NewReader(testCase.input))
			recorderListen := httptest.NewRecorder()
			recorderSay := httptest.NewRecorder()

			requestListen.Header.Set(echo.HeaderContentType, "no-cache")
			requestListen.Header.Set(echo.HeaderAccept, "text/event-stream")

			requestSay.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			controller := &ListenerController{
				logger: logrus.New(),
				state:  "",
				mState: &sync.RWMutex{},
			}

			echoContextListen := echoInstance.NewContext(requestListen, recorderListen)
			echoContextSay := echoInstance.NewContext(requestSay, recorderSay)

			if testCase.isSayRequestSended == true {
				assert.NoError(t, controller.Say(echoContextSay))
			}

			go controller.Subscribe(echoContextListen)

			time.Sleep(time.Second * 5)

			assert.Equal(t, testCase.exceptedStateValue, controller.state)
			assert.Equal(t, testCase.exceptedStatusCode, recorderListen.Code)
			assert.Contains(t, recorderListen.Body.String(), testCase.exceptedResponseBody)
		})
	}
}
