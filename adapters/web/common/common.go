package commonweb

import (
	"github.com/darllantissei/genealogy-tree/application"
	"github.com/labstack/echo"
)

type Common struct {
	ApplicationService application.Application
}

func (c Common) ParseResponse(ectx echo.Context, httpStatusCode int, response interface{}) error {
	request := ectx.Request()

	if request == nil {
		panic("the request is nil when parse respose")
	}

	contentType, ok := request.Header[echo.HeaderContentType]

	if !ok {
		return ectx.JSON(httpStatusCode, response)
	}

	contentTypeJson := []string{
		echo.MIMEApplicationJSON,
		echo.MIMEApplicationJSONCharsetUTF8,
	}

	contentTypeXML := []string{
		echo.MIMEApplicationXML,
		echo.MIMEApplicationXMLCharsetUTF8,
		echo.MIMETextXML,
		echo.MIMETextXMLCharsetUTF8,
	}

	for _, typeContent := range contentType {

		switch {
		case c.ApplicationService.CommonService.SliceExists(contentTypeJson, typeContent):
			return ectx.JSON(httpStatusCode, response)
		case c.ApplicationService.CommonService.SliceExists(contentTypeXML, typeContent):
			return ectx.XML(httpStatusCode, response)
		}
	}

	return ectx.JSON(httpStatusCode, response)
}
