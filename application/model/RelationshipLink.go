package model

import enumrelationship "github.com/darllantissei/genealogy-tree/application/enum/relationship"

type RelationshipLink struct {
	PersonID       string
	RelationshipID string
	Type           enumrelationship.Relationship
}
