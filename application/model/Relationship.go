package model

import "reflect"

type Relationship struct {
	ID       string
	PersonID string
	Members  []RelationshipMember
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
