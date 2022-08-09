package web

import (
	personweb "github.com/darllantissei/genealogy-tree/adapters/web/person"
	relationshipweb "github.com/darllantissei/genealogy-tree/adapters/web/relationship"
	"github.com/labstack/echo"
)

func (ws *WebServer) buildRoutes(echoServer *echo.Echo) {

	groupV1 := echoServer.Group("/v1")

	ws.personRoutes(groupV1)

	ws.relationshipRoutes(groupV1)

}

func (ws *WebServer) personRoutes(route *echo.Group) {

	personWeb := personweb.PersonWeb{
		ApplicationService: ws.ApplicationService,
		CommonWeb:          ws.commonWeb,
	}

	personGroup := route.Group("/person")

	personGroup.POST("", personWeb.RecordHandler)

	personGroup.GET("", personWeb.ListHandler)

	personGroup.GET("/:person_id", personWeb.GetHandler)

	personGroup.PUT("/:person_id", personWeb.UpdateHandler)

}

func (ws *WebServer) relationshipRoutes(route *echo.Group) {

	relationshipWeb := relationshipweb.RelationshipWeb{
		ApplicationService: ws.ApplicationService,
		CommonWeb:          ws.commonWeb,
	}

	relationshipGroup := route.Group("/relationship")

	relationshipGroup.POST("", relationshipWeb.CreateHandler)

	relationshipGroup.GET("/:person_id", relationshipWeb.GetHandler)
}
