package common

import (
	"context"

	"github.com/darllantissei/genealogy-tree/application/model"
)

type ICommonService interface {
	BuildError(ctx context.Context, msgErr []string) model.Returns
	SliceExists(slice interface{}, item interface{}) bool
}
