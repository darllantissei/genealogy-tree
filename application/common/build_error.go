package common

import (
	"context"

	statusapplication "github.com/darllantissei/genealogy-tree/application/enum/status_application"
	"github.com/darllantissei/genealogy-tree/application/model"
)

func BuildError(ctx context.Context, msgErr []string) model.Returns {

	var ret model.Returns

	ret.Status = statusapplication.Error
	ret.Message = msgErr

	return ret
}
