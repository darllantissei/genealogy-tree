package relationship

import (
	"context"
	"testing"

	statusapplication "github.com/darllantissei/genealogy-tree/application/enum/status_application"
	"github.com/darllantissei/genealogy-tree/application/model"
	mock_common "github.com/darllantissei/genealogy-tree/mocks/common"
	mock_person "github.com/darllantissei/genealogy-tree/mocks/person"
	mockrelationship "github.com/darllantissei/genealogy-tree/mocks/relationship"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRelationshipService_Create(t *testing.T) {

	const testCreateRelationship = "test_to_create_relationship"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	persistenceDB := mockrelationship.NewMockIRelationshipPersistenceDB(ctrl)

	personService := mock_person.NewMockIPersonService(ctrl)

	commonService := mock_common.NewMockICommonService(ctrl)

	relationshipService := RelationshipService{
		PersistenceDB: persistenceDB,
		PersonService: personService,
		CommonService: commonService,
	}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testCreateRelationship), "Don't allow create relationship with only person")

	mockRelationship := model.Relationship{
		PersonID: "84763d21-497d-41f3-8381-727724914576",
		// Type:     relationship.Child,
	}

	personService.EXPECT().Fetch(ctx, model.Person{ID: mockRelationship.PersonID}).Return(model.Person{ID: mockRelationship.PersonID}, nil)

	persistenceDB.EXPECT().Get(ctx, model.Relationship{PersonID: mockRelationship.PersonID}).Return(mockRelationship, nil)

	persistenceDB.EXPECT().GetKinship(ctx, model.RelationshipMember{PersonID: mockRelationship.PersonID}).Return(mockRelationship.Members, nil)

	commonService.EXPECT().BuildError(ctx, []string{"The person must have one or more members", "Please, declare some relationship"}).Return(model.Returns{Return: model.Return{Status: statusapplication.Error, Message: []string{"The person must have one or more members", "Please, declare some relationship"}}})

	relationship, err := relationshipService.Create(ctx, mockRelationship)

	require.NotNil(t, err, "Error occurred, the error must be returned")

	require.Empty(t, relationship, "Error occurred, the struct must be empty")

}
