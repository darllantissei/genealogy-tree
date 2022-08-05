package relationship

import (
	"context"
	"testing"

	"github.com/darllantissei/genealogy-tree/application/enum/kinship"
	"github.com/darllantissei/genealogy-tree/application/enum/relationship"
	"github.com/darllantissei/genealogy-tree/application/model"
	mock_relationship "github.com/darllantissei/genealogy-tree/mocks/relationship"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRelationshipService_getKinship(t *testing.T) {

	const testGetKinship = "test_get_kinship"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	persistenceDB := mock_relationship.NewMockIRelationshipPersistenceDB(ctrl)

	relationshipService := RelationshipService{
		PersistenceDB: persistenceDB,
	}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testGetKinship), "Get relationship with description kinship")

	mockToGetRelationship := model.Relationship{
		PersonID: "3211b2b2-3a5c-4ac0-9f80-f328b5f1c3c4",
	}

	mockGetRelationshipINDB := model.Relationship{
		PersonID: "3211b2b2-3a5c-4ac0-9f80-f328b5f1c3c4",
		Members: []model.RelationshipMember{
			{
				PersonID: "f24a439a-b5f0-4577-afcc-1b9e15b5b556",
				Type:     relationship.Parent,
				Kindship: kinship.Parent,
			},
			{
				PersonID: "238a257f-c38d-48bd-b7c0-ba4d31f8e7bd",
				Type:     relationship.Parent,
				Kindship: kinship.Parent,
			},
			{
				PersonID: "9bd0ba44-1c11-49b6-96f4-4b95e5c545e6",
				Type:     relationship.Sibling,
				Kindship: kinship.Sibling,
			},
			{
				PersonID: "ce1e6b7a-df80-41c3-99b4-51139162eb62",
				Type:     relationship.Sibling,
				Kindship: kinship.Sibling,
			},
		},
	}

	persistenceDB.EXPECT().Get(ctx, model.Relationship{PersonID: "3211b2b2-3a5c-4ac0-9f80-f328b5f1c3c4"}).Return(mockGetRelationshipINDB, nil)

	relationship, err := relationshipService.getKinship(ctx, mockToGetRelationship)

	require.Nil(t, err, "description of kinship of relationship returned with success, haven't error")

	_ = relationship

}
