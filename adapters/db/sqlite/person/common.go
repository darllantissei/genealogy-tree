package dbperson

import (
	"github.com/darllantissei/genealogy-tree/application/model"
)

func (db PersonDB) getFields(prsn model.Person) []interface{} {
	return []interface{}{
		prsn.FirstName,
		prsn.LastName,
		prsn.Gender,
	}
}
