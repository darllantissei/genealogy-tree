package model

import (
	enumkinship "github.com/darllantissei/genealogy-tree/application/enum/kinship"
	enumrelationship "github.com/darllantissei/genealogy-tree/application/enum/relationship"
)

type RelationshipMember struct {
	PersonID       string                        `valid:"required~The member of relationship is required"`
	RelationshipID string                        `valid:"-"`
	Type           enumrelationship.Relationship `valid:"required~The type of relationship is required"`
	Kindship       enumkinship.Kinship           `valid:"-"`
}
