package dbrelationship

import (
	"context"

	"github.com/darllantissei/genealogy-tree/application/model"
)

type RelationshipDB struct {
}

func NewRelationshipDB() *RelationshipDB {
	return &RelationshipDB{}

}

func (db *RelationshipDB) Record(ctx context.Context, rtshp model.Relationship) (model.Relationship, error) {
	panic("Database of relationship not implemented")
}
func (db *RelationshipDB) Get(ctx context.Context, rtshp model.Relationship) (model.Relationship, error) {
	panic("Database of relationship not implemented")
}
