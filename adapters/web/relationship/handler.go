package relationshipweb

import (
	"context"
	"net/http"
	"strings"

	commonweb "github.com/darllantissei/genealogy-tree/adapters/web/common"
	"github.com/darllantissei/genealogy-tree/application"
	"github.com/darllantissei/genealogy-tree/application/model"
	"github.com/labstack/echo"
)

type RelationshipWeb struct {
	ApplicationService application.Application
	CommonWeb          commonweb.Common
}

func (rw *RelationshipWeb) CreateHandler(ectx echo.Context) error {

	req := ectx.Request()

	if req == nil {

		ctx := context.Background()

		respError := rw.ApplicationService.CommonService.BuildError(ctx, []string{"the request is nil, impossible create a relationship"})

		return rw.CommonWeb.ParseResponse(ectx, http.StatusInternalServerError, respError)

	}

	ctx := req.Context()

	relationshipDTO := &RelationshipDTO{}

	err := ectx.Bind(relationshipDTO)

	if err != nil {
		respError := rw.ApplicationService.CommonService.BuildError(ctx, []string{"the struct to record a relationship is invalid"})
		return rw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, respError)
	}

	if relationshipDTO.IsEmpty() {
		respError := rw.ApplicationService.CommonService.BuildError(ctx, []string{"the data of relationship is empty, impossible to create a relationship"})
		return rw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, respError)
	}

	relationshipToRecord := rw.parseDataToApp(*relationshipDTO)

	relationshipCreated, err := rw.ApplicationService.RelationshipService.Create(ctx, relationshipToRecord)

	if err != nil {
		return rw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, err)
	}

	return rw.CommonWeb.ParseResponse(ectx, http.StatusCreated, rw.parseData(relationshipCreated))
}

func (pw *RelationshipWeb) GetHandler(ectx echo.Context) error {

	personID := ectx.Param("person_id")

	if strings.EqualFold(personID, "") {
		respError := pw.ApplicationService.CommonService.BuildError(ectx.Request().Context(), []string{"the ID of person is empty, impossible to fetch relationship"})
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, respError)
	}

	ctx := ectx.Request().Context()

	relationshipFetched, err := pw.ApplicationService.RelationshipService.Fetch(ctx, model.Relationship{PersonID: personID})

	if err != nil {
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, err)
	}

	if relationshipFetched.IsEmpty() {
		respErrNotFound := pw.ApplicationService.CommonService.BuildError(ectx.Request().Context(), []string{"The relationship of person not found"})
		return pw.CommonWeb.ParseResponse(ectx, http.StatusNotFound, respErrNotFound)
	}

	return pw.CommonWeb.ParseResponse(ectx, http.StatusOK, pw.parseData(relationshipFetched))
}
