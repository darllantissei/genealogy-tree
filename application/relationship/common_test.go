package relationship

import (
	"context"
	"testing"

	mock_person "github.com/darllantissei/genealogy-tree/mocks/person"
	mockrelationship "github.com/darllantissei/genealogy-tree/mocks/relationship"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

func TestPersonService_checkDependencyInjection(t *testing.T) {

	relationshipService := RelationshipService{}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey("check_dependecy"), "check dependecy injection from RelationshipService undeclared")

	err := relationshipService.checkDependencyInjection(ctx)

	require.NotNil(t, err, "The dependecy injection database undeclared")

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	persistenceDB := mockrelationship.NewMockIRelationshipPersistenceDB(ctrl)

	personService := mock_person.NewMockIPersonService(ctrl)

	relationshipService.PersistenceDB = persistenceDB
	relationshipService.PersonService = personService

	ctx = context.WithValue(ctx, contextKey("check_dependecy"), "check dependency injection from RelationshipService declared")

	err = relationshipService.checkDependencyInjection(ctx)

	require.Nil(t, err, "The dependecy injection declared, error must be nil")

}
