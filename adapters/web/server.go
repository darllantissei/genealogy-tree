package web

import (
	commonweb "github.com/darllantissei/genealogy-tree/adapters/web/common"
	"github.com/darllantissei/genealogy-tree/application"
)

type WebServer struct {
	ApplicationService application.Application
	commonWeb          commonweb.Common
}

func MakeNewWebServer(applicationService application.Application) *WebServer {
	return &WebServer{
		ApplicationService: applicationService,
		commonWeb: commonweb.Common{
			ApplicationService: applicationService,
		},
	}
}

func (ws *WebServer) Serve(port int, debugEnabled bool) {

	ws.Make(debugEnabled)

	ws.ProvideEchoInstance(ws.buildRoutes)

	ws.Run(port)
}
