package model

import enumrelationship "github.com/darllantissei/genealogy-tree/application/enum/relationship"

type Relationship struct {
	ID             string
	PersonID       string
	Type           enumrelationship.Relationship
	RelationshipID string
}
