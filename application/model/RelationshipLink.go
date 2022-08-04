package model

import enumrelationship "github.com/darllantissei/genealogy-tree/application/enum/relationship"

type RelationshipMember struct {
	PersonID       string
	RelationshipID string
	Type           enumrelationship.Relationship
}
