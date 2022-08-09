package relationship

import (
	"context"
	"testing"

	"github.com/darllantissei/genealogy-tree/application/enum/relationship"
	"github.com/darllantissei/genealogy-tree/application/model"
	mockperson "github.com/darllantissei/genealogy-tree/mocks/person"
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

	personService := mockperson.NewMockIPersonService(ctrl)

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testRelationshipUnallowed), "Unallowed incest between sibling")

	relationshipService := RelationshipService{
		PersistenceDB: persistenceDB,
		PersonService: personService,
	}

	mockRelationshipA := model.Relationship{
		ID:       "101",
		PersonID: "1",
		Members: []model.RelationshipMember{
			{
				PersonID:       "2",
				RelationshipID: "101",
				Type:           relationship.Spouse,
			},
		},
	}

	mockRelationshipB := model.Relationship{
		ID:       "201",
		PersonID: "6",
		Members: []model.RelationshipMember{
			{
				PersonID:       "5",
				RelationshipID: "201",
				Type:           relationship.Sibling,
			},
			{
				PersonID:       "1",
				RelationshipID: "101",
				Type:           relationship.Sibling,
			},
		},
	}

	persistenceDB.EXPECT().GetKinship(ctx, model.RelationshipMember{PersonID: mockRelationshipA.PersonID}).Return(mockRelationshipB.Members, nil)

	err := relationshipService.relationshipUnallowed(ctx, mockRelationshipA)

	require.Nil(t, err, "Relationship allowed")

	ctx = context.WithValue(ctx, contextKey(testRelationshipUnallowed), "Unallowed incest between child and mother")

	mockRelationshipA = model.Relationship{
		ID:       "104",
		PersonID: "5",
		Members: []model.RelationshipMember{
			{
				PersonID:       "12",
				RelationshipID: "104",
				Type:           relationship.Spouse,
			},
		},
	}

	mockRelationshipB = model.Relationship{
		ID:       "204",
		PersonID: "12",
		Members: []model.RelationshipMember{
			{
				PersonID:       "5",
				RelationshipID: "204",
				Type:           relationship.Child,
			},
		},
	}

	persistenceDB.EXPECT().GetKinship(ctx, model.RelationshipMember{PersonID: mockRelationshipA.PersonID}).Return(mockRelationshipB.Members, nil)

	err = relationshipService.relationshipUnallowed(ctx, mockRelationshipA)

	require.Nil(t, err, "Relationship allowed found")

	ctx = context.WithValue(ctx, contextKey(testRelationshipUnallowed), "Relationship sibling")

	mockRelationshipA = model.Relationship{
		ID:       "401",
		PersonID: "4",
		Members: []model.RelationshipMember{
			{
				PersonID:       "3",
				RelationshipID: "401",
				Type:           relationship.Sibling,
			},
		},
	}

	mockRelationshipB = model.Relationship{
		ID:       "501",
		PersonID: "3",
		Members: []model.RelationshipMember{
			{
				PersonID:       "4",
				RelationshipID: "501",
				Type:           relationship.Sibling,
			},
		},
	}

	persistenceDB.EXPECT().GetKinship(ctx, model.RelationshipMember{PersonID: mockRelationshipA.PersonID}).Return(mockRelationshipB.Members, nil)

	err = relationshipService.relationshipUnallowed(ctx, mockRelationshipA)

	require.Nil(t, err, "Relationship correct")

}

func TestRelationshipService_checkPersonExistis(t *testing.T) {

	const testCheckPersonExists = "test_person_exists"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	persistenceDB := mockrelationship.NewMockIRelationshipPersistenceDB(ctrl)

	personService := mockperson.NewMockIPersonService(ctrl)

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testCheckPersonExists), "person not exists")

	relationshipService := RelationshipService{
		PersistenceDB: persistenceDB,
		PersonService: personService,
	}

	mockPersonID := "12lk3j12k3l"

	personService.EXPECT().Fetch(ctx, model.Person{ID: mockPersonID}).Return(model.Person{}, nil)

	msgErr := relationshipService.checkPersonExistis(ctx, model.Relationship{PersonID: mockPersonID})

	require.Greater(t, len(msgErr), 0, "The person not found, will returned error in list")

	personService.EXPECT().Fetch(ctx, model.Person{ID: mockPersonID}).Return(model.Person{ID: mockPersonID}, nil)

	msgErr = relationshipService.checkPersonExistis(ctx, model.Relationship{PersonID: mockPersonID})

	require.Empty(t, msgErr, "The person exist, then list error will return empty")

}

func TestRelationshipService_checkRelationshipExists(t *testing.T) {

	const testRelationshipAlreadExists = "test_relationship_already_exists"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	persistenceDB := mockrelationship.NewMockIRelationshipPersistenceDB(ctrl)

	personService := mockperson.NewMockIPersonService(ctrl)

	relationshipService := RelationshipService{
		PersistenceDB: persistenceDB,
		PersonService: personService,
	}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testRelationshipAlreadExists), "check relationship already exists - relationship not exists")

	mockPersonID := "23lk4j3lk4j23"

	persistenceDB.EXPECT().Get(ctx, model.Relationship{PersonID: mockPersonID}).Return(model.Relationship{}, nil)

	err := relationshipService.checkRelationshipExists(ctx, model.Relationship{PersonID: mockPersonID})

	require.Nil(t, err, "Relationship not exists, then error must be nil")

	ctx = context.WithValue(ctx, contextKey(testRelationshipAlreadExists), "check relationship already exist - relationship exists")

	persistenceDB.EXPECT().Get(ctx, model.Relationship{PersonID: mockPersonID}).Return(model.Relationship{PersonID: mockPersonID, Members: []model.RelationshipMember{{PersonID: "34kl5j34lkj"}}}, nil)

	err = relationshipService.checkRelationshipExists(ctx, model.Relationship{PersonID: mockPersonID})

	require.NotNil(t, err, "The relationship exists, then error must be returned")

}
