package person

import (
	"context"
	"testing"

	"github.com/darllantissei/genealogy-tree/application/enum/gender"
	"github.com/darllantissei/genealogy-tree/application/model"
	mockperson "github.com/darllantissei/genealogy-tree/mocks/person"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPersonService_Record(t *testing.T) {

	const testRecordPerson = "test_record_person"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ctx := context.Background()

	persistenceDB := mockperson.NewMockIPersonPersistenceDB(ctrl)

	personService := PersonService{
		PersistenceDB: persistenceDB,
	}

	ctx = context.WithValue(ctx, contextKey(testRecordPerson), "Record a person")

	personToRecord := model.Person{
		FirstName: "John",
		LastName:  "Wick",
		Gender:    gender.Masculine,
	}

	personExpected := personToRecord
	personExpected.ID = "0509b0d9-72bb-4230-a9ad-7f763c0e0eb8"

	persistenceDB.EXPECT().Insert(ctx, personToRecord).Return(personExpected, nil)

	personRecorded, err := personService.Record(ctx, personToRecord)

	require.Nil(t, err, "Person recorded with success, error must be nil")

	require.NotEmpty(t, personRecorded.ID, "Person recorded, the field ID must be returned with value")

	ctx = context.WithValue(ctx, contextKey(testRecordPerson), "Update a person")

	personToRecord.LastName = "Foo"

	personExpected = personToRecord
	personExpected.ID = "0509b0d9-72bb-4230-a9ad-7f763c0e0eb8"

	personToRecord.ID = personExpected.ID

	persistenceDB.EXPECT().Update(ctx, personToRecord).Return(nil)

	personRecorded, err = personService.Record(ctx, personToRecord)

	require.Nil(t, err, "Person recorded with succes, error must be nil")

	require.Equal(t, personExpected.ID, personRecorded.ID, "The field ID must be returned equal when statement")

	require.Equal(t, personExpected, personRecorded, "The person updated must be returned as expected")

}

func TestPersonService_Fetch(t *testing.T) {

	const testFetchPerson = "test_fetch_person"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ctx := context.Background()

	persistenceDB := mockperson.NewMockIPersonPersistenceDB(ctrl)

	personService := PersonService{
		PersistenceDB: persistenceDB,
	}

	ctx = context.WithValue(ctx, contextKey(testFetchPerson), "Fetch a person")

	mockPersonFetch := model.Person{
		ID:        "15eec2f9-36ac-48b6-aa63-13d9fab89ec9",
		FirstName: "John",
		LastName:  "Wick",
		Gender:    gender.Masculine,
	}

	persistenceDB.EXPECT().Get(ctx, model.Person{ID: mockPersonFetch.ID}).Return(mockPersonFetch, nil)

	personFetched, err := personService.Fetch(ctx, model.Person{ID: "15eec2f9-36ac-48b6-aa63-13d9fab89ec9"})

	require.Nil(t, err, "The person fetched success, error must be nil")

	require.Equal(t, mockPersonFetch, personFetched, "The person fetched must be same of mock")

	ctx = context.WithValue(ctx, contextKey(testFetchPerson), "Fetch a person with no record")

	persistenceDB.EXPECT().Get(ctx, model.Person{ID: "84763d21-497d-41f3-8381-727724914576"}).Return(model.Person{}, nil)

	personFetched, err = personService.Fetch(ctx, model.Person{ID: "84763d21-497d-41f3-8381-727724914576"})

	require.Nil(t, err, "The person fetched success, error must be nil")

	require.Empty(t, personFetched, "The person not found, the struct must be returned empty")

}
