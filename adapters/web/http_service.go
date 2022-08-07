package web

import (
	"fmt"
	"net/http"

	statusapplication "github.com/darllantissei/genealogy-tree/application/enum/status_application"
	"github.com/darllantissei/genealogy-tree/application/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

var echoServer *echo.Echo

func (ws *WebServer) Make(debugEnabled bool) {

	if echoServer == nil {
		echoServer = echo.New()
	}

	// Esconde o cabeçalho do Echo
	echoServer.HideBanner = true

	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.Recover())

	if debugEnabled {
		echoServer.Debug = true
		echoServer.Use(middleware.Logger())
	}

	echoServer.HTTPErrorHandler = ws.errorServer
}

func (ws *WebServer) ProvideEchoInstance(task func(e *echo.Echo)) {
	task(echoServer)
}

func (ws *WebServer) errorServer(err error, ectx echo.Context) {
	var (
		code      = http.StatusInternalServerError
		errResult model.Returns
	)

	errResult.Return.Status = statusapplication.Error

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		errResult.Return.Message = append(errResult.Return.Message, fmt.Sprintf("%#v", he.Message))
		if he.Internal != nil {
			errResult.Return.Message = append(errResult.Return.Message, fmt.Sprintf("%v, %v", err, he.Internal))
		}
	} else {
		errResult.Return.Message = append(errResult.Return.Message, http.StatusText(code))
	}

	if !ectx.Response().Committed {
		if ectx.Request().Method == echo.HEAD {
			err = ectx.NoContent(code)
		} else {
			err = ws.commonWeb.ParseResponse(ectx, code, errResult)
		}
		if err != nil {
			ectx.Echo().Logger.Error(err)
		}
	} else {
		log.WithFields(log.Fields{
			"error": err.Error(),
		},
		).Error("err: error in handle erro web")
	}
}

func (ws *WebServer) Run(port int) {
	echoServer.Logger.Fatal(echoServer.Start(fmt.Sprintf(":%d", port)))
}
