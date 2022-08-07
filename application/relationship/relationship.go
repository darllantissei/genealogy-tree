package relationship

import (
	"context"

	"github.com/darllantissei/genealogy-tree/application/model"
)

func (r *RelationshipService) getKinship(ctx context.Context, rtshp model.Relationship) (model.Relationship, error) {
	// TODO: implementar o parentesco do relacionamento
	return model.Relationship{PersonID: "3211b2b2-3a5c-4ac0-9f80-f328b5f1c3c4"}, nil
}
