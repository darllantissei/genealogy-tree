package relationship

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/darllantissei/genealogy-tree/application/common"
	"github.com/darllantissei/genealogy-tree/application/model"
)

func (r *RelationshipService) checkRelationship(ctx context.Context, rtshp model.Relationship) (msgErr []string) {

	_, err := common.Validation(rtshp)

	if err != nil {
		msgErr = append(msgErr, strings.Split(err.Error(), ";")...)
	}

	if rtshp.IsEmpty() {
		msgErr = append(msgErr, "Please, the relationship is required")
	}

	if len(rtshp.Members) <= 0 {
		msgErr = append(msgErr, "Please, declare some relationship")
	}

	if len(rtshp.Members) == 1 {
		msgErr = append(msgErr, "Please, declare two or more relationship between the persons")
	}

	err = r.checkSamePerson(ctx, rtshp)

	if err != nil {
		msgErr = append(msgErr, err.Error())
	}

	err = r.relationshipUnallowed(ctx, rtshp)

	if err != nil {
		msgErr = append(msgErr, err.Error())
	}

	msgErr = append(msgErr, r.checkPersonExistis(ctx, rtshp)...)

	return msgErr
}

func (r *RelationshipService) checkSamePerson(ctx context.Context, rtshp model.Relationship) error {

	for _, link := range rtshp.Members {

		if link.PersonID == rtshp.PersonID {
			return errors.New("the person can't have a relationship with himself")
		}

	}

	return nil
}

func (r *RelationshipService) relationshipUnallowed(ctx context.Context, rtshp model.Relationship) error {

	err := r.checkDependencyInjection(ctx)

	if err != nil {
		return err
	}

	for _, memberRequest := range rtshp.Members {

		rtshpDB, err := r.PersistenceDB.Get(ctx, model.Relationship{PersonID: memberRequest.PersonID})

		if err != nil {
			return err
		}

		for _, memberDB := range rtshpDB.Members {

			if rtshp.PersonID == memberDB.PersonID && memberDB.Type != memberRequest.Type {
				person, err := r.PersonService.Fetch(ctx, model.Person{ID: rtshpDB.PersonID})

				if err != nil {
					return fmt.Errorf("fail get person for information that relationship unallowed")
				}

				return fmt.Errorf("this kind of relationship unallowed, because this person is %s of %s-%s", memberDB.Type, person.ID, person.FullName())
			}

		}
	}

	return nil
}

func (r *RelationshipService) checkPersonExistis(ctx context.Context, rtshp model.Relationship) []string {

	err := r.checkDependencyInjection(ctx)

	if err != nil {
		return []string{err.Error()}
	}

	getPerson := func(personID string) error {

		person, err := r.PersonService.Fetch(ctx, model.Person{ID: personID})

		if err != nil {
			return err
		}

		if person.IsEmpty() {
			return fmt.Errorf("the ID of person %s not exists", personID)
		}

		return nil
	}

	err = getPerson(rtshp.PersonID)

	if err != nil {
		return []string{err.Error()}
	}

	collectionErrors := []string{}
	for _, member := range rtshp.Members {

		err := getPerson(member.PersonID)

		if err != nil {
			collectionErrors = append(collectionErrors, err.Error())
		}
	}

	return collectionErrors
}
