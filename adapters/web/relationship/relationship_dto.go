package relationshipweb

import (
	"reflect"
)

type RelationshipDTO struct {
	XMLName  struct{}                `json:"-" xml:"relationship"`
	ID       string                  `json:"id,omitempty" xml:"id"`
	PersonID string                  `json:"person_id,omitempty" xml:"person_id"`
	Members  []RelationshipMemberDTO `json:"members,omitempty" xml:"members"`
}

func (rtshp *RelationshipDTO) IsEmpty() bool {

	if rtshp == nil {
		return true
	}

	if reflect.DeepEqual(rtshp, &RelationshipDTO{}) {
		return true
	}

	return false
}
