package model

import (
	enumkinship "github.com/darllantissei/genealogy-tree/application/enum/kinship"
	enumrelationship "github.com/darllantissei/genealogy-tree/application/enum/relationship"
)

type RelationshipMember struct {
	PersonID       string
	RelationshipID string
	Type           enumrelationship.Relationship
	Kindship       enumkinship.Kinship
}
