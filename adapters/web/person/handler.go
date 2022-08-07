package personweb

import (
	"context"
	"net/http"
	"strings"

	commonweb "github.com/darllantissei/genealogy-tree/adapters/web/common"
	"github.com/darllantissei/genealogy-tree/application"
	"github.com/darllantissei/genealogy-tree/application/model"
	"github.com/labstack/echo"
)

type PersonWeb struct {
	ApplicationService application.Application
	CommonWeb          commonweb.Common
}

func (pw *PersonWeb) RecordHandler(ectx echo.Context) error {

	req := ectx.Request()

	if req == nil {

		ctx := context.Background()

		respError := pw.ApplicationService.CommonService.BuildError(ctx, []string{"the request is nil, impossible record a person"})

		return pw.CommonWeb.ParseResponse(ectx, http.StatusInternalServerError, respError)

	}

	personID := ectx.Param("person_id")

	ctx := req.Context()

	personDTO := &PersonDTO{}

	err := ectx.Bind(personDTO)

	if err != nil {
		detail := ""
		switch echoErr := err.(type) {
		case *echo.HTTPError:
			detail = echoErr.Internal.Error()
		}
		respError := pw.ApplicationService.CommonService.BuildError(ctx, []string{"the struct to record a person is invalid", detail})
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, respError)
	}

	if personDTO.IsEmpty() {
		respError := pw.ApplicationService.CommonService.BuildError(ctx, []string{"the data of person is empty, impossible to create a person"})
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, respError)
	}

	personToRecord := pw.parseDataToApp(*personDTO)

	if strings.EqualFold(personID, "") {

		personRecorded, err := pw.ApplicationService.PersonService.Record(ctx, personToRecord)

		if err != nil {
			return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, err)
		}

		return pw.CommonWeb.ParseResponse(ectx, http.StatusCreated, pw.parseData(personRecorded))
	}

	personToRecord.ID = personID
	personUpdated, err := pw.ApplicationService.PersonService.Record(ctx, personToRecord)

	if err != nil {
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, err)
	}

	return pw.CommonWeb.ParseResponse(ectx, http.StatusOK, pw.parseData(personUpdated))
}

func (pw *PersonWeb) UpdateHandler(ectx echo.Context) error {

	personID := ectx.Param("person_id")

	if strings.EqualFold(personID, "") {
		respError := pw.ApplicationService.CommonService.BuildError(ectx.Request().Context(), []string{"the ID of person is empty, impossible to update"})
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, respError)
	}

	return pw.RecordHandler(ectx)

}

func (pw *PersonWeb) GetHandler(ectx echo.Context) error {

	personID := ectx.Param("person_id")

	if strings.EqualFold(personID, "") {
		respError := pw.ApplicationService.CommonService.BuildError(ectx.Request().Context(), []string{"the ID of person is empty, impossible to fetch"})
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, respError)
	}

	ctx := ectx.Request().Context()

	personFetched, err := pw.ApplicationService.PersonService.Fetch(ctx, model.Person{ID: personID})

	if err != nil {
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, err)
	}

	if personFetched.IsEmpty() {
		respErrNotFound := pw.ApplicationService.CommonService.BuildError(ectx.Request().Context(), []string{"Person not found"})
		return pw.CommonWeb.ParseResponse(ectx, http.StatusNotFound, respErrNotFound)
	}

	return pw.CommonWeb.ParseResponse(ectx, http.StatusOK, pw.parseData(personFetched))
}

func (pw *PersonWeb) ListHandler(ectx echo.Context) error {

	ctx := ectx.Request().Context()

	allPerson, err := pw.ApplicationService.PersonService.All(ctx)

	if err != nil {
		return pw.CommonWeb.ParseResponse(ectx, http.StatusBadRequest, err)
	}

	if len(allPerson) <= 0 {
		respErrNotFound := pw.ApplicationService.CommonService.BuildError(ectx.Request().Context(), []string{"No person"})
		return pw.CommonWeb.ParseResponse(ectx, http.StatusNoContent, respErrNotFound)
	}

	allPersonDTO := []PersonDTO{}

	for _, personAPP := range allPerson {
		allPersonDTO = append(allPersonDTO, pw.parseData(personAPP))
	}

	return pw.CommonWeb.ParseResponse(ectx, http.StatusOK, allPersonDTO)
}
