package relationship

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/darllantissei/genealogy-tree/application/common"
	"github.com/darllantissei/genealogy-tree/application/enum/relationship"
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

	if len(rtshp.Members) < 1 {
		msgErr = append(msgErr, "Please, declare some relationship")
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

	membersDB, err := r.PersistenceDB.GetKinship(ctx, model.RelationshipMember{PersonID: rtshp.PersonID})

	if err != nil {
		return err
	}

nextMemberRequest:
	for _, memberRequest := range rtshp.Members {

	checkMember:
		for _, memberDB := range membersDB {

			if strings.EqualFold(memberDB.PersonID, rtshp.PersonID) {
				continue checkMember
			}

			if !strings.EqualFold(memberDB.PersonID, memberRequest.PersonID) {
				continue checkMember
			}

			switch {
			case memberDB.Type == relationship.Parent && memberRequest.Type == relationship.Spouse:
				continue nextMemberRequest
			case memberDB.Type == relationship.Sibling && memberRequest.Type == relationship.Child:
				continue nextMemberRequest
			case memberDB.Type == relationship.Child && memberRequest.Type == relationship.Child:
				continue nextMemberRequest
			case memberDB.Type == relationship.Sibling && memberRequest.Type == relationship.Sibling:
				continue nextMemberRequest
			case memberDB.Type == relationship.Spouse && memberRequest.Type == relationship.Parent:
				continue nextMemberRequest
			case memberDB.Type == relationship.Child && memberRequest.Type == relationship.Sibling:
				continue nextMemberRequest
			default:

				return fmt.Errorf("invalid relationship")

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

func (r *RelationshipService) checkRelationshipExists(ctx context.Context, rtshp model.Relationship) error {

	err := r.checkDependencyInjection(ctx)

	if err != nil {
		return err
	}

	rtshpDB, err := r.PersistenceDB.Get(ctx, model.Relationship{PersonID: rtshp.PersonID})

	if err != nil {
		return fmt.Errorf("fail to check relationship exists. Details: %s", err.Error())
	}

	if !rtshpDB.IsEmpty() && len(rtshpDB.Members) > 0 {
		return fmt.Errorf("this relationship already exists")
	}

	return nil

}
