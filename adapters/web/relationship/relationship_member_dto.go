package relationshipweb

import (
	enumkinship "github.com/darllantissei/genealogy-tree/application/enum/kinship"
	enumrelationship "github.com/darllantissei/genealogy-tree/application/enum/relationship"
)

type RelationshipMemberDTO struct {
	XMLName        struct{}                      `json:"-" xml:"member"`
	PersonID       string                        `json:"person_id,omitempty" xml:"person_id"`
	RelationshipID string                        `json:"relationship_id,omitempty" xml:"relationship_id"`
	Type           enumrelationship.Relationship `json:"type,omitempty" xml:"type"`
	Kindship       enumkinship.Kinship           `json:"kindship,omitempty" xml:"kinship"`
}
