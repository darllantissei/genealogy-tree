package relationship

import (
	"context"

	"github.com/darllantissei/genealogy-tree/application/common"
	"github.com/darllantissei/genealogy-tree/application/model"
)

type RelationshipService struct {
	PersistenceDB IRelationshipPersistenceDB
	CommonService common.ICommonService
}

func (r *RelationshipService) Create(ctx context.Context, rtshp model.Relationship) (model.Relationship, error) {

	err := r.checkDependencyInjection(ctx)

	if err != nil {
		return model.Relationship{}, r.CommonService.BuildError(ctx, []string{err.Error()})
	}

	msgErr := r.checkRelationship(ctx, rtshp)

	if len(msgErr) > 0 {
		return model.Relationship{}, r.CommonService.BuildError(ctx, msgErr)
	}

	panic("not implemented")
}
