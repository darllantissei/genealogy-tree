package person

import (
	"context"
	"testing"

	mockperson "github.com/darllantissei/genealogy-tree/mocks/person"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

func TestPersonService_checkDependencyInjection(t *testing.T) {

	personService := PersonService{}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey("check_dependecy"), "check dependecy injection from PersonService undeclared")

	err := personService.checkDependencyInjection(ctx)

	require.NotNil(t, err, "The dependecy injection database undeclared")

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	persistenceDB := mockperson.NewMockIPersonPersistenceDB(ctrl)

	personService.PersistenceDB = persistenceDB

	ctx = context.WithValue(ctx, contextKey("check_dependecy"), "check dependency injection from PersonService declared")

	err = personService.checkDependencyInjection(ctx)

	require.Nil(t, err, "The dependecy injection declared, error must be nil")

}
