package relationship

import (
	"context"

	"github.com/darllantissei/genealogy-tree/application/model"
)

type IRelationshipService interface {
	Create(ctx context.Context, rtshp model.Relationship) (model.Relationship, error)
	Fetch(ctx context.Context, rtshp model.Relationship) (model.Relationship, error)
}

type IRelationshipWriteDB interface {
	Record(ctx context.Context, rtshp model.Relationship) (model.Relationship, error)
}

type IRelationshipReadDB interface {
	Get(ctx context.Context, rtshp model.Relationship) (model.Relationship, error)
	GetKinship(ctx context.Context, member model.RelationshipMember) ([]model.RelationshipMember, error)
}

type IRelationshipPersistenceDB interface {
	IRelationshipWriteDB
	IRelationshipReadDB
}
