package relationship

import (
	"context"
	"testing"

	"github.com/darllantissei/genealogy-tree/application/enum/relationship"
	"github.com/darllantissei/genealogy-tree/application/model"
	mockrelationship "github.com/darllantissei/genealogy-tree/mocks/relationship"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRelationshipService_checkRelationship(t *testing.T) {

	const testValidationRelationship = "test_validations_of_relationship"

	relationshipService := RelationshipService{}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testValidationRelationship), "test if relationship is empty")

	msgErr := relationshipService.checkRelationship(ctx, model.Relationship{})

	require.Greater(t, len(msgErr), 0, "There was validations, then messages will return")

	ctx = context.WithValue(ctx, contextKey(testValidationRelationship), "test if person with no link with relationship")

	msgErr = relationshipService.checkRelationship(ctx, model.Relationship{PersonID: "32423klj"})

	require.Greater(t, len(msgErr), 0, "There was validations, then messages will return")

	ctx = context.WithValue(ctx, contextKey(testValidationRelationship), "test if exists only person in relationship")

	msgErr = relationshipService.checkRelationship(ctx, model.Relationship{PersonID: "32423klj", Members: []model.RelationshipMember{{PersonID: "23423"}}})

	require.Greater(t, len(msgErr), 0, "There was validations, then messages will return")
}

func TestRelationshipService_checkSamePerson(t *testing.T) {

	const testCheckSamePerson = "test_check_same_person"

	relationshipService := RelationshipService{}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testCheckSamePerson), "test with link between different persons")

	mockRelationship := model.Relationship{
		PersonID: "234234",
		Members: []model.RelationshipMember{
			{
				PersonID: "645645",
			},
		},
	}

	err := relationshipService.checkSamePerson(ctx, mockRelationship)

	require.Nil(t, err, "The error must be nil")

	ctx = context.WithValue(ctx, contextKey(testCheckSamePerson), "test with link having the same person")

	mockRelationship.Members = append(mockRelationship.Members, model.RelationshipMember{PersonID: "234234"})

	err = relationshipService.checkSamePerson(ctx, mockRelationship)

	require.NotNil(t, err, "The error must be returned")
}

func TestRelationshipService_relationshipUnallowed(t *testing.T) {

	const testRelationshipUnallowed = "test_relationship_unallowed"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	persistenceDB := mockrelationship.NewMockIRelationshipPersistenceDB(ctrl)

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testRelationshipUnallowed), "Unallowed incest between sibling")

	relationshipService := RelationshipService{
		PersistenceDB: persistenceDB,
	}

	mockRelationship := model.Relationship{
		PersonID: "2910f588-9189-4f45-929b-3c3883326e2e",
		Members: []model.RelationshipMember{
			{
				PersonID: "aa0063dc-734b-4215-8d34-a65a598b1dca",
				Type:     relationship.Spouse,
			},
		},
	}

	persistenceDB.EXPECT().Get(ctx, model.Relationship{PersonID: "aa0063dc-734b-4215-8d34-a65a598b1dca"}).
		Return(model.Relationship{PersonID: "aa0063dc-734b-4215-8d34-a65a598b1dca", Members: []model.RelationshipMember{{PersonID: "2910f588-9189-4f45-929b-3c3883326e2e", Type: relationship.Sibling}}}, nil).AnyTimes()

	err := relationshipService.relationshipUnallowed(ctx, mockRelationship)

	require.NotNil(t, err, "Occurred a relationship unallowed, then the error will returned")

	mockRelationship = model.Relationship{
		ID:       "",
		PersonID: "2910f588-9189-4f45-929b-3c3883326e2e",
		Members: []model.RelationshipMember{
			{
				PersonID: "aa0063dc-734b-4215-8d34-a65a598b1dca",
				Type:     relationship.Sibling,
			},
		},
	}

	err = relationshipService.relationshipUnallowed(ctx, mockRelationship)

	require.Nil(t, err, "Don't occurred relationship unallowed")

}
