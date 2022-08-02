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

	if len(rtshp.Links) <= 0 {
		msgErr = append(msgErr, "Please, declare some relationship")
	}

	if len(rtshp.Links) == 1 {
		msgErr = append(msgErr, "Please, declare two or more relationship between the persons")
	}

	err := r.checkSamePerson(ctx, rtshp)

	if err != nil {
		msgErr = append(msgErr, err.Error())
	}

	return msgErr
}

func (r *RelationshipService) checkSamePerson(ctx context.Context, rtshp model.Relationship) error {

	for _, link := range rtshp.Links {

		if link.PersonID == rtshp.PersonID {
			return errors.New("the person can't have a relationship with himself")
		}

	}

	return nil
}
