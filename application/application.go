package application

import (
	"github.com/darllantissei/genealogy-tree/application/common"
	"github.com/darllantissei/genealogy-tree/application/person"
	"github.com/darllantissei/genealogy-tree/application/relationship"
)

func init() {
	common.SetGoValidator()
}

type Application struct {
	PersonService       person.IPersonService
	RelationshipService relationship.IRelationshipService
	CommonService       common.ICommonService
}
