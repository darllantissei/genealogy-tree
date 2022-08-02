package relationship

import (
	"context"
	"testing"

	"github.com/darllantissei/genealogy-tree/application/model"
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

	msgErr = relationshipService.checkRelationship(ctx, model.Relationship{PersonID: "32423klj", Links: []model.RelationshipLink{{PersonID: "23423"}}})

	require.Greater(t, len(msgErr), 0, "There was validations, then messages will return")
}

func TestRelationshipService_checkSamePerson(t *testing.T) {

	const testCheckSamePerson = "test_check_same_person"

	relationshipService := RelationshipService{}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testCheckSamePerson), "test with link between different persons")

	mockRelationship := model.Relationship{
		PersonID: "234234",
		Links: []model.RelationshipLink{
			{
				PersonID: "645645",
			},
		},
	}

	err := relationshipService.checkSamePerson(ctx, mockRelationship)

	require.Nil(t, err, "The error must be nil")

	ctx = context.WithValue(ctx, contextKey(testCheckSamePerson), "test with link having the same person")

	mockRelationship.Links = append(mockRelationship.Links, model.RelationshipLink{PersonID: "234234"})

	err = relationshipService.checkSamePerson(ctx, mockRelationship)

	require.NotNil(t, err, "The error must be returned")
}
