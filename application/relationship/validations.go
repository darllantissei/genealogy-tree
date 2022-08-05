package relationship

import (
	"context"
	"errors"

	"github.com/darllantissei/genealogy-tree/application/model"
)

func (r *RelationshipService) checkRelationship(ctx context.Context, rtshp model.Relationship) (msgErr []string) {

	if rtshp.IsEmpty() {
		msgErr = append(msgErr, "Please, the relationship is required")
	}

	if len(rtshp.Members) <= 0 {
		msgErr = append(msgErr, "Please, declare some relationship")
	}

	if len(rtshp.Members) == 1 {
		msgErr = append(msgErr, "Please, declare two or more relationship between the persons")
	}

	err := r.checkSamePerson(ctx, rtshp)

	if err != nil {
		msgErr = append(msgErr, err.Error())
	}

	err = r.relationshipUnallowed(ctx, rtshp)

	if err != nil {
		msgErr = append(msgErr, err.Error())
	}

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

	for _, checkMember := range rtshp.Members {

		rtshpDB, err := r.PersistenceDB.Get(ctx, model.Relationship{PersonID: checkMember.PersonID})

		if err != nil {
			return err
		}

		for _, kindship := range rtshpDB.Members {

			if kindship.Type != checkMember.Type {
				return errors.New("this kind of relationship unallowed")
			}
		}

	}

	return nil
}
