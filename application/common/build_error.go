package common

import (
	"context"

	statusapplication "github.com/darllantissei/genealogy-tree/application/enum/status_application"
	"github.com/darllantissei/genealogy-tree/application/model"
)

type CommonService struct{}

func (c CommonService) BuildError(ctx context.Context, msgErr []string) model.Returns {

	var ret model.Returns

	ret.Return = model.Return{
		Status:  statusapplication.Error,
		Message: msgErr,
	}

	return ret
}
