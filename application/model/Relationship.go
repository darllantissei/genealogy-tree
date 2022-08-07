package model

import "reflect"

type Relationship struct {
	ID       string               `valid:"-"`
	PersonID string               `valid:"required~The person is required"`
	Members  []RelationshipMember `valid:"required~The person must have one or more members"`
}

func (rtshp *Relationship) IsEmpty() bool {

	if rtshp == nil {
		return true
	}

	if reflect.DeepEqual(rtshp, &Relationship{}) {
		return true
	}

	return false
}
