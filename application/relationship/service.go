package relationship

import (
	"context"

	"github.com/darllantissei/genealogy-tree/application/common"
	"github.com/darllantissei/genealogy-tree/application/model"
	"github.com/darllantissei/genealogy-tree/application/person"
)

type RelationshipService struct {
	PersistenceDB IRelationshipPersistenceDB
	CommonService common.ICommonService
	PersonService person.IPersonService
}

func (r *RelationshipService) Create(ctx context.Context, rtshp model.Relationship) (model.Relationship, error) {

	err := r.checkDependencyInjection(ctx)

	if err != nil {
		return model.Relationship{}, r.CommonService.BuildError(ctx, []string{err.Error()})
	}

	msgErr := r.checkRelationship(ctx, rtshp)

	err = r.checkRelationshipExists(ctx, rtshp)

	if err != nil {
		msgErr = append(msgErr, err.Error())
	}

	if len(msgErr) > 0 {
		return model.Relationship{}, r.CommonService.BuildError(ctx, msgErr)
	}

	rtshp, err = r.PersistenceDB.Record(ctx, rtshp)

	if err != nil {
		return model.Relationship{}, r.CommonService.BuildError(ctx, []string{err.Error()})
	}

	return rtshp, nil

}
