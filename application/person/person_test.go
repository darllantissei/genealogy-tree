package person

import (
	"context"
	"testing"

	"github.com/darllantissei/genealogy-tree/application/model"
	"github.com/stretchr/testify/require"
)

func TestPersonService_checkDataPerson(t *testing.T) {

	const testCheckDataPerson = "test_check_data_person"

	personService := PersonService{}

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey(testCheckDataPerson), "test validations of person")

	person := model.Person{
		FirstName: "",
		LastName:  "",
	}

	msgErr := personService.checkDataPerson(ctx, person)

	require.Greater(t, len(msgErr), 0, "There are errors")

	ctx = context.WithValue(ctx, contextKey(testCheckDataPerson), "test validation gender of person")

	person.FirstName = "John"
	person.Gender = 1324

	msgErr = personService.checkDataPerson(ctx, person)

	require.Greater(t, len(msgErr), 0, "There are error")

}
