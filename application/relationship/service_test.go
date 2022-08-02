package relationship

import (
	"context"
	"testing"

	"github.com/darllantissei/genealogy-tree/application/model"
	mockrelationship "github.com/darllantissei/genealogy-tree/mocks/relationship"
	"github.com/golang/mock/gomock"
)

func TestRelationshipService_Create(t *testing.T) {

	const testCreateRelationship = "test_to_create_relationship"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	persistenceDB := mockrelationship.NewMockIRelationshipPersistenceDB(ctrl)

	relationshipService := RelationshipService{
		PersistenceDB: persistenceDB,
	}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testCreateRelationship), "Don't allow create relationship with only person")

	mockRelationship := model.Relationship{
		PersonID: "84763d21-497d-41f3-8381-727724914576",
		// Type:     relationship.Child,
	}

	relationship, err := relationshipService.Create(ctx, mockRelationship)

	_, _ = relationship, err

}
